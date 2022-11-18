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

// GetCertificateTemplates
// @Summary get all certificate templates
// @Description возвращает все шаблоны сертификатов
// @ID get-all-certificate-templates
// @Tags certificates, certificateTemplates
// @Produce json
// @Success 200 {object} model.CertificateTemplatesListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /certificates/templates [get]
func (h *CertificateTemplateHandlers) GetCertificateTemplates(c *gin.Context) {
}

// GetCertificateTemplateById
// @Summary get certificate template by id
// @ID get-crt-template-by-id
// @Description возвращает шаблон сертификата с указанным id
// @Tags certificates, certificateTemplates
// @Produce json
// @Param id path string true "crt_template_id"
// @Success 200 {object} model.CertificateTemplate
// @Failure 404 {object} model.ResponseMessage
// @Router /certificates/templates/{crt_template_id} [get]
func (h *CertificateTemplateHandlers) GetCertificateTemplateById(c *gin.Context) {
}

// AddCertificateTemplate
// @Summary add new certificate template
// @ID add-crt-template
// @Description добавляет новый шаблон сертификата
// @Tags certificates, certificateTemplates
// @Produce json
// @Param crt template body model.CertificateTemplate true "crt template"
// @Success 200 {object} model.CertificateTemplate
// @Failure 404 {object} model.ResponseMessage
// @Router /certificates/templates [post]
func (h *CertificateTemplateHandlers) AddCertificateTemplate(c *gin.Context) {
}

// UpdateCertificateTemplate
// @Summary update certificate template
// @ID put-crt-template
// @Description изменяет шаблон сертификата
// @Tags certificates, certificateTemplates
// @Produce json
// @Param crt template body model.CertificateTemplate true "crt template"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /certificates/templates [put]
func (h *CertificateTemplateHandlers) UpdateCertificateTemplate(c *gin.Context) {
}

// DeleteCertificateTemplate
// @Summary delete certificate template
// @ID delete-crt-template
// @Description удаляет шаблон сертификата
// @Tags certificates, certificateTemplates
// @Produce json
// @Param id path string true "crt_template_id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /certificates/templates/{crt_template_id} [delete]
func (h *CertificateTemplateHandlers) DeleteCertificateTemplate(c *gin.Context) {
}
