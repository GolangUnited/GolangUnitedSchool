package model

type Homework struct {
	ID        int64
	LectureID int64
	Title     string
	Task      string
	MaxScore  int
}
