package herror

import "github.com/rs/zerolog/log"

type DeferFuncErr struct {
	f func() error
}

func (m DeferFuncErr) Must() {
	if err := m.f(); err != nil {
		log.Fatal().Str("Must", "DeferErr err").Err(err).Send()
	}
}

func (m DeferFuncErr) Ignore() {
	if err := m.f(); err != nil {
		log.Debug().Str("Ignore", "DeferErr err").Err(err).Send()
	}
}
