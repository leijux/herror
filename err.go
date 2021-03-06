package herror

import (
	"github.com/rs/zerolog"
)

type Err struct {
	err error
	msg string
}

func (m Err) Msg(msg string) Err {
	if m.err != nil {
		m.msg = msg
	}
	return m
}

func (m Err) Must() {
	var event *zerolog.Event
	if m.err != nil {
		event = zlog.Panic().Err(m.err).Str("Must", "HandleErr err")
	}
	if m.msg != "" {
		event.Msg(m.msg)
	} else {
		event.Send()
	}
}

func (m Err) Ignore() {
	var event *zerolog.Event
	if m.err != nil {
		event = zlog.Debug().Err(m.err).Str("Ignore", "HandleErr err")
	}
	if m.msg != "" {
		event.Msg(m.msg)
	} else {
		event.Send()
	}
}
