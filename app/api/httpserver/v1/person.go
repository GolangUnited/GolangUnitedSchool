package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type PersonHandlers struct {
	lg      logger.Logger
	useCase usecase.PersonUseCaseInterface
}

func NewPersonHandler(lg logger.Logger,
	useCase usecase.PersonUseCaseInterface,
) *PersonHandlers {
	return &PersonHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *PersonHandlers) GetPersons(c *gin.Context)       {}
func (h *PersonHandlers) AddNewPerson(c *gin.Context)     {}
func (h *PersonHandlers) GetPersonById(c *gin.Context)    {}
func (h *PersonHandlers) EditPersonById(c *gin.Context)   {}
func (h *PersonHandlers) DeletePersonById(c *gin.Context) {}
