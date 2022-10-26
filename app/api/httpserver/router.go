package httpserver

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
)

func NewRouter(
	courseHandler *v1.CourseHandlers,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	courseRouter(api, courseHandler)
	personRouter(api)

	return router
}

func courseRouter(
	api *gin.RouterGroup,
	h *v1.CourseHandlers,
) {
	course := api.Group("/course")
	{
		course.GET("", h.SearchCourse)
		course.GET("/:course_id", h.GetCourseByID)
		course.POST("", h.CreateCourse)
		course.PATCH("", h.EditCourse)
		course.PATCH("/:course_id", h.EditCourseByID)
		course.PUT("", h.AddCourse)
		course.DELETE("", h.DeleteCourse)
		course.DELETE("/:course_id", h.DeleteCourseByID)
	}
}

func personRouter(
	api *gin.RouterGroup,
) {
	person := api.Group("/person")
	{
		person.GET("")
		person.GET("/:person_id")
		person.DELETE("/")
		person.DELETE("/:person_id")
	}
}
