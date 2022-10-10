package domain

type Person struct {
	ID        int64
	FirstName string
	LastName  string
	SurName   string
	Login     string
}

type Response struct {
	IsSuccess bool        `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
