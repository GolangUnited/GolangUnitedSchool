package model

import "time"

type StudentHomework struct {
	ID         int64     `json:"id"`
	StudentId  int64     `json:"student_id"`
	HomeworkId int64     `json:"homework_id"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	Score      float32   `json:"score"`
}
