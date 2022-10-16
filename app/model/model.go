package model

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	ID         int64     `json:"id"`
	FirstName  string    `json:"first_name" validate:"required,min=2,max=50,alphaunicode"`
	LastName   string    `json:"last_name" validate:"required,min=2,max=50,alphaunicode"`
	Patronymic string    `json:"patronymic" validate:"omitempty,min=2,max=50,alphaunicode"`
	Login      string    `json:"login" validate:"required,alphanum"`
	RoleId     int64     `json:"role_id" validate:"required,numeric,gte=0,lte=6"`
	Passwd     string    `json:"passwd" validate:"required,ascii,min=8,max=20"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (p *Person) Validate() error {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return fmt.Errorf("Validate: %w", err)
	}

	return nil
}
