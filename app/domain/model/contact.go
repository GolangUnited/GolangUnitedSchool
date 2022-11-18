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
