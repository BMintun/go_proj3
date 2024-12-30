build:
	@go build -o go_proj2

run: build
	@./go_proj2

test:
	@go test -v ./...