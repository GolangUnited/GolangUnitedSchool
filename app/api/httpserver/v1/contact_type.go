package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type ContactTypeHandlers struct {
	lg      logger.Logger
	usecase usecase.ContactTypeUsecaseInterface
}

func NewContactTypeHandler(lg logger.Logger, usecase usecase.ContactTypeUsecaseInterface) *ContactTypeHandlers {
	return &ContactTypeHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *ContactTypeHandlers) GetContactTypes(c *gin.Context) {

}

func (h *ContactTypeHandlers) GetContactTypeById(c *gin.Context) {

}

func (h *ContactTypeHandlers) AddContactType(c *gin.Context) {

}

func (h *ContactTypeHandlers) UpdateContactType(c *gin.Context) {

}
func (h *ContactTypeHandlers) DeleteContactType(c *gin.Context) {

}
