package model

type Lecture struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type LectureList struct {
	Metadata PaginationResponse `json:"_metadata"`
	Lectures []Lecture          `json:"lectures"`
}
