package postgres

import "time"

type DBPerson struct {
	ID         int64
	FirstName  *string
	LastName   *string
	Patronymic *string
	Login      *string
	RoleId     int64
	Passwd     *string
	UpdatedAt  time.Time
	Deleted    bool
}
