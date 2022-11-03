package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationTypeHandlers struct {
	lg      logger.Logger
	usecase usecase.OperationTypeUsecaseInterface
}

func NewOperationTypeHandler(lg logger.Logger, usecase usecase.OperationTypeUsecaseInterface) *OperationTypeHandlers {
	return &OperationTypeHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *OperationTypeHandlers) GetOperationTypes(c *gin.Context) {

}

func (h *OperationTypeHandlers) GetOperationTypeById(c *gin.Context) {

}

func (h *OperationTypeHandlers) AddOperationType(c *gin.Context) {

}

func (h *OperationTypeHandlers) UpdateOperationType(c *gin.Context) {

}
func (h *OperationTypeHandlers) DeleteOperationType(c *gin.Context) {

}
