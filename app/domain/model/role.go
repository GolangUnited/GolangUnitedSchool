package model

type Role struct {
	ID         int64
	RoleName   string
	IsReadOnly bool
}

type RoleCU struct {
	RoleName   *string
	IsReadOnly *bool
}
