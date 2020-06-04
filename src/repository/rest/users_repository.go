package rest

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/posol/bookstore_oauth-api/src/domain/users"
	"github.com/posol/bookstore_oauth-api/src/utils/errors"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type usersRepository struct {
}

func (u *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestError) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("api/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.NewIntrenalServerError("invalid rest client response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewIntrenalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewIntrenalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}
