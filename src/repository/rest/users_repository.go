package rest

import (
	"github.com/Komdosh/go-bookstore-oauth-api/src/domain/users"
	"github.com/Komdosh/go-bookstore-oauth-api/src/utils/errors_utils"
	"github.com/go-resty/resty"
)

const (
	UsersBaseURL = "http://127.0.0.1:8080"
)

var (
	usersRestClient = resty.New()
)

type HttpUserRepository interface {
	LoginUser(string, string) (*users.User, *errors_utils.RestErr)
}

type usersRepository struct {
}

func NewRepository() HttpUserRepository {
	return &usersRepository{}
}

func (r usersRepository) LoginUser(email string, password string) (*users.User, *errors_utils.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	var restErr *errors_utils.RestErr
	var user users.User
	_, err := usersRestClient.R().SetBody(request).SetError(&restErr).SetResult(&user).Post(UsersBaseURL + "/users/login")
	if err != nil {
		if restErr != nil {
			return nil, restErr
		}
		return nil, errors_utils.NewInternalServerError("error while login user")
	}

	return &user, nil
}
