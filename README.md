# No Zerolog .Msgf Usage After .Error

A custom linter that checks for incorrect usage of `.Msgf(...)` after chaining `.Error()` on a zerolog Event.

## Purpose

This linter helps maintain consistent logging practices by detecting cases where `.Msgf(...)` is used after `.Error()` on a zerolog Event. Instead, it suggests including additional log information as Event fields (e.g. `.Str("key", "value")`).

## Building the Linter

### Prerequisites
- Go (version 1.24.3)
- golangci-lint (version 2.1.6)

### Installing Go
There were some issues when installing and using Go via brew. I recommend using gvm:
```
gvm install go1.24.3
gvm use go1.24.3 --default
```
You may have to add `[[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"` to the top of your .zshrc.
A known problem is the IDE not using go tools (gopls, dlv) built with the new go version.
Try running:
```
gvm pkgset create go1.24.3-global
gvm pkgset use go1.24.3-global --default
```

### Installing golangci-lint
```
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6
```
If you previously had golangci-lint installed, make sure that your path is pointing to the correct one.

### Building

Once it is confirmed that golangci-lint was build with the same version of go you are running, build the plugin by running:

```bash
golangci-lint custom -v
```

This needs to be run from the directory with .custom-gcl.yml, which is currently in /stork-aggregator. This command will create a custom executable version of golangci-lint, called custom-gcl, that can be run by the IDE. 

### Settings

```json
{
   ...
   "go.lintTool": "golangci-lint",
    "go.lintFlags": [
        "--config=${workspaceFolder}/.golangci.yaml"
    ],
    "go.alternateTools": {
        "golangci-lint": "${workspaceFolder}/custom-gcl",
        "customFormatter": "${workspaceFolder}/custom-gcl",
    },
    "go.lintOnSave": "package",
   ... 
}
```

Make sure your settings.json uses the new custom-gcl. You must also set the Go extension to use the pre-release version.
Once this is done, make sure you set your IDE's linter to golangci-lint-v2.

Also, make sure that your IDE is running the correct version of the go language server, gopls (version 0.18.1)

### Usage
With the settings.json configured as above, the IDE should automatically lint and highlight on save.
To run from the command line, look at the `make lint` command.