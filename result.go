package herror

import (
	"github.com/rs/zerolog"
)

type ErrResult[T any] struct {
	err    error
	msg    string
	result T
}

func (m ErrResult[T]) Msg(msg string) ErrResult[T] {
	if m.err != nil {
		m.msg = msg
	}
	return m
}

func (m ErrResult[T]) Must() (result T) {
	var event *zerolog.Event
	if m.err != nil {
		event = zlog.Panic().Err(m.err).Str("Must", "ResultErr err")
	}
	if m.msg != "" {
		event.Msg(m.msg)
	} else {
		event.Send()
	}
	return m.result
}

func (m ErrResult[T]) Ignore() (result T) {
	var event *zerolog.Event
	if m.err != nil {
		event = zlog.Debug().Err(m.err).Str("Ignore", "ResultErr err")
	}
	if m.msg != "" {
		event.Msg(m.msg)
	} else {
		event.Send()
	}
	return m.result
}
