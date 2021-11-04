# Go Todo Code Gen App

A sample project which demonstrates code generation in Go. Following tools are used:

| Tool  |  Usecase |
|---|---|
|  [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc) |  sql queries |
|  [google/wire](https://github.com/google/wire) | dependency injections  |
|  [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen) | http handlers  |
| [vektra/mockery](https://github.com/vektra/mockery) |mocks|

**install all these tools to build the project on your machine**

## Generate HTTP routes

Run: `oapi-codegen -generate types,server,spec ./internal/todo/spec.yaml > ./internal/todo/boundary/router_delegate_gen.go`
to generate the http router.

Change the package of the generated code to ***boundary***

## Generate SQL queries

Run: `sqlc generate -f internal/todo/sqlc.yaml` to generate the mysql code

## Generate dependency injection

Run: `wire` to generate the dependency injection

## Generate mocks

Run: `mockery --all --recursive && rm mocks/DBTX.go mocks/EchoRouter.go mocks/ServerInterface.go`