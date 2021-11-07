package adapter

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoIdentityProvider struct {
	provider *cognitoidentityprovider.CognitoIdentityProvider
}

func NewCognitoIdentityProvider(provider *cognitoidentityprovider.CognitoIdentityProvider) *CognitoIdentityProvider {
	return &CognitoIdentityProvider{provider: provider}
}

func (cip CognitoIdentityProvider) AddUserToGroup(username, group, usersPoolID string) error {
	_, err := cip.provider.AdminAddUserToGroup(&cognitoidentityprovider.AdminAddUserToGroupInput{
		Username:   aws.String(username),
		GroupName:  aws.String(group),
		UserPoolId: aws.String(usersPoolID),
	})
	if err != nil {
		return fmt.Errorf("adding user to group error: %w", err)
	}

	return nil
}
