package db

import (
	"github.com/Komdosh/go-bookstore-oauth-api/src/app/clients/cassandra"
	"github.com/Komdosh/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/Komdosh/go-bookstore-oauth-api/src/utils/errors_utils"
)

type DbRepostiory interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepostiory {
	return &dbRepository{}
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors_utils.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors_utils.NewInternalServerError("db connection not implemented yet")
}
