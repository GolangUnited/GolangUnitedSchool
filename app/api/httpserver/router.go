package httpserver

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
)

func NewRouter(
	course *v1.CourseHandlers,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1/")
	courseRouter(api)

	return router
}

func courseRouter(
	api *gin.RouterGroup,
) {
	course := api.Group("/course")
	{
		course.GET("")

	}
}
