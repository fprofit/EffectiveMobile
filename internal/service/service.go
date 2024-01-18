package service

import (
	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/sirupsen/logrus"
)

type repository interface {
	GetPersons(filter models.PersonFilter) (models.PersonList, error)
	GetPersonByID(id int64) (models.EnrichedPerson, error)
	AddPerson(person models.EnrichedPerson) (models.EnrichedPerson, error)
	DelPerson(id int64) error
	UpdPerson(person models.EnrichedPerson) (models.EnrichedPerson, error)
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

func (s *Service) AddPerson(person models.Person) (models.EnrichedPerson, error) {
	s.log.Debug("Service AddPerson")
	age, err := s.fetchAgeFromAPI(*person.Name)
	if err != nil {
		return models.EnrichedPerson{}, err
	}
	countryID, err := s.fetchCountryIDFromAPI(*person.Name)
	if err != nil {
		return models.EnrichedPerson{}, err
	}
	gender, err := s.fetchGenderFromAPI(*person.Name)
	if err != nil {
		return models.EnrichedPerson{}, err
	}
	enrichedPerson := models.EnrichedPerson{Name: person.Name, Surname: person.Surname, Patronymic: person.Patronymic, Age: &age, CountryID: &countryID, Gender: &gender}
	return s.repository.AddPerson(enrichedPerson)
}

func (s *Service) UpdPerson(person models.EnrichedPerson) (models.EnrichedPerson, error) {
	s.log.Debug("Service UpdPerson")

	existingPerson, err := s.repository.GetPersonByID(person.ID)
	if err != nil {
		return existingPerson, err
	}
	flag := true
	s.log.Debug("Updating person data...")
	if person.Name != nil && *person.Name != *existingPerson.Name {
		s.log.Debugf("Updating Name: %s -> %s", *existingPerson.Name, *person.Name)
		existingPerson.Name = person.Name
		flag = false
	}
	if person.Surname != nil && *person.Surname != *existingPerson.Surname {
		s.log.Debugf("Updating Surname: %s -> %s", *existingPerson.Surname, *person.Surname)
		existingPerson.Surname = person.Surname
		flag = false
	}
	if person.Patronymic != nil && *person.Patronymic != *existingPerson.Patronymic {
		s.log.Debugf("Updating Patronymic: %s -> %s", *existingPerson.Patronymic, *person.Patronymic)
		existingPerson.Patronymic = person.Patronymic
		flag = false
	}
	if person.Age != nil && *person.Age != *existingPerson.Age {
		s.log.Debugf("Updating Age: %d -> %d", *existingPerson.Age, *person.Age)
		existingPerson.Age = person.Age
		flag = false
	}
	if person.Gender != nil && *person.Gender != *existingPerson.Gender {
		s.log.Debugf("Updating Gender: %s -> %s", *existingPerson.Gender, *person.Gender)
		existingPerson.Gender = person.Gender
		flag = false
	}
	if person.CountryID != nil && *person.CountryID != *existingPerson.CountryID {
		s.log.Debugf("Updating CountryID: %s -> %s", *existingPerson.CountryID, *person.CountryID)
		existingPerson.CountryID = person.CountryID
		flag = false
	}
	if flag {
		return existingPerson, nil
	}

	return s.repository.UpdPerson(existingPerson)
}

func (s *Service) DelPerson(id int64) error {
	s.log.Debug("Service DelPerson")
	return s.repository.DelPerson(id)
}

func (s *Service) GetPersonsByFilter(filter models.PersonFilter) (models.PersonList, error) {
	s.log.Debug("Service GetPersonsByFilter")
	return s.repository.GetPersons(filter)
}
