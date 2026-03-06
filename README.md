# About Go language practice project
The project is to practice Go language

## Tech Stack
- Go (v1.26)

## How to run locally
```bash
go run .
```
- Chose from the list of available exercises

## Run lint tests
### Built-in lint test
```bash
go vet ./...
```

### Production-like lint test
#### Install `golangci-lint`
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

#### Lint everything
```bash
golangci-lint run
```

#### Lint with more details
```bash
golangci-lint run --verbose
```

#### Run with its own auto-fix
```bash
golangci-lint run --fix
```

## Run test cases with race detector
```bash
go test -v -race ./...
```
