package httpserver

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
)

func NewRouter(
	course *v1.CourseHandlers,
	project *v1.ProjectHandlers,
	operationLog *v1.OperationLogHandlers,
	operation *v1.OperationHandlers,
	operationType *v1.OperationTypeHandlers,
	contactType *v1.ContactTypeHandlers,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1/")
	courseRouter(api, course)
	projectRouter(api, project)
	operationLogRouter(api, operationLog)
	operationRouter(api, operation)
	operationTypeRouter(api, operationType)
	contactTypeRouter(api, contactType)

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

func operationLogRouter(api *gin.RouterGroup, h *v1.OperationLogHandlers) {
	operationLog := api.Group("operation_log")
	{
		operationLog.POST("", h.CreateOperationLog)
		operationLog.GET("/:operation_log_id", h.GetOperationLog)
		operationLog.DELETE("/:operation_log_id", h.DeleteOperationLog)
	}
}

func operationRouter(api *gin.RouterGroup, h *v1.OperationHandlers) {
	operation := api.Group("operation")
	{
		operation.POST("", h.CreateOperation)
		operation.GET("", h.GetOperation)
		operation.GET("/:project_id", h.GetOperationByID)
		operation.PUT("/:project_id", h.EditOperation)
		operation.DELETE("/:project_id", h.DeleteOperation)
	}
}

func operationTypeRouter(api *gin.RouterGroup, h *v1.OperationTypeHandlers) {
	operationType := api.Group("operation_type")
	{
		operationType.POST("", h.CreateOperationType)
		operationType.GET("", h.GetOperationType)
		operationType.GET("/:project_id", h.GetOperationTypeByID)
		operationType.PUT("/:project_id", h.EditOperationType)
		operationType.DELETE("/:project_id", h.DeleteOperationType)
	}
}

func contactTypeRouter(api *gin.RouterGroup, h *v1.ContactTypeHandlers) {
	contactType := api.Group("contact_type")
	{
		contactType.POST("", h.CreateContactType)
		contactType.GET("", h.GetContactType)
		contactType.GET("/:project_id", h.GetContactTypeByID)
		contactType.PUT("/:project_id", h.EditContactType)
		contactType.DELETE("/:project_id", h.DeleteContactType)
	}
}
