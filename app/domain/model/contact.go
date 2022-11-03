package model

type Contact struct {
	ID            int64
	PersonID      int64
	ContactTypeID int64
	IsPrimary     bool
	ContactValue  string
}

type ContactType struct {
	ID    int64
	Title string
}
