default: run

run:
	@go run ./...

test:
	@go test ./... -short

it:
	@go test ./...
