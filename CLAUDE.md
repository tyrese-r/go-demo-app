# Go Demo App Development Guide

## Build & Test Commands

- Build: `go build -v ./...`
- Run: `go run main.go`
- Run all tests: `go test -v ./...`
- Run single test: `go test -v ./internal/utils/hello/hello_test.go -run TestHelloWorld`
- Lint: `golangci-lint run ./...`

## Code Style

- Use `gofmt` for formatting
- Organize imports in blocks: standard library, external, internal
- Use structured error handling with explicit returns
- Function parameters should use Go's standard parameter ordering
- Use meaningful variable names in camelCase
- Error messages should be lowercase without punctuation
- Use dependency injection for services and handlers
- Use the logger package for all logging (`Info`, `Warn`, `Error`)
- Package structure follows domain-driven design with internal modules

## Security

- Use the secrets package for all environment variables
- JWT tokens handled via the auth service
- Handle SQL with properly parameterized queries
- Apply `#nosec` comments only when necessary with justification