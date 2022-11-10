package model

import "time"

type MentorNote struct {
	MentorNoteId int64     `json:"mentor_note_id" validate:"required,numeric,gt=0"`
	StudentId    int64     `json:"student_id" validate:"required,numeric,gt=0"`
	MentorId     int64     `json:"mentor_id" validate:"required,numeric,gt=0"`
	Note         string    `json:"note" validate:"required,gte=2,lte=255"`
	CreatedAt    time.Time `json:"created_at" validate:"omitempty"`
}
