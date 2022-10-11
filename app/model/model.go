package model

type Response struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type Person struct {
	ID        int64
	FirstName string
	LastName  string
	SurName   string
	Login     string
}
