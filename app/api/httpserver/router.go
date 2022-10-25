package httpserver

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
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
