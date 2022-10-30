package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (u *repository) GetStudentCertificates(ctx context.Context) ([]model.StudentCertificate, error) {
	panic("not implemented")
}

func (u *repository) GetStudentCertificatesByStudentId(ctx context.Context, studentId int64) ([]model.StudentCertificate, error) {
	panic("not implemented")
}

func (u *repository) GetStudentCertificatesByCourseId(ctx context.Context, courseId int64) ([]model.StudentCertificate, error) {
	panic("not implemented")
}

func (u *repository) GetStudentCertificateById(ctx context.Context, id int64) (*model.StudentCertificate, error) {
	panic("not implemented")
}

func (u *repository) AddStudentCertificate(ctx context.Context, data *model.StudentCertificate) error {
	panic("not implemented")
}

func (u *repository) UpdateStudentCertificate(ctx context.Context, id int64, data *model.StudentCertificate) error {
	panic("not implemented")
}

func (u *repository) DeleteStudentCertificate(ctx context.Context, id int64) error {
	panic("not implemented")
}
