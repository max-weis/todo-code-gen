# Go Todo Code Gen App

A sample project which demonstrates code generation in Go. Following tools are used:

| Tool  |  Usecase |
|---|---|
|  [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc) |  sql queries |
|  [google/wire](https://github.com/google/wire) | dependency injections  |
|  [deepmap](https://github.com/deepmap/oapi-codegen) | http handlers  |

## Generate HTTP Router

Run: `oapi-codegen -generate types,server,spec ./internal/todo/spec.yaml > ./internal/todo/boundary/router_delegate_gen.go`

to generate the http router. Change the package of the generated code to ***boundary***
