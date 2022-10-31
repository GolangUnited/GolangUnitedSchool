package model

import "time"

type Mentor struct {
	ID int64
}

type MentorNote struct {
	ID        int64
	StudentID int64
	MentorID  int64
	Note      string
	CreatedAt time.Time
}
