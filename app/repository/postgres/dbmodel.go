package postgres

type DBPerson struct {
	ID        int64
	FirstName *string
	LastName  *string
	SurName   *string
	Login     *string
}
