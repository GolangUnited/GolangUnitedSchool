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
// @Description создает новый курс
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
// @Description возвращает все курсы
// @ID get-all-courses
// @Tags courses
// @Produce json
// @Success 200 {object} model.CourseList
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses [get]
func (h *CourseHandlers) GetCourses(c *gin.Context) {

}

// GetCourseByName
// @Summary get a course by name
// @ID get-course-by-name
// @Tags courses
// @Produce json
// @Param course_name path string true "course name"
// @Success 200 {object} model.Course
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/name/{course_name} [get]
func (h *CourseHandlers) GetCourseByName(c *gin.Context) {

}

// GetCourseById
// @Summary get a course by ID
// @Description возвращает курс с указанным id
// @ID get-course-by-id
// @Tags courses
// @Produce json
// @Param id path string true "course id"
// @Success 200 {object} model.Course
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/{id} [get]
func (h *CourseHandlers) GetCourseById(c *gin.Context) {

}

// PutCourseById
// @Summary put a course by ID
// @Description изменяет курс с указанным id
// @ID update-course-by-id
// @Tags courses
// @Param id path string true "course id"
// @Param course body model.CourseUpdate true "course"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/{id} [put]
func (h *CourseHandlers) UpdateCourseById(c *gin.Context) {

}

// DeleteCourseById
// @Summary delete a course by ID
// @Description удаляет курс с указанным id
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
