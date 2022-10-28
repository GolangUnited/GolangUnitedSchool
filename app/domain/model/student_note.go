package model

import "time"

type StudentNote struct {
	StudentId         int64
	StudentNoteTypeId int64
	Note              string
	CreatedAt         time.Time
}
