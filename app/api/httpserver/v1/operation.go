package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationHandlers struct {
	lg      logger.Logger
	usecase usecase.OperationUseCaseInterface
}

func NewOperationHandler(lg logger.Logger, usecase usecase.OperationUseCaseInterface) *OperationHandlers {
	return &OperationHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *OperationHandlers) CreateOperation(c *gin.Context) {

}

func (h *OperationHandlers) GetOperation(c *gin.Context) {

}

func (h *OperationHandlers) GetOperationByID(c *gin.Context) {

}

func (h *OperationHandlers) EditOperation(c *gin.Context) {

}

func (h *OperationHandlers) DeleteOperation(c *gin.Context) {

}
