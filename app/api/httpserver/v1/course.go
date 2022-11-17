package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type CourseHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseUsecaseInterface
}

func NewCourseHandler(
	lg logger.Logger,
	useCase usecase.CourseUsecaseInterface,
) *CourseHandlers {
	return &CourseHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// AddCourse
// @Summary add new course to the course list
// @ID create-course
// @Tags courses
// @Produce json
// @Consume json
// @Param course body model.Course true "course"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses [post]
func (h *CourseHandlers) AddCourse(c *gin.Context) {
	var course model.Course
	err := c.ShouldBind(&course)
	if err != nil {
		h.lg.Error(err, "couldn't bind course into data")
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseMessage{
			Status:  "error",
			Message: "bad request",
		})
		return
	}
}

// GetCourses
// @Summary get all items in the course list
// @ID get-all-courses
// @Tags courses
// @Produce json
// @Success 200 {object} model.CourseList
// @Failure 500 {object} model.ResponseMessage
// @Router /courses [get]
func (h *CourseHandlers) GetCourses(c *gin.Context) {

}

// GetCourseById
// @Summary get a course by ID
// @ID get-course_by_id
// @Tags courses
// @Produce json
// @Param id path string true "course id"
// @Success 200 {object} model.Course
// @Failure 404 {object} model.ResponseMessage
// @Router /courses/{id} [get]
func (h *CourseHandlers) GetCourseById(c *gin.Context) {

}

// UpdateCourseById
// @Summary update a course by ID
// @ID update-course-by-id
// @Tags courses
// @Param id path string true "course id"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /courses/{id} [post]
func (h *CourseHandlers) UpdateCourseById(c *gin.Context) {

}

// DeleteCourseById
// @Summary delete a course by ID
// @ID delete-course-by-id
// @Tags courses
// @Param id path string true "course id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/{id} [delete]
func (h *CourseHandlers) DeleteCourseById(c *gin.Context) {

}
