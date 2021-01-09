package db

import (
	"github.com/Komdosh/go-bookstore-oauth-api/src/app/clients/cassandra"
	"github.com/Komdosh/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/Komdosh/go-bookstore-oauth-api/src/utils/errors_utils"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM oauth.access_token WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO oauth.access_token(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE oauth.access_token SET expires=? WHERE access_token=?;"
)

type DbRepostiory interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestErr)
	Create(access_token.AccessToken) *errors_utils.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors_utils.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepostiory {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors_utils.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors_utils.NewNotFoundError("no access token found for given id")
		}
		return nil, errors_utils.NewInternalServerError(err.Error())
	}

	return &result, nil
}
func (r *dbRepository) Create(at access_token.AccessToken) *errors_utils.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors_utils.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}

	return nil
}
