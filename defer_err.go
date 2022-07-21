package herror

type DeferFuncErr struct {
	f func() error
}

func (m DeferFuncErr) Must() {
	if err := m.f(); err != nil {
		zlog.Panic().Str("Must", "DeferErr err").Err(err).Send()
	}
}

func (m DeferFuncErr) Ignore() {
	if err := m.f(); err != nil {
		zlog.Debug().Str("Ignore", "DeferErr err").Err(err).Send()
	}
}
