package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type ContactTypeHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseUsecaseInterface
}

func NewContactTypeHandler(
	lg logger.Logger,
	useCase usecase.ContactTypeUseCaseInterface,
) *ContactTypeHandlers {
	return &ContactTypeHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// GetContactTypes
// @Summary get all contact types
// @Description возвращает все типы контактов
// @ID get-all-contact-types
// @Tags contactTypes
// @Produce json
// @Success 200 {object} model.ContactTypesListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/contacts/types [get]
func (h *ContactTypeHandlers) GetContactTypes(c *gin.Context) {}

// GetContactType
// @Summary get contact type
// @Description возвращает тип контакта с указанным contact_type_id
// @ID get-contact-type
// @Param id path string true "contact_type_id"
// @Tags contactTypes
// @Produce json
// @Success 200 {object} model.ContactType
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/contacts/types/{contact_type_id} [get]
func (h *ContactTypeHandlers) GetContactType(c *gin.Context) {}

// AddContactType
// @Summary get contact type
// @Description добавляет новый тип контакта
// @ID add-contact-type
// @Param contact_type body model.NewContactTypeDto true "contact_type"
// @Tags contactTypes
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/contacts/types [post]
func (h *ContactTypeHandlers) AddContactType(c *gin.Context) {}

// UpdateContactType
// @Summary update contact type
// @Description изменяет тип контакта
// @ID update-contact-type
// @Param contact_type body model.ContactType true "contact_type"
// @Tags contactTypes
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/contacts/types [put]
func (h *ContactHandlers) UpdateContactType(c *gin.Context) {}

// DeleteContactType
// @Summary delete contact type
// @Description удаляет тип контакта с указанным contact_type_id
// @ID delete-contact-type
// @Param id path string true "contact_type_id"
// @Tags contactTypes
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/contacts/types/{:contact_type_id} [delete]
func (h *ContactTypeHandlers) DeleteContactType(c *gin.Context) {}
