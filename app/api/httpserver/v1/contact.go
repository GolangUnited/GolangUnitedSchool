package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type ContactHandler struct {
	lg      logger.Logger
	usecase usecase.ContactUsecaseInterface
}

func NewContactHandler(lg logger.Logger, usecase usecase.ContactUsecaseInterface) *ContactHandler {
	return &ContactHandler{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *ContactHandler) GetContactsByPersonId(c *gin.Context) {

}

func (h *ContactHandler) AddNewPersonContact(c *gin.Context) {

}

func (h *ContactHandler) DeletePersonContact(c *gin.Context) {

}

func (h *ContactHandler) UpdatePersonContact(c *gin.Context) {

}
