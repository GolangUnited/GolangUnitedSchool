package model

type Contact struct {
	Id            int64  `json:"id"`
	PersonId      int64  `json:"person_id"`
	ContactTypeId int64  `json:"contact_type_id"`
	ContactValue  string `json:"contact_value"`
}
