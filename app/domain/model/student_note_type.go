package model

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

/*
1 - опыт работы
2 - самопрезентация
3 - пройденные курсы
4 - что-то еще
*/

type StudentNoteType struct {
	StudentNoteTypeId int64  `json:"student_note_type_id" validate:"numeric,gt=0"`
	Title             string `json:"title" validate:"gte=2,lte=10"`
}

type StudentNoteTypesListDto struct {
	StudentNoteTypeList []StudentNoteType
}

type NewStudentNoteTypeDto struct {
	Title string `json:"title" validate:"gte=2,lte=10"`
}

type UpdateStudentNoteTypeDto struct {
	Title *string `json:"title" validate:"gte=2,lte=10"`
}

func (n *StudentNoteType) ValidateStudentNoteType() error {
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
