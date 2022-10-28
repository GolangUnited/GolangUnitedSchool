package model

import "time"

type Course struct {
	CourseId  int64
	Title     string    `json:"title" bind:""`
	Status    string    `json:"status"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
