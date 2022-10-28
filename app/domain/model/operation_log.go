package model

import "time"

type OperationLog struct {
	ID          int
	Description string    `json:"description"`
	OperationID int       `json:"operation_id"`
	RoleID      int       `json:"role_id"`
	PersonID    int       `json:"person_id"`
	CreatedAt   time.Time `json:"created_at"`
}
