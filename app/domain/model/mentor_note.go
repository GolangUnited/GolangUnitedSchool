package model

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"time"
)

// @Description
type MentorNote struct {
	MentorNoteId int64     `json:"mentor_note_id" validate:"numeric,gt=0"`
	StudentId    int64     `json:"student_id" validate:"numeric,gt=0"`
	MentorId     int64     `json:"mentor_id" validate:"numeric,gt=0"`
	Note         string    `json:"note" validate:"gte=2,lte=255"`
	CreatedAt    time.Time `json:"created_at" validate:"omitempty"`
}

type NewMentorNote struct {
	StudentId int64     `json:"student_id" validate:"numeric,gt=0"`
	MentorId  int64     `json:"mentor_id" validate:"numeric,gt=0"`
	Note      string    `json:"note" validate:"gte=2,lte=255"`
	CreatedAt time.Time `json:"created_at" validate:"omitempty"`
}

type UpdateMentorNote struct {
	StudentId *int64     `json:"student_id" validate:"numeric,gt=0"`
	MentorId  *int64     `json:"mentor_id" validate:"numeric,gt=0"`
	Note      *string    `json:"note" validate:"gte=2,lte=255"`
	CreatedAt *time.Time `json:"created_at" validate:"omitempty"`
}

func (n *MentorNote) ValidateMentorNote() error {
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
