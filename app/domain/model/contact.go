package model

type Contact struct {
	ID            int64
	PersonID      int64
	ContactTypeID int64
	IsPrimary     bool
	ContactValue  string
}

type ContactsListDto struct {
	Metadata     PaginationResponse
	ContactsList []Contact
}

type ContactType struct {
	ID    int64
	Title string
}

type ContactTypesListDto struct {
	Metadata         PaginationResponse
	ContactTypesList []ContactType
}

type NewContactTypeDto struct {
	Title string
}
