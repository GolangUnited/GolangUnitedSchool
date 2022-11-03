package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationLogHandlers struct {
	lg      logger.Logger
	usecase usecase.OperationLogUsecaseInterface
}

func NewOperationLogHandler(lg logger.Logger, usecase usecase.OperationLogUsecaseInterface) *OperationLogHandlers {
	return &OperationLogHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *OperationLogHandlers) GetOperationLogById(c *gin.Context) {

}

func (h *OperationLogHandlers) AddOperationLog(c *gin.Context) {

}

func (h *OperationLogHandlers) DeleteOperationLog(c *gin.Context) {

}
