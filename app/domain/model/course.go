package model

import "time"

type Course struct {
	ID        int64
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
