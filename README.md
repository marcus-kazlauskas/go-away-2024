# Advent of Code

[adventofcode.com](https://adventofcode.com)

## REST-service for puzzles solving

Version 1.0.0

### Supported puzzles

| Year of the event | Days with part 1 | Days with part 2 |
|:------------------|:-----------------|:-----------------|
| 2024              | 1                | 1                |
| 2025              | 1-6              | 1-6              |

### API description

Openapi description of supported methods is located in [openapi-go-away-2024.yml](api/openapi-go-away-2024.yml).

You can generate actual API interface:

```shell
go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
cd internal/api
go tool oapi-codegen -config api-codegen.yml ../../api/openapi-go-away-2024.yml
go tool oapi-codegen -config types-codegen.yml ../../api/openapi-go-away-2024.yml
cd ../..
```

### Launch application

Run local environment:

```shell
cd local-env
podman compose up -d
cd ..
```

Apply all available migrations:

```shell
cd db-migrations
go get -tool github.com/pressly/goose/v3/cmd/goose@latest
go tool goose up
cd ..
```

Start application:
```shell
cd cmd
go run main.go
```

### Application diagram

![go-away-2024.drawio](docs/go-away-2024.drawio.svg)
