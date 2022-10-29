package httpserver

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
)

func NewRouter(
	handlers *v1.Handlers,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	courseRouter(api, handlers)
	personRouter(api, handlers)

	return router
}

func courseRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	course := api.Group("/course")
	{
		course.GET("")
		course.GET("/:course_id")
	}
}

func personRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	person := api.Group("/person")
	{
		person.GET("")
		person.GET("/:person_id")
		person.DELETE("/")
		person.DELETE("/:person_id")
	}
}
