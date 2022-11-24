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
// @Description возвращает список всех пользователей
// @ID get-persons
// @Tags persons
// @Produce json
// @Success 200 {object} model.PersonListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /persons [get]
func (h *PersonHandlers) GetPersons(c *gin.Context) {}

// AddNewPerson
// @Summary add new person to database
// @Description добавляет нового пользователя
// @ID add-new-person
// @Tags persons
// @Produce json
// @Consume json
// @Param course body model.NewPersonDto true "person"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons [post]
func (h *PersonHandlers) AddNewPerson(c *gin.Context) {}

// GetPersonById
// @Summary get person by id
// @Description возвращает данные о пользователе с указанным id
// @ID get-person-by-id
// @Tags persons
// @Produce json
// @Param id path string true "person_id"
// @Success 200 {object} model.Person
// @Failure 404 {object} model.ResponseMessage
// @Router /persons/{person_id} [get]
func (h *PersonHandlers) GetPersonById(c *gin.Context) {}

// PutPersonById
// @Summary update person by id
// @ID update-person-by-id
// @Description изменяет данные пользователя с указанным id
// @Tags persons
// @Produce json
// @Param id path string true "person_id"
// @Param person body model.Person true "person"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/{person_id} [put]
func (h *PersonHandlers) PutPersonById(c *gin.Context) {}

// UpdatePersonById
// @Summary update person by person id
// @Description изменяет пользователя с указанным id, не затрагивая всех полей
// @ID update-person-by-id
// @Tags persons
// @Param id path string true "person_id"
// @Param update_person body model.UpdatePersonDto true "update_person"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/{person_id} [post]
func (h *PersonHandlers) UpdatePersonById(c *gin.Context) {}

// DeletePersonById
// @Summary delete person by id
// @Description удаляет пользователя с указанным id<
// @ID delete-person-by-id
// @Tags persons
// @Param id path string true "person_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/{person_id} [delete]
func (h *PersonHandlers) DeletePersonById(c *gin.Context) {}
