package model

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"time"
)

type StudentNote struct {
	StudentNoteId     int64     `json:"student_note_id" validate:"required,gte=1"`
	StudentId         int64     `json:"student_id" validate:"required,gt=0"`
	StudentNoteTypeId int64     `json:"student_note_type_id" validate:"required,gt=0"`
	Note              string    `json:"note" validate:"required,gte=2,lte=255"`
	CreatedAt         time.Time `json:"created_at" validate:"omitempty"`
}

func (n *StudentNote) ValidateStudentNote() error {
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
	return nil

}
