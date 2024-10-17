package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := internal.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := internal.Server{
		Config: config,
	}
	server.Run()
}
