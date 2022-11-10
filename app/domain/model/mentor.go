package model

type Mentor struct {
	MentorId int64 `json:"mentor_id" validate:"required,gt=0"`
}
