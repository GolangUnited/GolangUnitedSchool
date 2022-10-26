package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type PersonHandlers struct {
	lg      logger.Logger
	usecase usecase.PersonUsecaseIntarface
}

func NewPersonHandler(lg logger.Logger,
	usecase usecase.PersonUsecaseIntarface,
) *PersonHandlers {
	return &PersonHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *PersonHandlers) AddNewPerson(c *gin.Context) {

}

func (h *PersonHandlers) GetPersonByID(c *gin.Context) {

}
func (h *PersonHandlers) SearchPerson(c *gin.Context) {

}

func (h *PersonHandlers) EditPerson(c *gin.Context) {

}

func (h *PersonHandlers) EditPersonById(c *gin.Context) {

}

func (h *PersonHandlers) DeletePerson(c *gin.Context) {

}
func (h *PersonHandlers) DeletePersonByID(c *gin.Context) {

}
