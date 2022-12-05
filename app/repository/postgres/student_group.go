package postgres

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
	"log"
	"strings"
)

func (r *PostgresRepository) GetStudentGroups(ctx context.Context) ([]model.StudentGroup, error) {
	var groups []model.StudentGroup
	rows, err := r.pool.Query(ctx, `SELECT id, student_id, group_id FROM student_groups`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get student groups")
	}

	for rows.Next() {
		var c model.StudentGroup
		err := rows.Scan(&c.Id, &c.StudentId, &c.GroupId)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan student groups")
		}

		groups = append(groups, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "course statuses rows error")
	}

	return groups, nil
}
func (r *PostgresRepository) GetStudentGroupById(ctx context.Context, id int64) (*model.StudentGroup, error) {
	query := `SELECT student_id, group_id, id FROM student_groups WHERE id=$1`
	var c model.StudentGroup
	err := r.pool.QueryRow(ctx, query, id).
		Scan(&c.StudentId, &c.GroupId, &c.Id)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get student group by id")
	}
	return &c, nil
}
func (r *PostgresRepository) CreateStudentGroup(ctx context.Context, data *model.StudentGroup) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	// add title
	args = append(args, data.GroupId)
	keys = append(keys, "group_id")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	// add start date

	args = append(args, data.StudentId)
	keys = append(keys, "student_id")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	query := fmt.Sprintf(
		`INSERT INTO student_groups (%s) VALUES (%s) RETURNING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)
	log.Print(query)

	r.lg.Debug("query", query)

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "couldn't create new student group")
	}

	return id, nil

}
func (r *PostgresRepository) PutStudentGroupById(ctx context.Context, id int64, data *model.UpdateStudentGroup) error {

	query := `UPDATE student_groups 
				SET 
					student_id = $1, 
					group_id = $2 
				WHERE id=$3`

	cmn, err := r.pool.Exec(ctx, query,
		data.StudentId,
		data.GroupId,
		id,
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't put student group by id: %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("couldn't put student group"),
			"couldn't put student group by id: %d", id)
	}

	return nil
}
func (r *PostgresRepository) DeleteStudentGroupById(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM student_groups WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete student group by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("student group doesn't not exist"),
			"couldn't delete student group by id %d", id)
	}
	return nil
}
