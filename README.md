# Advent of Code

[adventofcode.com](https://adventofcode.com)

## REST-service for puzzles solving

Version 1.0.0

Requirements:

- ***Go 1.26*** or higher
- ***Podman Desktop*** (or Docker Desktop)
- [***Goose***](https://github.com/pressly/goose)
- [***Task***](https://taskfile.dev/docs/installation)
- [***golangci-lint***](https://golangci-lint.run/docs/welcome/install/local/)
- [***oapi-codegen***](https://github.com/oapi-codegen/oapi-codegen)

### Supported puzzles

| Year of the event | Days with part 1 | Days with part 2 |
|:------------------|:-----------------|:-----------------|
| 2024              | 1                | 1                |
| 2025              | 1-6              | 1-6              |

### API description

Openapi description of supported methods is located in [openapi-go-away-2024.yml](api/openapi-go-away-2024.yml).

You can generate actual API interface:

```shell
cd internal/api
oapi-codegen -config api-codegen.yml ../../api/openapi-go-away-2024.yml
oapi-codegen -config types-codegen.yml ../../api/openapi-go-away-2024.yml
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
goose up
cd ..
```

Start application:
```shell
cd cmd
go run main.go
```

### Application diagram

![go-away-2024.drawio](docs/go-away-2024.drawio.svg)
