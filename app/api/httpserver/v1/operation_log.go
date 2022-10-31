package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationLogHandlers struct {
	lg      logger.Logger
	usecase usecase.OperationLogUseCaseInterface
}

func NewOperationLogHandler(lg logger.Logger, usecase usecase.OperationLogUseCaseInterface) *OperationLogHandlers {
	return &OperationLogHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *OperationLogHandlers) CreateOperationLog(c *gin.Context) {

}

func (h *OperationLogHandlers) GetOperationLog(c *gin.Context) {

}

func (h *OperationLogHandlers) DeleteOperationLog(c *gin.Context) {

}
