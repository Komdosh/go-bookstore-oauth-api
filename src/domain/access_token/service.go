package access_token

import (
	"github.com/Komdosh/go-bookstore-oauth-api/src/utils/errors_utils"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors_utils.RestErr)
	Create(AccessToken) *errors_utils.RestErr
	UpdateExpirationTime(AccessToken) *errors_utils.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors_utils.RestErr)
	Create(AccessToken) *errors_utils.RestErr
	UpdateExpirationTime(AccessToken) *errors_utils.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors_utils.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors_utils.NewBadRequestError("invalid access token")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors_utils.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors_utils.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
