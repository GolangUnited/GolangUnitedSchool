package domain

import "time"

type Person struct {
	ID        int64
	FirstName string
	LastName  string
	SurName   string
	Login     string
}

type Course struct {
	ID        int64
	Title     string
	Status    string
	StartDate time.Time
	EndDate   time.Time
}
