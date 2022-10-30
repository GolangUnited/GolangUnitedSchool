package usecase

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type CertificateTemplateUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewCertificateTemplate(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *CertificateTemplateUsecase {
	return &CertificateTemplateUsecase{lg: lg, repo: repo}
}

func (u *CertificateTemplateUsecase) GetCertificateTemplates(ctx context.Context) ([]model.CertificateTemplate, error) {
	panic("not implemented")
}

func (u *CertificateTemplateUsecase) GetCertificateTemplateById(ctx context.Context, id int64) (*model.CertificateTemplate, error) {
	panic("not implemented")
}

func (u *CertificateTemplateUsecase) AddCertificateTemplate(ctx context.Context, data *model.CertificateTemplate) error {
	panic("not implemented")
}

func (u *CertificateTemplateUsecase) UpdateCertificateTemplate(ctx context.Context, id int64, data *model.CertificateTemplate) error {
	panic("not implemented")
}

func (u *CertificateTemplateUsecase) DeleteCertificateTemplate(ctx context.Context, id int64) error {
	panic("not implemented")
}
