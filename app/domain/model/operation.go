package model

import "time"

type OperationLog struct {
	ID          int64
	Operation   Operation
	PersonID    int64
	RoleID      int64
	Description string
	CreatedAt   time.Time
}
type Operation struct {
	ID            int64
	OperationType OperationType
	Title         string
	IsLogging     bool
}

type OperationType struct {
	ID    int64
	Title string
}
