package postgres

import m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

func (r *PostgresRepository) GetMentors(ctx cntxt) ([]m.Mentor, error) {
	panic("empty")
}
func (r *PostgresRepository) GetMentorById(ctx cntxt, id int64) (*m.Mentor, error) {
	panic("empty")
}
func (r *PostgresRepository) AddMentor(ctx cntxt, data *m.Mentor) error {
	panic("empty")
}
func (r *PostgresRepository) EditMentorByMentorId(ctx cntxt, id int64, data *m.Mentor) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteMentorByMentorId(ctx cntxt, id int64) error {
	panic("empty")
}
