package access_token

import (
	"fmt"
	"github.com/Komdosh/go-bookstore-utils/crypto_utils"
	"github.com/Komdosh/go-bookstore-utils/rest_errors"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// grant_type: password
	Username string `json:"username"`
	Password string `json:"password"`
}

func (at *AccessTokenRequest) Validate() rest_errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		at.Username = strings.TrimSpace(at.Username)
		at.Password = strings.TrimSpace(at.Password)
		if at.Username == "" {
			return rest_errors.NewBadRequestError("invalid username")
		}
		if at.Password == "" {
			return rest_errors.NewBadRequestError("invalid password")
		}
		break
	case grantTypeClientCredentials:
		return rest_errors.NewBadRequestError("client_credentials grant type does not supported yet")
	default:
		return rest_errors.NewBadRequestError("invalid grant_type parameter")
	}

	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) Validate() rest_errors.RestErr {
	accessTokenId := strings.TrimSpace(at.AccessToken)
	if len(accessTokenId) == 0 {
		return rest_errors.NewBadRequestError("invalid access token")
	}
	if at.UserId < 0 {
		return rest_errors.NewBadRequestError("invalid userId")
	}
	if at.ClientId < 0 {
		return rest_errors.NewBadRequestError("invalid clientId")
	}
	if at.Expires < 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetSha1(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
