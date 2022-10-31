package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type ContactTypeHandlers struct {
	lg      logger.Logger
	usecase usecase.ContactTypeUseCaseInterface
}

func NewContactTypeHandler(lg logger.Logger, usecase usecase.ContactTypeUseCaseInterface) *ContactTypeHandlers {
	return &ContactTypeHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *ContactTypeHandlers) CreateContactType(c *gin.Context) {

}

func (h *ContactTypeHandlers) GetContactType(c *gin.Context) {

}

func (h *ContactTypeHandlers) GetContactTypeByID(c *gin.Context) {

}

func (h *ContactTypeHandlers) EditContactType(c *gin.Context) {

}

func (h *ContactTypeHandlers) DeleteContactType(c *gin.Context) {

}
