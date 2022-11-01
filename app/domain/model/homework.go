package model

type Homework struct {
	ID        int64
	Title     string
	Task      string
	MaxScore  float32
	LectureID int64
}
