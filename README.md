# No Zerolog .Msgf Usage After .Error

A custom linter that checks for incorrect usage of `.Msgf(...)` after chaining `.Error()` on a zerolog Event.

## Purpose

This linter helps maintain consistent logging practices by detecting cases where `.Msgf(...)` is used after `.Error()` on a zerolog Event. Instead, it suggests including additional log information as Event fields (e.g. `.Str("key", "value")`).

## Building the Linter

### Prerequisites
- golangci-lint
- Go (version matching your golangci-lint installation)

### Version Compatibility

The linter must be built with dependencies matching your golangci-lint version. To check your golangci-lint version:

```bash
golangci-lint version
```

### Required Dependencies

1. Check your golangci-lint's tools version by running:
```bash
golangci-lint version --debug | grep "golang.org/x/tools"
```

2. Update `go.mod` to match:
   - Go version from golangci-lint
   - `golang.org/x/tools` version from the debug output

For example, with golangci-lint v1.64.8:
- Set Go version to 1.23.8
- Set `golang.org/x/tools` to v0.31.0
- Run `go mod tidy`

### Building

Once dependencies are aligned, build the plugin by running:

```bash
go build -buildmode=plugin plugin/nozerologmsgf.go
```

This will create a version of nozerologmsgf compatible with your golangci-lint installation.