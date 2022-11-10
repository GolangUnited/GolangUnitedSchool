package model

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

/*
типы контактов
- 1 - телеграм  @username
- 2 - телеграм группа t.me/go_united_chat
- 3 - дискорд   Kek#3873
- 4 - дискорд канал https://discord.com/channels/516715744646660106/1012067296744841246
- 5 - почта  kek@gmail.com
- 6 - телефон +79080490909
- 7 - вк  vk.com/funeral
*/

type GroupContact struct {
	GroupContactId int64  `json:"group_contact_id" validate:"required,numeric,gte=0"`
	GroupId        int64  `json:"group_id" validate:"required,numeric,gte=0"`
	ContactTypeId  int64  `json:"contact_type_id" validate:"required,numeric,gte=0,lte=7"`
	IsPrimary      bool   `json:"is_primary" validate:"required"`
	ContactValue   string `json:"contact_value" validate:"required,min=2,max=70"`
}

func (n *GroupContact) ValidateGroupContact() error {
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
