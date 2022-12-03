package model

import "time"

type OperationLog struct {
	Id          int64
	Operation   Operation
	PersonId    int64
	RoleId      int64
	Description string
	CreatedAt   time.Time
}
type Operation struct {
	Id            int64
	OperationType OperationType
	Title         string
	IsLogging     bool
}

type OperationType struct {
	Id    int64
	Title string
}
