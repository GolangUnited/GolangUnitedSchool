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

// GetStudentHomeworks
// @Summary get all homeworks
// @ID get-all-students-homeworks
// @Description возвращает все домашние задания всех студентов в виде списка
// @Tags studentHomeworks
// @Produce json
// @Success 200 {object} model.StudentHomeworkList
// @Failure 500 {object} model.ResponseMessage
// @Router /students/homeworks [get]
func (h *StudentHomeworkHandlers) GetStudentHomeworks(c *gin.Context) {
}

// GetStudentHomeworksByStudentId
// @Summary get student homeworks by student ID
// @Description возвращает все домашние задания студента с указанным id студента
// @ID get-student-homeworks-by-student-id
// @Tags studentHomeworks
// @Produce json
// @Param student_id path string true "student id"
// @Success 200 {object} model.StudentHomeworkList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/homeworks [get]
func (h *StudentHomeworkHandlers) GetStudentHomeworksByStudentId(c *gin.Context) {
}

// GetStudentHomeworkById
// @Summary get student homework by ID
// @Description возвращает домашнее задание студента с указанным id
// @ID get-student-homework-by-id
// @Param id path string true "homework id"
// @Tags studentHomeworks
// @Produce json
// @Success 200 {object} model.StudentHomework
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/homeworks/{id} [get]
func (h *StudentHomeworkHandlers) GetStudentHomeworkById(c *gin.Context) {
}

// AddStudentHomework
// @Summary create student homework
// @Description добавляет новое домашнее задание студента
// @ID create-student-homework
// @Tags studentHomeworks
// @Produce json
// @Consume json
// @Param student_homework body model.StudentHomework true "student homework"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/homeworks [post]
func (h *StudentHomeworkHandlers) AddStudentHomework(c *gin.Context) {
}

// UpdateStudentHomework
// @Summary update student homework by ID
// @Description изменяет домашнее задание студента с указанным id
// @ID update-student-homework-by-id
// @Tags studentHomeworks
// @Produce json
// @Consume json
// @Param id path string true "student homework id"
// @Param student_homework body model.StudentHomework true "student homework"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/homeworks/{id} [put]
func (h *StudentHomeworkHandlers) UpdateStudentHomework(c *gin.Context) {
}

// DeleteStudentHomework
// @Summary delete student homework by ID
// @Description удаляет домашнее задание студента с указанным id
// @ID delete-student-homework-by-id
// @Tags studentHomeworks
// @Produce json
// @Param id path string true "student homework id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/homeworks/{id} [delete]
func (h *StudentHomeworkHandlers) DeleteStudentHomework(c *gin.Context) {
}
