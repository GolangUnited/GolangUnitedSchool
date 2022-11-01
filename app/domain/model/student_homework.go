package model

import "time"

type StudentHomework struct {
	ID         int64
	StudentID  int64
	HomeworkID int64
	StartedAt  time.Time
	FinishedAt time.Time
	Score      float32
}
