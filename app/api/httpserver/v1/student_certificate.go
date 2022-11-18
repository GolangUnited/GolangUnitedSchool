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
// @Summary get all student certificates
// @ID get-all-certificates
// @Description возвращает все сертификаты всех студентов
// @Tags students, certificates
// @Produce json
// @Success 200 {object} model.StudentCertificatesListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates [get]
func (h *StudentCertificateHandlers) GetStudentCertificates(c *gin.Context) {
}

// GetStudentCertificatesByStudentId
// @Summary get all certificates of student
// @ID get-student-certificates
// @Description возвращает все сертификаты студента с указанным id
// @Param id path string true "student_id"
// @Tags students, certificates
// @Produce json
// @Success 200 {object} model.StudentCertificatesListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/certificates [get]
func (h *StudentCertificateHandlers) GetStudentCertificatesByStudentId(c *gin.Context) {
}

// GetStudentCertificatesByCourseId
// @Summary get certificates by course
// @ID get-course-certificates
// @Description возвращает сертификаты курса
// @Param id path string true "course_id"
// @Tags students, certificates, courses
// @Produce json
// @Success 200 {object} model.StudentCertificatesListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates/courses/{course_id} [get]
func (h *StudentCertificateHandlers) GetStudentCertificatesByCourseId(c *gin.Context) {
}

// GetStudentCertificateById
// @Summary get student certificate
// @ID get-student-certificate
// @Description возвращает сертификат студента с указанным id
// @Param id path string true "certificate_id"
// @Param id path string true "student_id"
// @Tags students, certificates
// @Produce json
// @Success 200 {object} model.StudentCertificate
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/certificates/{certificate_od} [get]
func (h *StudentCertificateHandlers) GetStudentCertificateById(c *gin.Context) {
}

// AddStudentCertificate
// @Summary add student certificate
// @Description добавляет новый сертификат студента
// @ID add-student-certificate
// @Param certificate body model.StudentCertificate true "certificate"
// @Tags students, certificates
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates [post]
func (h *StudentCertificateHandlers) AddStudentCertificate(c *gin.Context) {
}

// UpdateStudentCertificate
// @Summary update student certificate
// @Description изменяет новый сертификат студента
// @ID put-student-certificate
// @Param certificate body model.StudentCertificate true "certificate"
// @Tags students, certificates
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates [put]
func (h *StudentCertificateHandlers) UpdateStudentCertificate(c *gin.Context) {
}

// DeleteStudentCertificate
// @Summary delete student certificate
// @Description удаляет сертификат студента
// @ID delete-student-certificate
// @Param id path string true "certificate_id"
// @Tags students, certificates
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/certificates/{certificate_id} [delete]
func (h *StudentCertificateHandlers) DeleteStudentCertificate(c *gin.Context) {
}
