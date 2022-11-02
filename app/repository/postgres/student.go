package postgres

import m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"

func (r *PostgresRepository) GetStudents(ctx cntxt) ([]m.Student, error) {
	panic("empty")
}
func (r *PostgresRepository) GetStudentByStudentId(ctx cntxt, id int64) (*m.Student, error) {
	panic("empty")
}
func (r *PostgresRepository) AddStudent(ctx cntxt, data *m.Student) error {
	panic("empty")
}
func (r *PostgresRepository) EditStudentByStudentId(ctx cntxt, id int64, data *m.Student) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteStudentByStudentId(ctx cntxt, id int64) error {
	panic("empty")
}
