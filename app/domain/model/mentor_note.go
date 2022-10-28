package model

import "time"

type MentorNote struct {
	StudentId int64
	MentorId  int64
	Note      string
	CreatedAt time.Time
}
