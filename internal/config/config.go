package config

import (
	"github.com/rs/zerolog"
)

func StartConfig() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
