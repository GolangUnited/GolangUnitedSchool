package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentCertificateHandlers struct {
	lg      logger.Logger
	usecase usecase.StudentCertificateUsecaseInterface
}

func NewStudentCertificateHandler(
	lg logger.Logger,
	usecase usecase.StudentCertificateUsecaseInterface,
) *StudentCertificateHandlers {
	return &StudentCertificateHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *StudentCertificateHandlers) GetStudentCertificates(c *gin.Context) {
}

func (h *StudentCertificateHandlers) GetStudentCertificatesByStudentId(c *gin.Context) {
}

func (h *StudentCertificateHandlers) GetStudentCertificatesByCourseId(c *gin.Context) {
}

func (h *StudentCertificateHandlers) GetStudentCertificateById(c *gin.Context) {
}

func (h *StudentCertificateHandlers) AddStudentCertificate(c *gin.Context) {
}

func (h *StudentCertificateHandlers) UpdateStudentCertificate(c *gin.Context) {
}

func (h *StudentCertificateHandlers) DeleteStudentCertificate(c *gin.Context) {
}
