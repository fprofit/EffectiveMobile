package handler

import (
	"fmt"
	"strings"

	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/fprofit/EffectiveMobile/internal/utils"
)

func checkJSONAddPerson(person models.Person) error {
	switch {
	case person.Name == nil:
		return fmt.Errorf("Name is a required field")
	case person.Surname == nil:
		return fmt.Errorf("Surname is a required field")
	}
	return nil
}

func checkFilterGetPerson(filter models.PersonFilter) error {
	if filter.Gender != nil && !(*filter.Gender == "male" || *filter.Gender == "female") {
		return fmt.Errorf("Invalid value for the 'Gender' field. It should be 'male' or 'female'.")
	}
	if filter.Age != nil && *filter.Age < 1 {
		return fmt.Errorf("Age must be greater than 0.")
	}
	if filter.Country != nil {
		*filter.Country = strings.ToUpper(*filter.Country)
		if err := utils.CheckCountry(*filter.Country); err != nil {
			return err
		}
	}

	if filter.Sort != nil {
		sort := map[string]interface{}{"name": nil, "surname": nil, "patronymic": nil, "age": nil, "gender": nil, "country_id": nil}
		if _, ok := sort[*filter.Sort]; !ok {
			*filter.Sort = "id"
		}
	}

	return nil
}

func checkJSONUpdPerson(person models.EnrichedPerson) error {
	if person.Gender != nil && !(*person.Gender == "male" || *person.Gender == "female") {
		return fmt.Errorf("Invalid value for the 'Gender' field. It should be 'male' or 'female'.")
	}
	if person.Age != nil && *person.Age < 1 {
		return fmt.Errorf("Age must be greater than 0.")
	}
	if person.CountryID != nil {
		*person.CountryID = strings.ToUpper(*person.CountryID)
		return utils.CheckCountry(*person.CountryID)
	}
	return nil
}
