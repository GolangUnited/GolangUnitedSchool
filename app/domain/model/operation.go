package model

type Operation struct {
	ID        int
	Title     string `json:"title"`
	IsLogging bool   `json:"is_logging"`
}
