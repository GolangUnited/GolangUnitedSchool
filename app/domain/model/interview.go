package model

import "time"

type Interview struct {
	ID        int64
	StudentID int64
	MentorID  int64
	Note      string
	Score     int
	CreatedAt time.Time
}
