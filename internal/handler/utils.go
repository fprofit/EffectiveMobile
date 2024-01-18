package handler

import (
	"fmt"

	"github.com/fprofit/EffectiveMobile/internal/models"
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
