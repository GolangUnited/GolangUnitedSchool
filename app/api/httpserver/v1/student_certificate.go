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

// @Summary get all items in the student certificate list
// @ID get-all-student-certificates
// @Tags students
// @Produce json
// @Success 200 {object} model.StudentCertificateList
// @Failure 500 {object} model.ResponseMessage
// @Router /student-certificates [get]
func (h *StudentCertificateHandlers) GetStudentCertificates(c *gin.Context) {
}

// @Summary get items in the student certificate list by student ID
// @ID get-student-certificates-by-student-id
// @Tags students
// @Produce json
// @Param student_id path string true "student id"
// @Success 200 {object} model.StudentCertificateList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/certificates [get]
func (h *StudentCertificateHandlers) GetStudentCertificatesByStudentId(c *gin.Context) {
}

// @Summary get items in the student certificate list by course ID
// @ID get-student-certificates-by-course-id
// @Tags courses, students
// @Produce json
// @Param course_id path string true "course id"
// @Success 200 {object} model.StudentCertificateList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/{course_id}/certificates [get]
func (h *StudentCertificateHandlers) GetStudentCertificatesByCourseId(c *gin.Context) {
}

// @Summary get a student certificate by ID
// @ID get-student-certificate-by-id
// @Tags students
// @Produce json
// @Param id path string true "student certificate id"
// @Success 200 {object} model.StudentCertificate
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student-certificates/{id} [get]
func (h *StudentCertificateHandlers) GetStudentCertificateById(c *gin.Context) {
}

// @Summary add new student certificate to the student certificate list
// @ID create-student-certificate
// @Tags courses
// @Produce json
// @Consume json
// @Param student_certificate body model.StudentCertificate true "student certificate"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student-certificates [post]
func (h *StudentCertificateHandlers) AddStudentCertificate(c *gin.Context) {
}

// @Summary update a student certificate by ID
// @ID update-student-certificate-by-id
// @Tags courses
// @Produce json
// @Consume json
// @Param id path string true "student certificate id"
// @Param student_certificate body model.StudentCertificate true "student certificate"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student-certificates/{id} [put]
func (h *StudentCertificateHandlers) UpdateStudentCertificate(c *gin.Context) {
}

// @Summary delete a student certificate by ID
// @ID delete-student-certificate-by-id
// @Tags courses
// @Produce json
// @Param id path string true "student certificate id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student-certificates/{id} [delete]
func (h *StudentCertificateHandlers) DeleteStudentCertificate(c *gin.Context) {
}
