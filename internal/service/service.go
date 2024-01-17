package service

import (
	// "fmt"

	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/gin-gonic/gin"
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

func (s *Service) AddUser(c *gin.Context, addUser models.AddUser) (models.ResponseUser, error) {
	age, err := s.fetchAgeFromAPI(*addUser.Name)
	if err != nil {
		return models.ResponseUser{}, err
	}
	countryID, err := s.fetchCountryIDFromAPI(*addUser.Name)
	if err != nil {
		return models.ResponseUser{}, err
	}
	gender, err := s.fetchGenderFromAPI(*addUser.Name)
	if err != nil {
		return models.ResponseUser{}, err
	}
	t := models.ResponseUser{Name: *addUser.Name, Surname: *addUser.Surname, Patronymic: *addUser.Patronymic, Age: age, CountryID: countryID, Gender: gender}
	return t, nil
}

func (s *Service) DelUser(c *gin.Context, id int64) error {

	return nil
}
