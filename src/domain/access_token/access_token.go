package access_token

import (
	"fmt"
	"github.com/posol/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/posol/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:scope`

	// used for password grant type
	UserName string `json:"username"`
	Password string `json:"password"`

	// used for client credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestError {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	// TODO: validate parametrs for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (a *AccessToken) Validate() *errors.RestError {
	a.AccessToken = strings.TrimSpace(a.AccessToken)
	if a.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if a.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if a.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if a.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (a AccessToken) isExpired() bool {
	return time.Unix(a.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
