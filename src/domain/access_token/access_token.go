package access_token

import (
	"github.com/posol/bookstore_oauth-api/src/utils/errors"
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

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (a AccessToken) isExpired() bool {
	return time.Unix(a.Expires, 0).Before(time.Now().UTC())
}
