package domain

import "time"

type Person struct {
	ID         int64
	FirstName  string
	LastName   string
	Patronymic string
	Login      string
	RoleId     int
	Passwd     string
	UpdatedAt  time.Time
	Deleted    bool
}