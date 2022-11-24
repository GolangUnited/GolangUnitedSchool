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

// GetStudentCertificates
// @Summary get all items in the student certificate list
// @Description возвращает все сертификаты
// @ID get-all-student-certificates
// @Tags studentCertificates
// @Produce json
// @Success 200 {object} model.StudentCertificateList
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates [get]
func (h *StudentCertificateHandlers) GetStudentCertificates(c *gin.Context) {
}

// GetStudentCertificatesByStudentId
// @Summary get items in the student certificate list by student ID
// @Description возвращает сертификаты с указанным id студента
// @ID get-student-certificates-by-student-id
// @Tags studentCertificates
// @Produce json
// @Param student_id path string true "student id"
// @Success 200 {object} model.StudentCertificateList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/certificates [get]
func (h *StudentCertificateHandlers) GetStudentCertificatesByStudentId(c *gin.Context) {
}

// GetStudentCertificatesByCourseId
// @Summary get certificates by course ID
// @Description возвращает сертификаты с указанным id курса
// @ID get-student-certificates-by-course-id
// @Tags studentCertificates
// @Produce json
// @Param course_id path string true "course id"
// @Success 200 {object} model.StudentCertificateList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/{course_id}/certificates [get]
func (h *StudentCertificateHandlers) GetStudentCertificatesByCourseId(c *gin.Context) {
}

// GetStudentCertificateById
// @Summary get student certificate by ID
// @ID get-student-certificate-by-id
// @Description возвращает сертификат с указанным id
// @Tags studentCertificates
// @Produce json
// @Param id path string true "certificate id"
// @Success 200 {object} model.StudentCertificate
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates/{id} [get]
func (h *StudentCertificateHandlers) GetStudentCertificateById(c *gin.Context) {
}

// AddStudentCertificate
// @Summary add student certificate
// @Description добавляет новый сертификат
// @ID add-student-certificate
// @Tags studentCertificates
// @Produce json
// @Consume json
// @Param certificate body model.StudentCertificate true "certificate"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates [post]
func (h *StudentCertificateHandlers) AddStudentCertificate(c *gin.Context) {
}

// UpdateStudentCertificate
// @Summary update student certificate by ID
// @Description изменяет сертификат с указанным id
// @ID update-student-certificate-by-id
// @Tags studentCertificates
// @Produce json
// @Consume json
// @Param id path string true "certificate id"
// @Param certificate body model.StudentCertificate true "certificate"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates/{id} [put]
func (h *StudentCertificateHandlers) UpdateStudentCertificate(c *gin.Context) {
}

// DeleteStudentCertificate
// @Summary delete student certificate by ID
// @Description удаляет сертификат с указанным id
// @ID delete-student-certificate-by-id
// @Tags studentCertificates
// @Produce json
// @Param id path string true "certificate id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates/{id} [delete]
func (h *StudentCertificateHandlers) DeleteStudentCertificate(c *gin.Context) {
}
