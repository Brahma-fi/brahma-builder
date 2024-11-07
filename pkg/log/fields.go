package log

import "github.com/rs/zerolog"

type Field func(e *zerolog.Event)

func Str(key, val string) Field {
	return func(e *zerolog.Event) {
		e.Str(key, val)
	}
}

func Err(err error) Field {
	return func(e *zerolog.Event) {
		e.Err(err)
	}
}

func Int(key string, val int) Field {
	return func(e *zerolog.Event) {
		e.Int(key, val)
	}
}

func Int64(key string, val int64) Field {
	return func(e *zerolog.Event) {
		e.Int64(key, val)
	}
}

func Any(key string, val any) Field {
	return func(e *zerolog.Event) {
		e.Any(key, val)
	}
}
