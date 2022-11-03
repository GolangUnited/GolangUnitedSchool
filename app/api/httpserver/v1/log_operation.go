package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type LogOperationHandlers struct {
	lg      logger.Logger
	usecase usecase.LogOperationUsecaseInterface
}

func NewLogOperationHandler(lg logger.Logger, usecase usecase.LogOperationUsecaseInterface) *LogOperationHandlers {
	return &LogOperationHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *LogOperationHandlers) GetLogOperationById(c *gin.Context) {

}

func (h *LogOperationHandlers) AddLogOperation(c *gin.Context) {

}

func (h *LogOperationHandlers) DeleteLogOperation(c *gin.Context) {

}
