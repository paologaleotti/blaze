package main

import (
	"blaze/internal/todoservice"
	"blaze/pkg/util"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	router := todoservice.InitService()

	util.Log.Info("Server started at http://localhost:3000")
	lambda.Start(httpadapter.New(router).ProxyWithContext)
}
