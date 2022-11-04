package model

import "time"

type MentorNote struct {
	MentorNoteId int64
	StudentId    int64
	MentorId     int64
	Note         string
	CreatedAt    time.Time
}
