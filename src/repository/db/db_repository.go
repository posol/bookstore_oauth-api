package db

import (
	"github.com/posol/bookstore_oauth-api/clients/cassandra"
	"github.com/posol/bookstore_oauth-api/src/domain/access_token"
	"github.com/posol/bookstore_oauth-api/src/utils/errors"
	"log"
)

type dbRepository struct {
}

func (r dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestError) {
	session, err := cassandra.GetSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	return nil, errors.NewIntrenalServerError("database connection not implemented yet")
}

func NewRepository() access_token.Repository {
	return &dbRepository{}
}
