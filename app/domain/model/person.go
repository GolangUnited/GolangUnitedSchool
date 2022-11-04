package model

import "time"

type Person struct {
	PersonId   int64
	FirstName  string    `json:"first_name" validate:"required,min=2,max=50"`
	LastName   string    `json:"last_name" validate:"required,min=2,max=50"`
	Patronymic string    `json:"patronymic" validate:"omitempty,min=2,max=50"`
	Login      string    `json:"login" validate:"required,ascii"`
	RoleId     int       `json:"role_id" validate:"required,numeric,gte=0,lte=6"`
	Passwd     string    `json:"passwd" validate:"required,ascii,min=8,max=20"`
	UpdatedAt  time.Time `json:"updated_at" validate:"omitempty"`
	Deleted    bool      `json:"deleted" validate:"omitempty"`
}
