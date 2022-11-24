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
// @Description добавить пользователя в группу студентов
// @ID add-person-to-students
// @Tags students
// @Produce json
// @Consume json
// @Param id path string true "person_id"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{person_id} [post]
func (h *StudentHandlers) AddStudent(c *gin.Context) {}

// DeleteStudentByStudentId
// @Summary delete person from students
// @Description удалить пользователя с указанным id is группы "студенты"
// @ID delete-person-from-students
// @Tags students
// @Param id path string true "student_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id} [delete]
func (h *StudentHandlers) DeleteStudentByStudentId(c *gin.Context) {}

// PutStudentByStudentId
// @Summary update student by student id
// @Description пока не слишком понятно, нужен ли этот хендлер
// @ID update-student-by-student-id
// @Tags students
// @Param id path string true "student_id"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /students/{student_id} [put]
func (h *StudentHandlers) PutStudentByStudentId(c *gin.Context) {}

// GetStudentByStudentId
// @Summary get student by student id
// @Description получить данные о студенте по его id, возвращает структуру person
// @ID get-student-by-student-id
// @Tags students
// @Produce json
// @Param id path string true "student_id"
// @Success 200 {object} model.Person
// @Failure 404 {object} model.ResponseMessage
// @Router /students/{student_id} [get]
func (h *StudentHandlers) GetStudentByStudentId(c *gin.Context) {}

// GetStudents
// @Summary get all students
// @Description возвращает список всех пользователей группы "student"
// @ID get-all-students
// @Tags students
// @Produce json
// @Success 200 {object} model.StudentsListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /students [get]
func (h *StudentHandlers) GetStudents(c *gin.Context) {}
