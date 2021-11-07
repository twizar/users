package ports

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/twizar/users/internal/application/service"
)

type LambdaHandler struct {
	users *service.Users
}

func NewLambdaHandler(usersService *service.Users) *LambdaHandler {
	return &LambdaHandler{users: usersService}
}

func (lh LambdaHandler) Handle(event *events.CognitoEventUserPoolsPostConfirmation) (*events.CognitoEventUserPoolsPostConfirmation, error) {
	err := lh.users.AddUserToRegularGroup(event.UserName, event.UserPoolID)
	if err != nil {
		return nil, fmt.Errorf("adding user to regular group error: %w", err)
	}

	return event, nil
}
