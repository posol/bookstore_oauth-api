package db

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/posol/bookstore_oauth-api/clients/cassandra"
	"github.com/posol/bookstore_oauth-api/src/domain/access_token"
	"github.com/posol/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "select access_token, user_id, client_id, expires from access_tokens where access_token = ?;"
	queryCreateAccessToken = "insert into access_tokens(access_token, user_id, client_id, expires) values(?,?,?,?);"
	queryUpdateExpires     = "update access_token set expires = ? where access_token = ?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewIntrenalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		fmt.Println(err)
		if err == gocql.ErrNotFound {
			return nil, errors.NewIntrenalServerError("no access token found with given id")
		}
		return nil, errors.NewIntrenalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewIntrenalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		token.AccessToken,
		token.UserId,
		token.ClientId,
		token.Expires,
	).Exec(); err != nil {
		return errors.NewIntrenalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewIntrenalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateExpires,
		token.Expires,
		token.AccessToken,
	).Exec(); err != nil {
		return errors.NewIntrenalServerError(err.Error())
	}
	return nil
}
