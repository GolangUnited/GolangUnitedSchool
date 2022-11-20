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
// @ID get-students-homeworks
// @Description возвращает все домашние работы всех студентов в виде списка
// @Tags studentHomeworks
// @Produce json
// @Success 200 {object} model.StudentHomeworksListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /students/homeworks [get]
func (h *StudentHomeworkHandlers) GetStudentHomeworks(c *gin.Context) {
}

// GetStudentHomeworksByStudentId
// @Summary get contact type
// @Description возвращает все домашние работы студента с указанным student_id
// @ID get-all-student-homeworks
// @Param id path string true "student_id"
// @Tags studentHomeworks
// @Produce json
// @Success 200 {object} model.StudentHomeworksListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/homeworks [get]
func (h *StudentHomeworkHandlers) GetStudentHomeworksByStudentId(c *gin.Context) {
}

// GetStudentHomeworkById
// @Summary get student's homework
// @Description возвращает определенную домашку указанного студента
// @ID get-student-homework
// @Param id path string true "student_id"
// @Param id path string true "homework_id"
// @Tags studentHomeworks
// @Produce json
// @Success 200 {object} model.StudentHomework
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/homeworks/{homework_id} [get]
func (h *StudentHomeworkHandlers) GetStudentHomeworkById(c *gin.Context) {
}

// AddStudentHomework
// @Summary get contact type
// @Description добавляет новый тип контакта
// @ID add-contact-type
// @Param contact_type body model.NewContactTypeDto true "contact_type"
// @Tags studentHomeworks
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/homeworks [post]
func (h *StudentHomeworkHandlers) AddStudentHomework(c *gin.Context) {
}

// UpdateStudentHomework
// @Summary update student homework
// @Description изменяет домашнюю работу студента
// @ID update-student-homework
// @Param student_homework body model.StudentHomework true "student_homework"
// @Tags studentHomeworks
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/homeworks [put]
func (h *StudentHomeworkHandlers) UpdateStudentHomework(c *gin.Context) {
}

// DeleteStudentHomework
// @Summary delete student homework
// @Description удаляет домашнюю работу студента
// @ID delete-student-homework
// @Param id path string true "student_homework_id"
// @Tags studentHomeworks
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/homeworks/{:contact_type_id} [delete]
func (h *StudentHomeworkHandlers) DeleteStudentHomework(c *gin.Context) {
}
