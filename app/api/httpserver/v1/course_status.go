package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type CourseStatusHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseStatusUseCaseInterface
}

func NewCourseStatusHandler(
	lg logger.Logger,
	useCase usecase.CourseStatusUseCaseInterface,
) *CourseStatusHandlers {
	return &CourseStatusHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// GetCourseStatuses
// @Summary get all course statuses
// @Description возвращает список всех статусов курса
// @ID get-all-course-statuses
// @Tags courseStatuses
// @Produce json
// @Success 200 {object} model.CourseStatusesListDto
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses [get]
func (h *CourseStatusHandlers) GetCourseStatuses(c *gin.Context) {}

// GetCourseStatusById
// @Summary get course status by id
// @Description возвращает статус курса с status_id
// @ID get-course-status
// @Tags courseStatuses
// @Produce json
// @Param id path string true "status_id"
// @Success 200 {object} model.CourseStatus
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses/{status_id} [get]
func (h *CourseStatusHandlers) GetCourseStatusById(c *gin.Context) {}

// AddCourseStatus
// @Summary add new course status
// @Description создает новый статус курса
// @ID add-course-status
// @Tags courseStatuses
// @Produce json
// @Param new_course_status body model.NewCourseStatusDto true "new_course_status"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses [post]
func (h *CourseStatusHandlers) AddCourseStatus(c *gin.Context) {}

// UpdateCourseStatusById
// @Summary update course status
// @Description изменяет статус курса с указанным status_id
// @ID put-course-status
// @Tags courseStatuses
// @Produce json
// @Param id path string true "status_id"
// @Param update_course_status body model.NewCourseStatusDto true "course_status"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses/{status_id} [post]
func (h *CourseStatusHandlers) UpdateCourseStatusById(c *gin.Context) {}

// PutCourseStatusById
// @Summary put course status by id
// @Description изменяет статус курса с указанным id с доступом ко всем полям
// @ID зге-operation-type-by-id
// @Tags courseStatuses
// @Param id path string true "course_status_id"
// @Param operation_type body model.CourseStatus true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses/{id} [put]
func (h *CourseStatusHandlers) PutCourseStatusById(c *gin.Context) {}

// DeleteCourseStatusById
// @Summary delete course status
// @Description удаляет статус курса с указанным status_id
// @ID delete-course-status
// @Tags courseStatuses
// @Produce json
// @Param id path string true "status_id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses/{status_id} [delete]
func (h *CourseStatusHandlers) DeleteCourseStatusById(c *gin.Context) {}
