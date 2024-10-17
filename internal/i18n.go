package internal

import (
	"github.com/invopop/ctxi18n"
	"github.com/leedrum/ikarus_travel/locales"
	"github.com/rs/zerolog/log"
)

func InitI18n() {
	if err := ctxi18n.Load(locales.Content); err != nil {
		log.Fatal().Err(err).Msg("Cannot load locales")
	}
}
