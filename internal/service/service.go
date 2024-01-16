package service

import (
	"github.com/sirupsen/logrus"
)

type repository interface {
}

type Service struct {
	repository repository
	apiUrl     ApiUrl
	log        *logrus.Logger
}

type ApiUrl struct {
	ApiGetAge     string
	ApiGetCountry string
	ApiGetGender  string
}

func NewService(repository repository, apiUrl ApiUrl, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		apiUrl:     apiUrl,
		log:        log,
	}
}
