package access_token

import (
	"github.com/posol/bookstore_oauth-api/src/utils/errors"
	"strings"
)

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type service struct {
	repository Repository
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(token AccessToken) *errors.RestError {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.Create(token)
}

func (s *service) UpdateExpirationTime(token AccessToken) *errors.RestError {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(token)
}
