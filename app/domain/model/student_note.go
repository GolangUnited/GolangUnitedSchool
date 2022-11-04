package model

import "time"

type StudentNote struct {
	StudentNoteId     int64
	StudentId         int64
	StudentNoteTypeId int64
	Note              string
	CreatedAt         time.Time
}
