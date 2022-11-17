package model

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type CourseStatus struct {
	CourseStatusId int64  `json:"course_status_id" validate:"numeric,gt=0"`
	Title          string `json:"title" validate:"min=2,max=20,alphaunicode"`
}

type CourseStatusesListDto struct {
	Metadata         PaginationResponse `json:"_metadata"`
	CourseStatusList []CourseStatus
}

func (n *CourseStatus) ValidateCourseStatus() error {
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
