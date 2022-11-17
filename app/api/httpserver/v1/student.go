package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentHandlers struct {
	lg      logger.Logger
	useCase usecase.StudentUseCaseInterface
}

func NewStudentHandler(
	lg logger.Logger,
	useCase usecase.StudentUseCaseInterface,
) *StudentHandlers {
	return &StudentHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// AddStudent
// @Summary add person to students
// @ID cadd-person-to-students
// @Tags students
// @Produce json
// @Consume json
// @Param id path string true "person_id"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student/{person_id} [post]
func (h *StudentHandlers) AddStudent(c *gin.Context) {}

// DeleteStudentByStudentId
// @Summary delete person from students
// @ID delete-person-from-students
// @Tags students
// @Param id path string true "student_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student/{student_id} [delete]
func (h *StudentHandlers) DeleteStudentByStudentId(c *gin.Context) {}

// UpdateStudentByStudentId
// @Summary update student by student id
// @ID update-student-by-student-id
// @Tags students
// @Param id path string true "student_id"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /student/{student_id} [put]
func (h *StudentHandlers) UpdateStudentByStudentId(c *gin.Context) {}

// GetStudentByStudentId
// @Summary get student by student id
// @ID get-student-by-student-id
// @Tags students
// @Produce json
// @Param id path string true "student_id"
// @Success 200 {object} model.Course
// @Failure 404 {object} model.ResponseMessage
// @Router /student/{student_id} [get]
func (h *StudentHandlers) GetStudentByStudentId(c *gin.Context) {}

// GetStudents
// @Summary get all students
// @ID get-all-students
// @Tags students
// @Produce json
// @Success 200 {object} model.CourseList
// @Failure 500 {object} model.ResponseMessage
// @Router /students [get]
func (h *StudentHandlers) GetStudents(c *gin.Context) {}
