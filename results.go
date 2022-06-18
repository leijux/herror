package herror

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ErrResults[T1, T2 any] struct {
	err     error
	msg     string
	result1 T1
	result2 T2
}

func (m ErrResults[T1, T2]) Msg(msg string) ErrResults[T1, T2] {
	if m.err != nil {
		m.msg = msg
	}
	return m
}

func (m ErrResults[T1, T2]) Must() (T1, T2) {
	var event *zerolog.Event
	if m.err != nil {
		event = log.Fatal().Err(m.err).Str("Must", "ResultsErr err")
	}
	if m.msg != "" {
		event.Msg(m.msg)
	} else {
		event.Send()
	}
	return m.result1, m.result2
}

func (m ErrResults[T1, T2]) Ignore() (T1, T2) {
	var event *zerolog.Event
	if m.err != nil {
		event = log.Debug().Err(m.err).Str("Ignore", "ResultsErr err")
	}
	if m.msg != "" {
		event.Msg(m.msg)
	} else {
		event.Send()
	}
	return m.result1, m.result2
}
