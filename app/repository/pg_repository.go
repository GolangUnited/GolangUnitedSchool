package repository

type PostgreSQLRepository struct {
	connectionString string
}

func (r *PostgreSQLRepository) GetPersonById(id int64) DBPerson {
	return DBPerson{
		ID:        id,
		FirstName: "test",
		LastName:  "test",
		SurName:   "test",
	}
}

func NewPostgreSQLRepository(connectionString string) *PostgreSQLRepository {
	return &PostgreSQLRepository{
		connectionString: connectionString,
	}
}
