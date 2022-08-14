package rules

import (
	"errors"
	"github.com/devazizi/go-crm/infrastructure"
	"github.com/devazizi/go-crm/repository"
	validation "github.com/go-ozzo/ozzo-validation"
)

func CheckEmailMustUnique(DB infrastructure.DB) validation.RuleFunc {

	return func(value interface{}) error {
		email, _ := value.(string)

		emailIsUnique := repository.CheckEmailIsUnique(DB, email)

		if emailIsUnique == false {
			return errors.New("another user registered with current email")
		}

		return nil
	}

}
