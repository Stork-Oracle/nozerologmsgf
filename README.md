# No Zerolog .Msgf Usage After .Error
Simple linter checks for .Msgf(...) usage after chaining .Error() on a zerolog Event.

The linter suggests including the extra log information as Event fields (e.g. .Str("key", "value")).

The compiled linter for golangci-lint use is `cmd/no-zerolog-msgf/nozerologmsgf.so`.
