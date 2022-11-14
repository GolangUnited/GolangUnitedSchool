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

// GetPersons
// @Summary get all person from database
// @ID get-persons
// @Tags person
// @Produce json
// @Success 200 {object} model.PersonListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /person [get]
func (h *PersonHandlers) GetPersons(c *gin.Context) {}

// AddNewPerson
// @Summary add new person to database
// @ID add-new-person
// @Tags person
// @Produce json
// @Consume json
// @Param course body model.Person true "person"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /person [post]
func (h *PersonHandlers) AddNewPerson(c *gin.Context) {}

// GetPersonById
// @Summary get person by id
// @ID get-person-by-id
// @Tags person
// @Produce json
// @Param id path string true "person id"
// @Success 200 {object} model.Person
// @Failure 404 {object} model.ResponseMessage
// @Router /person/{id} [get]
func (h *PersonHandlers) GetPersonById(c *gin.Context) {}

// UpdatePersonById
// @Summary update person by id
// @ID update-person-by-id
// @Tags person
// @Param id path string true "person id"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /person/{id} [post]
func (h *PersonHandlers) UpdatePersonById(c *gin.Context) {}

// DeletePersonById
// @Summary delete person by id
// @ID delete-person-by-id
// @Tags person
// @Param id path string true "person id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /person/{id} [delete]
func (h *PersonHandlers) DeletePersonById(c *gin.Context) {}
