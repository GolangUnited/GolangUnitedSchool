package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetStudentCertificates(ctx context.Context) ([]model.StudentCertificate, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetStudentCertificatesByStudentId(ctx context.Context, studentId int64) ([]model.StudentCertificate, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetStudentCertificatesByCourseId(ctx context.Context, courseId int64) ([]model.StudentCertificate, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetStudentCertificateById(ctx context.Context, id int64) (*model.StudentCertificate, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddStudentCertificate(ctx context.Context, data *model.StudentCertificate) error {
	panic("not implemented")
}

func (r *PostgresRepository) UpdateStudentCertificate(ctx context.Context, id int64, data *model.StudentCertificate) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteStudentCertificate(ctx context.Context, id int64) error {
	panic("not implemented")
}
