package model

import "time"

type Course struct {
	ID        int64
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type CourseCreate struct {
	Title     string     `json:"title" bind:"required"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type CourseUpdate struct {
	Title     *string    `json:"title"`
	Status    *string    `json:"status"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type CourseList struct {
	Metadata PaginationResponse `json:"_metadata"`
	Courses  []Course           `json:"courses"`
}
