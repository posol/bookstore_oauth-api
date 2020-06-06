package rest

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/api/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"posol2008@gmail.com","password":"123"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("posol2008@gmail.com", "123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/api/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"posol2008@gmail.com","password":"123"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials","code": "404","error": "not_found"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("posol2008@gmail.com", "123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/api/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"posol2008@gmail.com","password":"123"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials","code": 404,"error": "not_found"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("posol2008@gmail.com", "123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/api/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"posol2008@gmail.com","password":"123"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
  			"id": "25"",
			"first_name": "Aleksandr",
  			"last_name": "Poslovskii",
  			"email": "posol2008@gmail.com"
		}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("posol2008@gmail.com", "123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal users login response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/api/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"posol2008@gmail.com","password":"123"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
  			"id": 25,
			"first_name": "Aleksandr",
  			"last_name": "Poslovskii",
  			"email": "posol2008@gmail.com"
		}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("posol2008@gmail.com", "123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 25, user.Id)
	assert.EqualValues(t, "Aleksandr", user.FirstName)
	assert.EqualValues(t, "Poslovskii", user.LastName)
	assert.EqualValues(t, "posol2008@gmail.com", user.Email)
}
