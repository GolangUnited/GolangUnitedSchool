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
// @Tags courses, courseStatuses
// @Produce json
// @Success 200 {object} model.CourseStatusesListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses [get]
func (h *CourseStatusHandlers) GetCourseStatuses(c *gin.Context) {}

// GetCourseStatusById
// @Summary get course status by id
// @Description возвращает статус курса с status_id
// @ID get-course-status
// @Tags courses, courseStatuses
// @Produce json
// @Param id path string true "status_id"
// @Success 200 {object} model.CourseStatus
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses/{status_id} [get]
func (h *CourseStatusHandlers) GetCourseStatusById(c *gin.Context) {}

// AddCourseStatus
// @Summary add new course status
// @Description создает новый статус курса
// @ID add-course-status
// @Tags courses, courseStatuses
// @Produce json
// @Param new_course_status body model.CourseStatus true "new_course_status"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses [post]
func (h *CourseStatusHandlers) AddCourseStatus(c *gin.Context) {}

// UpdateCourseStatusById
// @Summary update course status
// @Description изменяет статус курса с указанным status_id
// @ID put-course-status
// @Tags courses, courseStatuses
// @Produce json
// @Param id path string true "status_id"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses/{status_id} [put]
func (h *CourseStatusHandlers) UpdateCourseStatusById(c *gin.Context) {}

// DeleteCourseStatusById
// @Summary delete course status
// @Description удаляет статус курса с указанным status_id
// @ID delete-course-status
// @Tags courses, courseStatuses
// @Produce json
// @Param id path string true "status_id"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/statuses/{status_id} [delete]
func (h *CourseStatusHandlers) DeleteCourseStatusById(c *gin.Context) {}
