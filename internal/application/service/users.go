package service

import "fmt"

type identityProvider interface {
	AddUserToGroup(username, group, usersPoolID string) error
}

type Users struct {
	regularUsersGroup string
	identityProvider  identityProvider
}

func NewUsers(regularUsersGroup string, identityProvider identityProvider) *Users {
	return &Users{
		regularUsersGroup: regularUsersGroup,
		identityProvider:  identityProvider,
	}
}

func (pc Users) AddUserToRegularGroup(username, usersPoolID string) error {
	err := pc.identityProvider.AddUserToGroup(username, pc.regularUsersGroup, usersPoolID)
	if err != nil {
		return fmt.Errorf("adding user to identity provider group error: %w", err)
	}

	return nil
}
