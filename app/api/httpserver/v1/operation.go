package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationHandlers struct {
	lg      logger.Logger
	usecase usecase.OperationUsecaseInterface
}

func NewOperationHandler(lg logger.Logger, usecase usecase.OperationUsecaseInterface) *OperationHandlers {
	return &OperationHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *OperationHandlers) GetOperations(c *gin.Context) {

}

func (h *OperationHandlers) GetOperationById(c *gin.Context) {

}

func (h *OperationHandlers) AddOperation(c *gin.Context) {

}

func (h *OperationHandlers) UpdateOperation(c *gin.Context) {

}
func (h *OperationHandlers) DeleteOperation(c *gin.Context) {

}
