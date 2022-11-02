package student

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentCertificateUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewStudentCertificate(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *StudentCertificateUsecase {
	return &StudentCertificateUsecase{lg: lg, repo: repo}
}

func (u *StudentCertificateUsecase) GetStudentCertificates(
	ctx context.Context) ([]model.StudentCertificate, error) {
	panic("not implemented")
}

func (u *StudentCertificateUsecase) GetStudentCertificatesByStudentId(
	ctx context.Context,
	studentId int64) ([]model.StudentCertificate, error) {
	panic("not implemented")
}

func (u *StudentCertificateUsecase) GetStudentCertificatesByCourseId(
	ctx context.Context,
	courseId int64) ([]model.StudentCertificate, error) {
	panic("not implemented")
}

func (u *StudentCertificateUsecase) GetStudentCertificateById(
	ctx context.Context,
	id int64) (*model.StudentCertificate, error) {
	panic("not implemented")
}

func (u *StudentCertificateUsecase) AddStudentCertificate(
	ctx context.Context,
	data *model.StudentCertificate) error {
	panic("not implemented")
}

func (u *StudentCertificateUsecase) UpdateStudentCertificate(
	ctx context.Context,
	id int64,
	data *model.StudentCertificate) error {
	panic("not implemented")
}

func (u *StudentCertificateUsecase) DeleteStudentCertificate(
	ctx context.Context,
	id int64) error {
	panic("not implemented")
}
