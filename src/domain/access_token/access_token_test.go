package access_token

import (
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.isExpired() {
		t.Error("brand new access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("new access token should not have defined access token id")
	}

	if at.UserId != 0 {
		t.Error("new access token should not have an associated user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.isExpired() {
		t.Error("empty access token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.isExpired() {
		t.Error("access token expiring three hours from now should not be expired")
	}

}
