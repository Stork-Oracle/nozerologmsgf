# No Zerolog .Msgf Usage After .Error
Simple linter checks for .Msgf(...) usage after chaining .Error() on a zerolog Event.

The linter suggests including the extra log information as Event fields (e.g. .Str("key", "value")).

## Build

```bash
go build -buildmode=plugin plugin/nozerologmsgf.go
```
