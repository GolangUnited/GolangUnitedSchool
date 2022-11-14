package model

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"regexp"
	"time"
)

type Person struct {
	PersonId   int64     `json:"person_id" validate:"omitempty,numeric,gt=0"`
	FirstName  string    `json:"first_name" validate:"required,min=2,max=50"`
	LastName   string    `json:"last_name" validate:"required,min=2,max=50"`
	Patronymic string    `json:"patronymic" validate:"omitempty,min=2,max=50"`
	Login      string    `json:"login" validate:"required,ascii,min=2,max=50"`
	RoleId     int       `json:"role_id" validate:"required,numeric,gte=1,lte=6"`
	Passwd     string    `json:"passwd" validate:"required,ascii,min=8,max=20"`
	UpdatedAt  time.Time `json:"updated_at" validate:"omitempty"`
	Deleted    bool      `json:"deleted" validate:"omitempty"`
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf("%s", e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	if len(errs) != 0 {
		fmt.Println(len(errs))
		return errs
	}
	return nil

}

func (n *Person) ValidatePerson() error {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(n)
	errs := translateError(err, trans)
	if errs != nil {
		return fmt.Errorf("%s", errs)
	}
	reg, _ := regexp.Compile("^[а-яА-Яa-zA-Z-]+$")
	firstNameValid := reg.MatchString(n.FirstName)
	lastNameValid := reg.MatchString(n.LastName)
	patReg, _ := regexp.Compile("(^[a-zA-Zа-яА-Я]+$|^$)")
	patronymicValid := patReg.MatchString(n.Patronymic)
	if !firstNameValid {
		return fmt.Errorf("invalid chars in firstname field")
	}
	if !lastNameValid {
		return fmt.Errorf("invalid chars in lastname field")
	}
	if !patronymicValid {
		return fmt.Errorf("invalid chars in patronymic field")
	}

	return nil

}
