package model

import "time"

type OperationLog struct {
	Id          int64     `json:"id"`
	Operation   Operation `json:"operation"`
	PersonId    int64     `json:"person_id"`
	RoleId      int64     `json:"role_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type OperationLogList struct {
	Metadata      PaginationResponse `json:"_metadata"`
	OperationLogs []OperationLog     `json:"operation_logs"`
}

type Operation struct {
	Id            int64         `json:"id"`
	OperationType OperationType `json:"operation_type"`
	Title         string        `json:"title"`
	IsLogging     bool          `json:"is_logging"`
}

type OperationList struct {
	Metadata   PaginationResponse `json:"_metadata"`
	Operations []Operation        `json:"operations"`
}

type OperationCreate struct {
	Title           string `json:"title"`
	OperationTypeId *int64 `json:"operation_id"`
	IsLogging       *bool  `json:"is_logging"`
}

type OperationUpdate struct {
	OperationTypeId *int64  `json:"operation_type_id"`
	Title           *string `json:"title"`
	IsLogging       *bool   `json:"is_logging"`
}

type OperationType struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
