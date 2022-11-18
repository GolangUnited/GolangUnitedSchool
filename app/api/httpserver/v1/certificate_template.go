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

// @Summary get all items in the certificate template list
// @ID get-all-certificate-templates
// @Tags courses
// @Produce json
// @Success 200 {object} model.CertificateTemplateList
// @Failure 500 {object} model.ResponseMessage
// @Router /certificate-templates [get]
func (h *CertificateTemplateHandlers) GetCertificateTemplates(c *gin.Context) {
}

// @Summary get a certificate template by ID
// @ID get-certificate-template-by-id
// @Tags courses
// @Produce json
// @Param id path string true "certificate template id"
// @Success 200 {object} model.CertificateTemplate
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /certificate-templates/{id} [get]
func (h *CertificateTemplateHandlers) GetCertificateTemplateById(c *gin.Context) {
}

// @Summary add new certificate template to the certificate template list
// @ID create-certificate-template
// @Tags courses
// @Produce json
// @Consume json
// @Param certificate_template body model.CertificateTemplate true "certificate template"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /certificate-templates [post]
func (h *CertificateTemplateHandlers) AddCertificateTemplate(c *gin.Context) {
}

// @Summary update a certificate template by ID
// @ID update-certificate-template-by-id
// @Tags courses
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

// @Summary delete a certificate template by ID
// @ID delete-certificate-template-by-id
// @Tags courses
// @Produce json
// @Param id path string true "certificate template id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /certificate-templates/{id} [delete]
func (h *CertificateTemplateHandlers) DeleteCertificateTemplate(c *gin.Context) {
}
