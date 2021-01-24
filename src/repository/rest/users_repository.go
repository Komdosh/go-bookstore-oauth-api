package rest

import (
	"github.com/Komdosh/go-bookstore-oauth-api/src/domain/users"
	"github.com/Komdosh/go-bookstore-utils/rest_errors"

	"github.com/go-resty/resty"
)

const (
	usersBaseURL = "http://127.0.0.1:8081"
)

var (
	usersRestClient = resty.New()
)

type HttpUserRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() HttpUserRepository {
	return &usersRepository{}
}

func (r usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	var restErr rest_errors.RestErr
	var user users.User
	_, err := usersRestClient.R().SetBody(request).SetError(&restErr).SetResult(&user).Post(usersBaseURL + "/users/login")
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error while login user", err)
	}
	if restErr != nil {
		return nil, restErr
	}

	return &user, nil
}
