package postgres

import m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

func (r *PostgresRepository) GetStudentGroups(ctx cntxt) ([]m.StudentGroup, error) {
	panic("empty")
}
func (r *PostgresRepository) GetStudentGroupById(ctx cntxt, id int64) (*m.StudentGroup, error) {
	panic("empty")
}
func (r *PostgresRepository) AddStudentGroup(ctx cntxt, data *m.StudentGroup) error {
	panic("empty")
}
func (r *PostgresRepository) EditStudentGroupById(ctx cntxt, id int64, data *m.StudentGroup) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteStudentGroupById(ctx cntxt, id int64) error {
	panic("empty")
}
