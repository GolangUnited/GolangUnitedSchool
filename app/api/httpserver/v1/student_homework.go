package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentHomeworkHandlers struct {
	lg      logger.Logger
	usecase usecase.StudentHomeworkUsecaseInterface
}

func NewStudentHomeworkHandler(
	lg logger.Logger,
	usecase usecase.StudentHomeworkUsecaseInterface,
) *StudentHomeworkHandlers {
	return &StudentHomeworkHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

// @Summary get all items in the student homework list
// @ID get-all-student-homeworks
// @Tags students
// @Produce json
// @Success 200 {object} model.StudentHomeworkList
// @Failure 500 {object} model.ResponseMessage
// @Router /student-homeworks [get]
func (h *StudentHomeworkHandlers) GetStudentHomeworks(c *gin.Context) {
}

// @Summary get items in the student homework list by student ID
// @ID get-student-homeworks-by-student-id
// @Tags students
// @Produce json
// @Param student_id path string true "student id"
// @Success 200 {object} model.StudentHomeworkList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/homeworks [get]
func (h *StudentHomeworkHandlers) GetStudentHomeworksByStudentId(c *gin.Context) {
}

// @Summary get a student homework by ID
// @ID get-student-homework-by-id
// @Tags students
// @Produce json
// @Param id path string true "student homework id"
// @Success 200 {object} model.StudentHomework
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student-homeworks/{id} [get]
func (h *StudentHomeworkHandlers) GetStudentHomeworkById(c *gin.Context) {
}

// @Summary add new student homework to the student homework list
// @ID create-student-homework
// @Tags students
// @Produce json
// @Consume json
// @Param student_homework body model.StudentHomework true "student homework"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student-homeworks [post]
func (h *StudentHomeworkHandlers) AddStudentHomework(c *gin.Context) {
}

// @Summary update a student homework by ID
// @ID update-student-homework-by-id
// @Tags students
// @Produce json
// @Consume json
// @Param id path string true "student homework id"
// @Param student_homework body model.StudentHomework true "student homework"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student-homeworks/{id} [put]
func (h *StudentHomeworkHandlers) UpdateStudentHomework(c *gin.Context) {
}

// @Summary delete a student homework by ID
// @ID delete-student-homework-by-id
// @Tags students
// @Produce json
// @Param id path string true "student homework id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /student-homeworks/{id} [delete]
func (h *StudentHomeworkHandlers) DeleteStudentHomework(c *gin.Context) {
}
