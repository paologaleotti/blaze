package main

import (
	"blaze/internal/api"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	router := api.InitService()

	log.Info().Msg("Server started at http://localhost:3000")
	err := http.ListenAndServe("0.0.0.0:3000", router)
	if err != nil {
		log.Fatal().Msgf("Could not start http server: %v\n", err)
	}

}
