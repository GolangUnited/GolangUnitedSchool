package model

import "time"

type CourseDTO struct {
	Id        int64        `json:"id"`
	Title     string       `json:"title"`
	Status    CourseStatus `json:"status"`
	StartDate time.Time    `json:"start_date"`
	EndDate   time.Time    `json:"end_date"`
}

type Course struct {
	Id        int64
	Title     string    `json:"title"`
	StatusId  int64     `json:"status_id"`
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
	StatusId  *int64     `json:"status"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type CourseList struct {
	Metadata PaginationResponse `json:"_metadata"`
	Courses  []Course           `json:"courses"`
}
