package access_token

import (
	"github.com/Komdosh/go-bookstore-oauth-api/src/utils/errors_utils"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) Validate() *errors_utils.RestErr {
	accessTokenId := strings.TrimSpace(at.AccessToken)
	if len(accessTokenId) == 0 {
		return errors_utils.NewBadRequestError("invalid access token")
	}
	if at.UserId < 0 {
		return errors_utils.NewBadRequestError("invalid userId")
	}
	if at.ClientId < 0 {
		return errors_utils.NewBadRequestError("invalid clientId")
	}
	if at.Expires < 0 {
		return errors_utils.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}
