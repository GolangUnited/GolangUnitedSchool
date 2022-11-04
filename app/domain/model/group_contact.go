package model

type GroupContact struct {
	GroupContactId int64
	GroupId        int64
	ContactTypeId  int64
	IdPrimary      bool
	ContactValue   string
}
