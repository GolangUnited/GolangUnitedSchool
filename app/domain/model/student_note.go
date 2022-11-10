package model

import "time"

type StudentNote struct {
	StudentNoteId     int64     `json:"student_note_id" validate:"required,gte=0"`
	StudentId         int64     `json:"student_id" validate:"required,gte=0"`
	StudentNoteTypeId int64     `json:"student_note_type_id" validate:"required,gte=0"`
	Note              string    `json:"note" validate:"required,gte=2,lte=255"`
	CreatedAt         time.Time `json:"created_at" validate:"omitempty"`
}
