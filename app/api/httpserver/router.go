package httpserver

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
)

func NewRouter(
	course *v1.CourseHandlers,
	project *v1.ProjectHandlers,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1/")
	courseRouter(api, course)
	projectRouter(api, project)

	return router
}

func courseRouter(
	api *gin.RouterGroup,
	h *v1.CourseHandlers,
) {
	course := api.Group("/course")
	{
		course.GET("")

	}
}

func projectRouter(api *gin.RouterGroup, h *v1.ProjectHandlers) {
	project := api.Group("/project")
	{
		project.POST("", h.CreateProject)
		project.GET("", h.GetProject)
		project.GET("/:project_id", h.GetProjectByID)
		project.PUT("/:project_id", h.EditProject)
		project.DELETE("/:project_id", h.DeleteProject)
	}
}
