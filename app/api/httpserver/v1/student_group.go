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
// @ID get-student-groups
// @Tags student_group
// @Produce json
// @Success 200 {object} model.StudentGroupListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /group/students [get]
func (h *StudentGroupHandlers) GetStudentGroups(c *gin.Context) {}

// GetStudentGroupById
// @Summary get student group by id
// @ID get_student_group_by_id
// @Tags student_group
// @Produce json
// @Param id path string true "group_id"
// @Success 200 {object} model.StudentGroup
// @Failure 404 {object} model.ResponseMessage
// @Router /group/students{group_id} [get]
func (h *StudentGroupHandlers) GetStudentGroupById(c *gin.Context) {}

// AddStudentGroup
// @Summary add new course to the course list
// @ID add-student-group
// @Tags student_group
// @Produce json
// @Consume json
// @Param course body model.StudentGroup true "student_group"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /group/students [post]
func (h *StudentGroupHandlers) AddStudentGroup(c *gin.Context) {}

// UpdateStudentGroupById
// @Summary update student group by group id
// @ID update-student-group-by-id
// @Tags student_group
// @Param id path string true "group_id"
// @Success 200 {object} model.StudentGroup
// @Param course body model.StudentGroup true "student_group"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /group/students/{group_id} [put]
func (h *StudentGroupHandlers) UpdateStudentGroupById(c *gin.Context) {}

// DeleteStudentGroupById
// @Summary delete student group by id
// @ID delete-student-group-by-id
// @Tags student_group
// @Param id path string true "group_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /group/students/{group_id} [delete]
func (h *StudentGroupHandlers) DeleteStudentGroupById(c *gin.Context) {}
