package postgres

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strings"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetGroupById(ctx context.Context, id int64) (*model.Group, error) {
	var group model.Group
	err := r.pool.QueryRow(ctx, `
		SELECT 
			id, 
			course_id, 
			mentor_id,   
			title 
		FROM groups 
		WHERE id = $1 
	`, id).Scan(
		&group.ID,
		&group.CourseId,
		&group.MentorId,
		&group.Title,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get group by id: %d", id)
	}

	return &group, nil
}
func (r *PostgresRepository) GetGroups(ctx context.Context) ([]model.Group, error) {
	var groups []model.Group
	rows, err := r.pool.Query(ctx, `SELECT * FROM groups`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of groups")
	}

	for rows.Next() {
		var c model.Group
		err := rows.Scan(&c.ID,
			&c.CourseId,
			&c.MentorId,
			&c.Title,
		)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan group")
		}

		groups = append(groups, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "group rows error")
	}

	return groups, nil
}
func (r *PostgresRepository) CreateGroup(ctx context.Context, group *model.GroupCU) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	// add title
	if group.MentorId != 0 {
		args = append(args, group.MentorId)
		keys = append(keys, "mentor_id")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	// add start date
	if group.Title != "" {
		args = append(args, group.Title)
		keys = append(keys, "title")
		values = append(values, fmt.Sprintf("$%d", len(args)))
	}

	// prepare query to execute
	query := fmt.Sprintf(
		`INSERT INTO groups (%s) VALUES (%s) RETURNING id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)
	log.Print(query)

	r.lg.Debug("query", query)

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "couldn't create a group")
	}

	return id, nil
}
func (r *PostgresRepository) UpdateGroupById(ctx context.Context, id int64, group *model.GroupCU) error {

	var args []interface{}
	var keys []string

	// title
	if group.Title != "" {
		args = append(args, group.Title)
		keys = append(keys, fmt.Sprintf("title = $%d", len(args)))
	}

	if group.MentorId != 0 {
		args = append(args, group.MentorId)
		keys = append(keys, fmt.Sprintf("mentor_id = $%d", len(args)))
	}

	// id
	args = append(args, id)

	query := fmt.Sprintf(`UPDATE groups
				SET %s
				WHERE id = $%d`, strings.Join(keys, ","), len(args))

	cmn, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update group by id %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("group doesn't exist"),
			"couldn't update group by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) PutGroupById(ctx context.Context, id int64, group *model.GroupCU) error {
	query := `UPDATE groups 
				SET 
					mentor_id = $1, 
					title = $2 
				WHERE id=$3`

	cmn, err := r.pool.Exec(ctx, query,
		&group.MentorId,
		&group.Title,
		&id,
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't put group by id: %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("group doesn't exist"),
			"couldn't put group by id: %d", id)
	}

	return nil
}
func (r *PostgresRepository) DeleteGroupById(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM groups WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete group by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("group doesn't exist"),
			"couldn't delete group by id %d", id)
	}
	return nil
}
