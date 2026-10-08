# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

A Go REST service that solves Advent of Code puzzles asynchronously. A client uploads puzzle input for a given year/day/part, gets a task ID back, and polls for the result. The supported puzzles are listed in the README table. Update that table when you add a puzzle.

## Commands

Requires Go 1.26+, podman (or docker), and `goose`, `golangci-lint` (v2) and `oapi-codegen` installed as standalone binaries. They are not `go tool` dependencies in `go.mod`.

Several paths are relative, so commands must be run from specific directories:

```shell
# Local infra (Postgres, Kafka + kafka-ui on :8080, MinIO on :9000/:9001), via podman
cd local-env && podman compose up -d

# DB migrations (connection string in db-migrations/.env)
cd db-migrations && goose up

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
oapi-codegen -config api-codegen.yml ../../api/openapi-go-away-2024.yml
oapi-codegen -config types-codegen.yml ../../api/openapi-go-away-2024.yml
```

`internal/database` tests are integration tests. They need the local Postgres running with migrations applied, and they load config from `config.TEST_PATH` (`../../properties/go-away-2024.yml`). `internal/puzzles` tests are pure and open their `*_test.txt` fixtures by relative path.

Imports: keep `go-away-2024/...` in the same block as stdlib (the existing style); do not regroup them. goimports accepts either style, so the linter will not catch this.

## Architecture

One process (`cmd/main.go`) runs two goroutines; the first error from either kills the process (no graceful shutdown).

- **HTTP server** (`internal/aoc_server`, Fiber): `POST /task/create` stores the request in Postgres, the input in MinIO, and publishes a task to Kafka; `GET /task/{id}` reads request + result. Requests are validated against the OpenAPI spec embedded in `internal/api`. All errors go through `SendServerError` (the app's `ErrorHandler`) and are returned as JSON `ErrorResponse`, so handlers just `return err`.
- **Calculator** (`internal/aoc_calc`): polls Kafka, downloads the input from MinIO, runs the solver from `internal/puzzles`, and writes the answer or error to the `result` table.

`internal/api` is generated (`*.gen.go`): never edit it by hand, regenerate it from the spec.

## Adding a puzzle

1. Add `internal/puzzles/yearYYYYdayD.go` with `func YearYYYYDayDPart1(scan *bufio.Scanner) (*string, error)` and a matching `Part2`. Return `DataError()` when the input is malformed.
2. Add the example input as `yearYYYYdayD_test.txt` and add `t.Run("Year YYYY day D part N", ...)` cases to `puzzles_test.go`.
3. Wire both parts into the nested year/day/part `switch` in `aoc_calc.Calculator.calculate`.
4. Update the supported-puzzles table in `README.md`.
