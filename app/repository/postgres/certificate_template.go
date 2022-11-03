package postgres

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

func (r *PostgresRepository) GetCertificateTemplates(ctx context.Context) ([]model.CertificateTemplate, error) {
	panic("not implemented")
}

func (r *PostgresRepository) GetCertificateTemplateById(ctx context.Context, id int64) (*model.CertificateTemplate, error) {
	panic("not implemented")
}

func (r *PostgresRepository) AddCertificateTemplate(ctx context.Context, data *model.CertificateTemplate) error {
	panic("not implemented")
}

func (r *PostgresRepository) UpdateCertificateTemplate(ctx context.Context, id int64, data *model.CertificateTemplate) error {
	panic("not implemented")
}

func (r *PostgresRepository) DeleteCertificateTemplate(ctx context.Context, id int64) error {
	panic("not implemented")
}
