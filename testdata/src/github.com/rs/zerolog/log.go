package zerolog

type Logger struct{}

func New(w any) Logger {
	return Logger{}
}

type Event struct{}

func (l Logger) Error() *Event {
	return &Event{}
}

func (l Logger) Warn() *Event {
	return &Event{}
}

func (e *Event) Str(key, value string) *Event {
	return e
}

func (e *Event) Bool(key string, value bool) *Event {
	return e
}

func (e *Event) Msg(msg string) {}

func (e *Event) Msgf(format string, v ...any) {}
