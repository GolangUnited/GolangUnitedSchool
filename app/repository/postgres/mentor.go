package postgres

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
	"log"
	"strings"
)

func (r *PostgresRepository) GetMentors(ctx context.Context) ([]model.Mentor, error) {
	var mentors []model.Mentor
	rows, err := r.pool.Query(ctx, `SELECT * FROM mentors`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of mentors")
	}

	for rows.Next() {
		var c model.Mentor
		err := rows.Scan(&c.MentorId,
			&c.UserId,
		)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan mentors")
		}

		mentors = append(mentors, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "mentors rows error")
	}

	return mentors, nil
}
func (r *PostgresRepository) GetMentorById(ctx context.Context, id int64) (*model.Mentor, error) {
	var mentor model.Mentor
	err := r.pool.QueryRow(ctx, `
		SELECT 
			mentor_id, 
			user_id 
		FROM mentors
		WHERE mentor_id = $1 
	`, id).Scan(
		&mentor.MentorId,
		&mentor.UserId,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get mentor by id: %d", id)
	}

	return &mentor, nil
}
func (r *PostgresRepository) AddMentor(ctx context.Context, data *model.UpdateMentor) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	args = append(args, data.UserId)
	keys = append(keys, "user_id")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	// prepare query to execute
	query := fmt.Sprintf(
		`INSERT INTO mentors (%s) VALUES (%s) RETURNING mentor_id`,
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	)
	log.Print(query)

	r.lg.Debug("query", query)

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return id, errors.Wrap(err, "couldn't create a group contact")
	}

	return id, nil
}
func (r *PostgresRepository) UpdateMentorByMentorId(ctx context.Context, id int64, data *model.UpdateMentor) error {
	var args []interface{}
	var keys []string

	if data.UserId != nil {
		args = append(args, data.UserId)
		keys = append(keys, fmt.Sprintf("user_id = $%d", len(args)))
	}

	args = append(args, id)

	query := fmt.Sprintf(`UPDATE mentors
				SET %s
				WHERE mentor_id = $%d`, strings.Join(keys, ","), len(args))

	cmn, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update mentor by id %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("mentor doesn't exist"),
			"couldn't update mentor by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) DeleteMentorByMentorId(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM mentors WHERE mentor_id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete mentor by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("mentor doesn't exist"),
			"couldn't delete mentor by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) PutMentorById(ctx context.Context, id int64, data *model.UpdateMentor) error {
	query := `UPDATE mentors 
				SET 
					user_id = $1 
					
				WHERE mentor_id=$2`

	cmn, err := r.pool.Exec(ctx, query,
		&data.UserId,
		&id,
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't put mentor by id: %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("mentor doesn't exist"),
			"couldn't put mentor by id: %d", id)
	}

	return nil
}

//GetMentors(ctx context.Context) ([]model.Mentor, error)
//GetMentorById(ctx context.Context, id int64) (*model.Mentor, error)
//AddMentor(ctx context.Context, data *model.UpdateMentor) (int64, error)
//UpdateMentorByMentorId(ctx context.Context, id int64, data *model.UpdateMentor) error
//DeleteMentorByMentorId(ctx context.Context, id int64) error
//PutMentorById(ctx context.Context, id int64, data *model.UpdateMentor) error
