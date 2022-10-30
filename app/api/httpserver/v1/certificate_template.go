package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type CertificateTemplateHandlers struct {
	lg      logger.Logger
	usecase usecase.CertificateTemplateUsecaseInterface
}

func NewCertificateTemplateHandler(
	lg logger.Logger,
	usecase usecase.CertificateTemplateUsecaseInterface,
) *CertificateTemplateHandlers {
	return &CertificateTemplateHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *CertificateTemplateHandlers) GetCertificateTemplates(c *gin.Context) {
}

func (h *CertificateTemplateHandlers) GetCertificateTemplateById(c *gin.Context) {
}

func (h *CertificateTemplateHandlers) AddCertificateTemplate(c *gin.Context) {
}

func (h *CertificateTemplateHandlers) UpdateCertificateTemplate(c *gin.Context) {
}

func (h *CertificateTemplateHandlers) DeleteCertificateTemplate(c *gin.Context) {
}
