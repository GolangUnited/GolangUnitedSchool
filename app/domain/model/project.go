package model

type Project struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseID    int64  `json:"course_id"`
	GroupID     int64  `json:"group_id"`
}

type ProjectList struct {
	Metadata PaginationResponse `json:"_metadata"`
	Projects []Project          `json:"projects"`
}
