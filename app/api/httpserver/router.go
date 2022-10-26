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

//func NewPersonRouter(
//	person *v1.PersonHandlers,
//) *gin.Engine {
//	router := gin.Default()
//
//	api := router.Group("/api/v1/")
//	personRouter(api)
//
//	return router
//}

func courseRouter(
	api *gin.RouterGroup,
) {
	course := api.Group("/course")
	{
		course.GET("")

	}
}

func personRouter(
	api *gin.RouterGroup,
) {
	person := api.Group("/person")
	{
		person.GET("")

	}
}
