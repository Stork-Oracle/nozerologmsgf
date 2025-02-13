package p

import zl "github.com/rs/zerolog"

func alias_example() {
	logger := zl.New(nil)
	// This should trigger our custom linter
	logger.Error().Msgf("This is a formatted error message %s", "test")                                        // want "Do not use zerolog .Msgf after zerolog .Error; include extra info in Event fields"
	logger.Error().Str("key", "value").Msgf("This is a formatted error message %s", "test")                    // want "Do not use zerolog .Msgf after zerolog .Error; include extra info in Event fields"
	logger.Error().Str("key", "value").Bool("bool", true).Msgf("This is a formatted error message %s", "test") // want "Do not use zerolog .Msgf after zerolog .Error; include extra info in Event fields"

	// These should be fine
	logger.Error().Msg("This is a regular error message")
	logger.Error().Str("key", "value").Msg("This is a regular error message")
	logger.Warn().Msg("This is a regular warn message")
	logger.Warn().Str("key", "value").Msg("This is a regular warn message")
	logger.Warn().Msgf("This is a formatted warn message %s", "test")
	logger.Warn().Str("key", "value").Msgf("This is a formatted warn message %s", "test")
}
