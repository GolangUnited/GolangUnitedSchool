package model

type Group struct {
	ID       int64
	CourseId int64
	MentorId int64
	Title    string
}

// group create and update struct
type GroupCU struct {
	MentorId int64
	Title    string
}
