package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type ProjectHandlers struct {
	lg      logger.Logger
	usecase usecase.ProjectUseCaseInterface
}

func NewProjectHandler(lg logger.Logger, usecase usecase.ProjectUseCaseInterface) *ProjectHandlers {
	return &ProjectHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *ProjectHandlers) CreateProject(c *gin.Context) {

}

func (h *ProjectHandlers) GetProject(c *gin.Context) {

}

func (h *ProjectHandlers) GetProjectByID(c *gin.Context) {

}

func (h *ProjectHandlers) EditProject(c *gin.Context) {

}

func (h *ProjectHandlers) DeleteProject(c *gin.Context) {

}
