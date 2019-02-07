package presenter

import "github.com/tomocy/archs/domain/model"

type AuthenticationPresenter interface {
	PresentAuthenticUser(user *model.User) *AuthenticUserPresent
	PresentAuthenticUserID(id model.UserID) string
}

func NewAuthenticationPresenter() AuthenticationPresenter {
	return new(authenticationPresenter)
}

type authenticationPresenter struct {
}

type AuthenticUserPresent struct {
	ID    string
	Email string
}

func (p authenticationPresenter) PresentAuthenticUser(user *model.User) *AuthenticUserPresent {
	return &AuthenticUserPresent{
		ID:    string(user.ID),
		Email: user.Email,
	}
}

func (p authenticationPresenter) PresentAuthenticUserID(id model.UserID) string {
	return string(id)
}