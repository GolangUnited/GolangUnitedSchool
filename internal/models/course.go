package models

import "time"

type Course struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// course struct to create and update
type CourseCUSt struct {
	ID        int64
	Title     *string
	Status    *string
	StartDate *time.Time
	EndDate   *time.Time
}
