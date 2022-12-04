package postgres

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strings"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetPersons(ctx context.Context) ([]model.Person, error) {
	var persons []model.Person
	rows, err := r.pool.Query(ctx, `SELECT * FROM person`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of persons")
	}

	for rows.Next() {
		var c model.Person
		err := rows.Scan(&c.PersonId,
			&c.FirstName,
			&c.LastName,
			&c.Patronymic,
			&c.Login,
			&c.RoleId,
			&c.Passwd,
			&c.UpdatedAt,
			&c.Deleted,
			&c.CreatedAt,
			&c.Email,
			&c.Birthday,
		)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan person")
		}

		persons = append(persons, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "person rows error")
	}

	return persons, nil

}
func (r *PostgresRepository) GetPersonById(ctx context.Context, id int64) (*model.Person, error) {
	query := `SELECT * FROM person WHERE id=$1`
	var c model.Person
	err := r.pool.QueryRow(ctx, query, id).
		Scan(&c.PersonId,
			&c.FirstName,
			&c.LastName,
			&c.Patronymic,
			&c.Login,
			&c.RoleId,
			&c.Passwd,
			&c.UpdatedAt,
			&c.Deleted,
			&c.CreatedAt,
			&c.Email,
			&c.Birthday)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get person by id")
	}
	return &c, nil
}
func (r *PostgresRepository) AddNewPerson(ctx context.Context, data *model.NewPersonDto) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	if data.FirstName != "" {
		args = append(args, data.FirstName)
		keys = append(keys, "first_name")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}
	if data.LastName != "" {
		args = append(args, data.LastName)
		keys = append(keys, "last_name")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	if data.Patronymic != "" {
		args = append(args, data.Patronymic)
		keys = append(keys, "patronymic")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	if data.Login != "" {
		args = append(args, data.Login)
		keys = append(keys, "login")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	if data.Email != "" {
		args = append(args, data.Email)
		keys = append(keys, "email")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	if data.Birthday != "" {
		args = append(args, data.Birthday)
		keys = append(keys, "birthday")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	if data.RoleId != 0 {
		args = append(args, data.RoleId)
		keys = append(keys, "role_id")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}
	if data.Passwd != "" {
		args = append(args, data.Passwd)
		keys = append(keys, "passwd")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	query := fmt.Sprintf(
		`INSERT INTO person (%s) VALUES (%s) RETURNING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)

	r.lg.Info("query to add new person: ", query)
	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "couldn't add new person")
	}

	return id, nil
}
func (r *PostgresRepository) UpdatePersonByID(ctx context.Context, id int64, data *model.UpdatePerson) error {
	var args []interface{}
	var keys []string

	// title
	if data.FirstName != nil {
		args = append(args, data.FirstName)
		keys = append(keys, fmt.Sprintf("first_name = $%d", len(args)))
	}

	// status
	if data.LastName != nil {
		args = append(args, data.LastName)
		keys = append(keys, fmt.Sprintf("last_name = $%d", len(args)))
	}

	// start date
	if data.Patronymic != nil {
		args = append(args, data.Patronymic)
		keys = append(keys, fmt.Sprintf("patronymic = $%d", len(args)))
	}

	// end date
	if data.Login != nil {
		args = append(args, data.Login)
		keys = append(keys, fmt.Sprintf("login = $%d", len(args)))
	}

	if data.RoleId != nil {
		args = append(args, data.RoleId)
		keys = append(keys, fmt.Sprintf("role_id = $%d", len(args)))
	}

	if data.Passwd != nil {
		args = append(args, data.Passwd)
		keys = append(keys, fmt.Sprintf("passwd = $%d", len(args)))
	}

	if data.UpdatedAt != nil {
		args = append(args, data.UpdatedAt)
		keys = append(keys, fmt.Sprintf("updated_at = $%d", len(args)))
	}

	if data.Email != nil {
		args = append(args, data.Email)
		keys = append(keys, fmt.Sprintf("email = $%d", len(args)))
	}

	if data.Birthday != nil {
		args = append(args, data.Birthday)
		keys = append(keys, fmt.Sprintf("birthday = $%d", len(args)))
	}

	// id
	args = append(args, id)

	query := fmt.Sprintf(`UPDATE person 
				SET %s
				WHERE id = $%d`, strings.Join(keys, ","), len(args))

	r.lg.Debug("query:", query)
	cmn, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update person by id %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("person doesn't exist"),
			"couldn't update person by id %d", id)
	}
	return nil
}

func (r *PostgresRepository) PutPersonByID(ctx context.Context, id int64, data *model.UpdatePerson) error {
	query := `UPDATE person 
				SET 
					first_name = $1, 
					last_name = $2, 
					patronymic = $3, 
					login = $4,
					email = $5,
					birthday = $6,
					role_id = $7,
					passwd = $8,
					updated_at = $9,
					deleted = $10
				WHERE id=$11`

	cmn, err := r.pool.Exec(ctx, query,
		data.FirstName,
		data.LastName,
		data.Patronymic,
		data.Login,
		data.Email,
		data.Birthday,
		data.RoleId,
		data.Passwd,
		data.UpdatedAt,
		data.Deleted,
		id,
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't put person by id: %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("person doesn't exist"),
			"couldn't put person by id: %d", id)
	}

	return nil
}

func (r *PostgresRepository) DeletePersonById(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM person WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete person by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("person doesn't exist"),
			"couldn't delete person by id %d", id)
	}
	return nil
}
