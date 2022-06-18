package herror

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.NewConsoleWriter()).
		With().
		CallerWithSkipFrameCount(3).
		Logger()
}
