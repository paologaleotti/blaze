package main

import (
	"blaze/internal/api"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/rs/zerolog/log"
)

func main() {
	router := api.InitService()

	log.Info().Msg("Server started at http://localhost:3000")
	lambda.Start(httpadapter.New(router).ProxyWithContext)
}
