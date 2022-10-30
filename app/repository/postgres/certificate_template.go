package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (u *repository) GetCertificateTemplates(ctx context.Context) ([]model.CertificateTemplate, error) {
	panic("not implemented")
}

func (u *repository) GetCertificateTemplateById(ctx context.Context, id int64) (*model.CertificateTemplate, error) {
	panic("not implemented")
}

func (u *repository) AddCertificateTemplate(ctx context.Context, data *model.CertificateTemplate) error {
	panic("not implemented")
}

func (u *repository) UpdateCertificateTemplate(ctx context.Context, id int64, data *model.CertificateTemplate) error {
	panic("not implemented")
}

func (u *repository) DeleteCertificateTemplate(ctx context.Context, id int64) error {
	panic("not implemented")
}
