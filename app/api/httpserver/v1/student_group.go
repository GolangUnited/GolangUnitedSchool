package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentGroupHandlers struct {
	lg      logger.Logger
	useCase usecase.StudentGroupUseCaseInterface
}

func NewStudentGroupHandler(
	lg logger.Logger,
	useCase usecase.StudentGroupUseCaseInterface,
) *StudentGroupHandlers {
	return &StudentGroupHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// GetStudentGroups
// @Summary get all student groups
// @Description возвращает список всех студенческих групп
// @ID get-student-groups
// @Tags students, studentGroups
// @Produce json
// @Success 200 {object} model.StudentGroupListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /students/groups [get]
func (h *StudentGroupHandlers) GetStudentGroups(c *gin.Context) {}

// GetStudentGroupById
// @Summary get student group by id
// @Description возвращает группу с указанным id
// @ID get_student_group_by_id
// @Tags students, studentGroups
// @Produce json
// @Param id path string true "group_id"
// @Success 200 {object} model.StudentGroup
// @Failure 404 {object} model.ResponseMessage
// @Router /students/groups/{group_id} [get]
func (h *StudentGroupHandlers) GetStudentGroupById(c *gin.Context) {}

// AddStudentGroup
// @Summary add new student group
// @Description добавляет новую студенческую группу
// @ID add-student-group
// @Tags students, studentGroups
// @Produce json
// @Consume json
// @Param course body model.StudentGroup true "student_group"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/groups [post]
func (h *StudentGroupHandlers) AddStudentGroup(c *gin.Context) {}

// UpdateStudentGroupById
// @Summary update student group by group id
// @Description изменяет данные студенческой группы
// @ID update-student-group-by-id
// @Tags students, studentGroups
// @Param id path string true "group_id"
// @Param student group body model.StudentGroup true "student_group"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /students/groups/{group_id} [put]
func (h *StudentGroupHandlers) UpdateStudentGroupById(c *gin.Context) {}

// DeleteStudentGroupById
// @Summary delete student group by id
// @Description удаляет студенческую группу с указанным id
// @ID delete-student-group-by-id
// @Tags students, studentGroups
// @Param id path string true "group_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/groups/{group_id} [delete]
func (h *StudentGroupHandlers) DeleteStudentGroupById(c *gin.Context) {}
