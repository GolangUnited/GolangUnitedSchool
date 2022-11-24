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
// @Summary get all items in the certificate template list
// Description возвращает все доступные шаблоны сертификатов
// @ID get-all-certificate-templates
// @Tags certificateTemplates
// @Produce json
// @Success 200 {object} model.CertificateTemplateList
// @Failure 500 {object} model.ResponseMessage
// @Router /certificate-templates [get]
func (h *CertificateTemplateHandlers) GetCertificateTemplates(c *gin.Context) {
}

// GetCertificateTemplateById
// @Summary get a certificate template by ID
// @Description возвращает шаблон сертификата с указанным id
// @ID get-certificate-template-by-id
// @Tags certificateTemplates
// @Produce json
// @Param id path string true "certificate template id"
// @Success 200 {object} model.CertificateTemplate
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /certificate-templates/{id} [get]
func (h *CertificateTemplateHandlers) GetCertificateTemplateById(c *gin.Context) {
}

// AddCertificateTemplate
// @Summary add new certificate template to the certificate template list
// @Description добавляет новый шаблон сертификата
// @ID create-certificate-template
// @Tags certificateTemplates
// @Produce json
// @Consume json
// @Param certificate_template body model.CertificateTemplate true "certificate template"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /certificate-templates [post]
func (h *CertificateTemplateHandlers) AddCertificateTemplate(c *gin.Context) {
}

// UpdateCertificateTemplate
// @Summary update a certificate template by ID
// @Description изменяет существующий щаблон сертификата с указанным id
// @ID update-certificate-template-by-id
// @Tags certificateTemplates
// @Produce json
// @Consume json
// @Param id path string true "certificate template id"
// @Param certificate_template body model.CertificateTemplate true "certificate template"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /certificate-templates/{id} [put]
func (h *CertificateTemplateHandlers) UpdateCertificateTemplate(c *gin.Context) {
}

// DeleteCertificateTemplate
// @Summary delete a certificate template by ID
// @Description удаляет шаблон сертификата с указанным id
// @ID delete-certificate-template-by-id
// @Tags certificateTemplates
// @Produce json
// @Param id path string true "certificate template id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /certificate-templates/{id} [delete]
func (h *CertificateTemplateHandlers) DeleteCertificateTemplate(c *gin.Context) {
}
