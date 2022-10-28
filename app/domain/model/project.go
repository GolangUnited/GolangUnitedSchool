package model

type Project struct {
	ID          int
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseID    int    `json:"course_id"`
	GroupID     int    `json:"group_id"`
}
