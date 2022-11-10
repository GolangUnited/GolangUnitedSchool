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

func (n *GroupContact) ValidateGroupContact() error {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(n)
	errs := translateError(err, trans)
	telegramContactType, _ := regexp.Compile("^@[a-zA-Z_0-9]+$")
	telegramContactGroup, _ := regexp.Compile("")
	discordContactType, _ := regexp.Compile("")
	discordChannelContactType, _ := regexp.Compile("")
	emailContactType, _ := regexp.Compile("")
	phoneContactType, _ := regexp.Compile("")
	vkProfileContactType, _ := regexp.Compile("")
	switch n.ContactTypeId {
	case 1:
		if telegramContactValid := telegramContactType.MatchString(n.ContactValue); !telegramContactValid {
			_ = fmt.Errorf("invalid telegram contact")
		}
	case 2:
		if telegramContactGroup := telegramContactGroup.MatchString(n.ContactValue); !telegramContactGroup {
			_ = fmt.Errorf("invalid telegram group contact")
		}
	case 3:
		if discordContactType := discordContactType.MatchString(n.ContactValue); !discordContactType {
			_ = fmt.Errorf("invalid discord contact")
		}
	case 4:
		if discordChannelContactType := discordChannelContactType.MatchString(n.ContactValue); !discordChannelContactType {
			_ = fmt.Errorf("invalid discord channel contact")
		}
	case 5:
		if emailContactType := emailContactType.MatchString(n.ContactValue); !emailContactType {
			_ = fmt.Errorf("invalid email contact")
		}
	case 6:
		if phoneContactType := phoneContactType.MatchString(n.ContactValue); !phoneContactType {
			_ = fmt.Errorf("invalid phone contact")
		}
	case 7:
		if vkProfileContactType := vkProfileContactType.MatchString(n.ContactValue); !vkProfileContactType {
			_ = fmt.Errorf("invalid vk profile")
		}

	}

	if errs != nil {
		return fmt.Errorf("%s", errs)
	}
	return nil

}
