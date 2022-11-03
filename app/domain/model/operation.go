package model

import "time"

type LogOperation struct {
	ID          int64
	OperationId int64
	PersonID    int64
	RoleID      int64
	Description string
	CreatedAt   time.Time
}
type Operation struct {
	ID              int64
	OperationTypeID int64
	Title           string
	IsLogging       bool
}

type OperationType struct {
	ID    int64
	Title string
}
