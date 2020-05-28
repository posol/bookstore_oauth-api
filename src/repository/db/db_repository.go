package db

import (
	"github.com/posol/bookstore_oauth-api/src/domain/access_token"
	"github.com/posol/bookstore_oauth-api/src/utils/errors"
)

type dbRepository struct {
}

func (r dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestError) {
	return nil, errors.NewIntrenalServerError("database connection not implemented yet")
}

func NewRepository() access_token.Repository {
	return &dbRepository{}
}
