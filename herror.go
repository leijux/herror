package herror

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var zlog zerolog.Logger

func init() {
	zlog = log.Output(zerolog.NewConsoleWriter()).
		With().
		CallerWithSkipFrameCount(3).
		Logger()
}
