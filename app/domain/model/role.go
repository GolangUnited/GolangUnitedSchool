package model

type Role struct {
	Id         int64
	RoleName   string
	IsReadOnly bool
}

type RoleCU struct {
	RoleName   *string
	IsReadOnly *bool
}
