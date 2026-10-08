# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

A Go REST service that solves Advent of Code puzzles asynchronously. A client uploads puzzle input for a given year/day/part, gets a task ID back, and polls for the result. The supported puzzles are listed in the README table. Update that table when you add a puzzle.

## Commands

Several paths are relative, so commands must be run from specific directories:

```shell
# Local infra (Postgres, Kafka + kafka-ui on :8080, MinIO on :9000/:9001), via podman
cd local-env && podman compose up -d

# DB migrations (goose; connection string in db-migrations/.env)
cd db-migrations && go tool goose up

# Run the app — MUST be run from cmd/ (config.MAIN_PATH = "../properties/go-away-2024.yml")
cd cmd && go run main.go

# Tests
go test ./...
go test ./internal/puzzles -run 'TestPuzzles/Year_2025_day_6_part_1'   # single subtest

# Lint
golangci-lint run          # config: .golangci.yml (v2)
golangci-lint fmt          # gofmt + goimports

# Regenerate API code after editing api/openapi-go-away-2024.yml — run from internal/api
cd internal/api
go tool oapi-codegen -config api-codegen.yml ../../api/openapi-go-away-2024.yml
go tool oapi-codegen -config types-codegen.yml ../../api/openapi-go-away-2024.yml
```

`internal/database` tests are integration tests. They need the local Postgres running with migrations applied, and they load config from `config.TEST_PATH` (`../../properties/go-away-2024.yml`). `internal/puzzles` tests are pure and open their `*_test.txt` fixtures by relative path.

## Architecture

`cmd/main.go` runs two long-lived components in one process, and both share the same Repository, MinIO client and Kafka connection:

1. **HTTP server** (`internal/aoc_server`, Fiber). It implements `api.ServerInterface`, which is generated from the OpenAPI spec, and requests are validated against the embedded spec by `oapi-codegen/fiber-middleware`. `POST /task/create`:
   - saves a `request` row
   - uploads the raw body to MinIO (object name from `minio.NewPattern`) and stores the S3 link
   - publishes a `TaskMessage` to Kafka
   - inserts a `result` row with status `CREATED`

   `GET /task/{id}` joins request and result.
2. **Calculator** (`internal/aoc_calc`). This is a polling loop: it sleeps `calculator.sleep`, then reads a message from Kafka. Read timeouts are ignored. It skips tasks that are unknown or already `COMPLETED`/`ERROR`, downloads the input from MinIO, dispatches to a puzzle solver, and writes the answer (or the error text, with status `ERROR`) to `result`.

Supporting packages:
- `internal/api`: generated code (`api.gen.go`, `types.gen.go`). Do not hand-edit it. Regenerate it from the spec instead.
- `internal/database`: sqlx + pgx repository. The schema lives in `db-migrations/*.sql` (goose). Results are linked to requests with `on delete cascade`.
- `internal/utils`: mappers between API types, DB entities and Kafka messages, plus small generic math helpers.
- `internal/config`: a YAML config struct loaded from `properties/go-away-2024.yml`.

## Adding a puzzle

1. Add `internal/puzzles/yearYYYYdayD.go` with `func YearYYYYDayDPart1(scan *bufio.Scanner) (*string, error)` and a matching `Part2`. Return `DataError()` when the input is malformed.
2. Add the example input as `yearYYYYdayD_test.txt` and add `t.Run("Year YYYY day D part N", ...)` cases to `puzzles_test.go`.
3. Wire both parts into the nested year/day/part `switch` in `aoc_calc.Calculator.calculate`.
4. Update the supported-puzzles table in `README.md`.
