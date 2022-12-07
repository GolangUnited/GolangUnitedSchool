package postgres

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/pkg/errors"
	"log"
	"strings"
)

func (r *PostgresRepository) GetGroupContacts(ctx context.Context) ([]model.GroupContact, error) {
	var groups []model.GroupContact
	rows, err := r.pool.Query(ctx, `SELECT * FROM group_contacts`)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of group contacts")
	}

	for rows.Next() {
		var c model.GroupContact
		err := rows.Scan(&c.GroupContactId,
			&c.ContactValue,
			&c.IsPrimary,
			&c.ContactTypeId,
			&c.GroupId,
		)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan group contacts")
		}

		groups = append(groups, c)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "group rows error")
	}

	return groups, nil
}
func (r *PostgresRepository) GetGroupContactById(ctx context.Context, id int64) (*model.GroupContact, error) {
	var groupContact model.GroupContact
	err := r.pool.QueryRow(ctx, `
		SELECT 
			id, 
			contact_value, 
			is_primary,   
			contact_type,
			group_id
		FROM group_contacts
		WHERE id = $1 
	`, id).Scan(
		&groupContact.GroupContactId,
		&groupContact.ContactValue,
		&groupContact.IsPrimary,
		&groupContact.ContactTypeId,
		&groupContact.GroupId,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get group contact by id: %d", id)
	}

	return &groupContact, nil
}
func (r *PostgresRepository) AddGroupContact(ctx context.Context, data *model.GroupContactCU) (int64, error) {
	var args []interface{}
	var keys []string
	var values []string

	args = append(args, data.ContactValue)
	keys = append(keys, "contact_value")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	args = append(args, data.IsPrimary)
	keys = append(keys, "is_primary")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	args = append(args, data.ContactTypeId)
	keys = append(keys, "contact_type")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	args = append(args, data.GroupId)
	keys = append(keys, "group_id")
	values = append(values, fmt.Sprintf("$%d", len(args)))

	// prepare query to execute
	query := fmt.Sprintf(
		`INSERT INTO group_contacts (%s) VALUES (%s) RETURNING id`,
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
func (r *PostgresRepository) PutGroupContactById(ctx context.Context, id int64, data *model.GroupContactCU) error {
	query := `UPDATE group_contacts 
				SET 
					contact_value = $1, 
					is_primary = $2,
					contact_type = $3,
					group_id = $4
				WHERE id=$5`

	cmn, err := r.pool.Exec(ctx, query,
		&data.ContactValue,
		&data.IsPrimary,
		&data.ContactTypeId,
		&data.GroupId,
		&id,
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't put group contact by id: %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("group contact doesn't exist"),
			"couldn't put group contact by id: %d", id)
	}

	return nil
}
func (r *PostgresRepository) UpdateGroupContactById(ctx context.Context, id int64, data *model.GroupContactUpdate) error {

	var args []interface{}
	var keys []string

	if data.ContactValue != nil {
		args = append(args, data.ContactValue)
		keys = append(keys, fmt.Sprintf("contact_value = $%d", len(args)))
	}
	if data.IsPrimary != nil {
		args = append(args, data.IsPrimary)
		keys = append(keys, fmt.Sprintf("is_primary = $%d", len(args)))
	}
	if data.ContactTypeId != nil {
		args = append(args, data.ContactTypeId)
		keys = append(keys, fmt.Sprintf("contact_type = $%d", len(args)))
	}
	if data.GroupId != nil {
		args = append(args, data.GroupId)
		keys = append(keys, fmt.Sprintf("group_id = $%d", len(args)))
	}

	args = append(args, id)

	query := fmt.Sprintf(`UPDATE group_contacts
				SET %s
				WHERE id = $%d`, strings.Join(keys, ","), len(args))

	cmn, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrapf(err,
			"couldn't update group contact by id %d", id)
	}

	if cmn.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("group contact doesn't exist"),
			"couldn't update group contact by id %d", id)
	}
	return nil
}
func (r *PostgresRepository) DeleteGroupContactById(ctx context.Context, id int64) error {
	pgt, err := r.pool.Exec(ctx, `DELETE FROM group_contacts WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete group contact by id %d", id)
	}
	if pgt.RowsAffected() != 1 {
		return errors.Wrapf(errors.New("group contact doesn't exist"),
			"couldn't delete group contact by id %d", id)
	}
	return nil
}

///еще одна функция на все контакты одной группы
/// еще одна функция на все коннтакты в принципе
