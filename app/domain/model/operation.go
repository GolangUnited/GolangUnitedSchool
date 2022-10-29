package model

import "time"

type Operation struct {
	ID              int64
	OperationTypeID int64
	Title           string
	IsLogging       bool
}

type OperationLog struct {
	ID          int64
	OperationID int64
	PersonID    int64
	RoleID      int64
	Description string
	CreatedAt   time.Time
}

type OperationType struct {
	ID    int64
	Title string
}
