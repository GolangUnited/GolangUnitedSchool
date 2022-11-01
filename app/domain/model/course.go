package model

import "time"

type Course struct {
	ID        int64
	Title     string
	Status    string
	StartDate time.Time
	EndDate   time.Time
}
