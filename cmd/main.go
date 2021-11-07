package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/twizar/users/internal/adapter"
	"github.com/twizar/users/internal/application/service"
	"github.com/twizar/users/internal/ports"
)

const (
	regularGroupEnv = "REGULAR_USERS_GROUP"
)

func main() {
	group, ok := os.LookupEnv(regularGroupEnv)
	if !ok {
		log.Panicf("required env `%s` is missing", regularGroupEnv)
	}

	cognitoClient := cognitoidentityprovider.New(session.Must(session.NewSession()))
	users := service.NewUsers(group, adapter.NewCognitoIdentityProvider(cognitoClient))
	handler := ports.NewLambdaHandler(users)
	lambda.Start(handler.Handle)
}
