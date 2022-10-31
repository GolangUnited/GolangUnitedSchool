package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationTypeHandlers struct {
	lg      logger.Logger
	usecase usecase.OperationTypeUseCaseInterface
}

func NewOperationTypeHandler(lg logger.Logger, usecase usecase.OperationTypeUseCaseInterface) *OperationTypeHandlers {
	return &OperationTypeHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *OperationTypeHandlers) CreateOperationType(c *gin.Context) {

}

func (h *OperationTypeHandlers) GetOperationType(c *gin.Context) {

}

func (h *OperationTypeHandlers) GetOperationTypeByID(c *gin.Context) {

}

func (h *OperationTypeHandlers) EditOperationType(c *gin.Context) {

}

func (h *OperationTypeHandlers) DeleteOperationType(c *gin.Context) {

}
