package handler

import (
	"fmt"

	"github.com/fprofit/EffectiveMobile/internal/models"
)

func checkJSONAddUser(addUser models.AddUser) error {
	switch {
	case addUser.Name == nil:
		return fmt.Errorf("Name is a required field")
	case addUser.Surname == nil:
		return fmt.Errorf("Surname is a required field")
	}
	return nil
}
