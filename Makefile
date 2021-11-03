default: run

run:
	@go run ./...

test:
	@go test ./... -tags unit integration

test-unit:
	@go test ./... -tags unit

test-integration:
	@go test ./... -tags integration
