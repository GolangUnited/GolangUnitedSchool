package model

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"regexp"
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
	GroupContactId int64  `json:"group_contact_id" validate:"numeric,gt=0"`
	GroupId        int64  `json:"group_id" validate:"numeric,gt=0"`
	ContactTypeId  int64  `json:"contact_type_id" validate:"numeric,gt=0,lte=7"`
	IsPrimary      bool   `json:"is_primary" validate:"required"`
	ContactValue   string `json:"contact_value" validate:"min=2,max=70"`
}

type GroupContactsListDto struct {
	Metadata      PaginationResponse `json:"_metadata"`
	GroupContacts []GroupContact
}

type GroupContactsAddDto struct {
	GroupId       int64  `json:"group_id" validate:"numeric,gt=0"`
	ContactTypeId int64  `json:"contact_type_id" validate:"numeric,gt=0,lte=7"`
	IsPrimary     bool   `json:"is_primary" validate:"required"`
	ContactValue  string `json:"contact_value" validate:"min=2,max=70"`
}

type GroupContactsUpdateDto struct {
	GroupId       *int64  `json:"group_id" validate:"numeric,gt=0"`
	ContactTypeId *int64  `json:"contact_type_id" validate:"numeric,gt=0,lte=7"`
	IsPrimary     *bool   `json:"is_primary" validate:"required"`
	ContactValue  *string `json:"contact_value" validate:"min=2,max=70"`
}

func (n *GroupContact) ValidateGroupContact() error {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(n)
	errs := translateError(err, trans)
	telegramContactType, _ := regexp.Compile("^@[a-zA-Z_0-9]+$")
	telegramContactGroup, _ := regexp.Compile("^t[.]me/[a-zA-Z_0-9]+$")
	discordContactType, _ := regexp.Compile("^[a-zA-Z_0-9]+#[a-zA-Z_0-9]+$")
	discordChannelContactType, _ := regexp.Compile("^https://discord.com/channels/[0-9]+/[0-9]+$")
	emailContactType, _ := regexp.Compile("^[a-zA-Z_0-9]+@[a-zA-Z_0-9]+[.][a-z]+$")
	phoneContactType, _ := regexp.Compile("\\+[0-9]")
	vkProfileContactType, _ := regexp.Compile("^vk.com/[a-zA-Z_0-9]+$")
	switch n.ContactTypeId {
	case 1:
		if telegramContactValid := telegramContactType.MatchString(n.ContactValue); !telegramContactValid {
			return fmt.Errorf("invalid telegram contact")
		}
	case 2:
		if telegramContactGroup := telegramContactGroup.MatchString(n.ContactValue); !telegramContactGroup {
			return fmt.Errorf("invalid telegram group contact")
		}
	case 3:
		if discordContactType := discordContactType.MatchString(n.ContactValue); !discordContactType {
			return fmt.Errorf("invalid discord contact")
		}
	case 4:
		if discordChannelContactType := discordChannelContactType.MatchString(n.ContactValue); !discordChannelContactType {
			return fmt.Errorf("invalid discord channel contact")
		}
	case 5:
		if emailContactType := emailContactType.MatchString(n.ContactValue); !emailContactType {
			return fmt.Errorf("invalid email contact")
		}
	case 6:
		if phoneContactType := phoneContactType.MatchString(n.ContactValue); !phoneContactType {
			return fmt.Errorf("invalid phone contact")
		}
		if len(n.ContactValue) != 12 {
			return fmt.Errorf("invalid phone contact")
		}
	case 7:
		if vkProfileContactType := vkProfileContactType.MatchString(n.ContactValue); !vkProfileContactType {
			return fmt.Errorf("invalid vk profile")
		}

	}

	if errs != nil {
		return fmt.Errorf("%s", errs)
	}
	return nil

}
