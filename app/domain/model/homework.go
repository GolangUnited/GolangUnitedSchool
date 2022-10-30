package model

type Homework struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	Task      string  `json:"task"`
	MaxScore  float32 `json:"max_score"`
	LectureId int64   `json:"lecture_id"`
}
