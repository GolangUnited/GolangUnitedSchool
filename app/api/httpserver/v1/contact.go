package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type ContactHandlers struct {
	lg      logger.Logger
	useCase usecase.ContactUseCaseInterface
}

func NewContactHandler(
	lg logger.Logger,
	useCase usecase.ContactUseCaseInterface,
) *ContactHandlers {
	return &ContactHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// GetContacts
// @Summary get all contacts
// @Description возвращает все контакты всех пользователей
// @ID get-all-contacts
// @Tags persons, contacts
// @Produce json
// @Success 200 {object} model.ContactsListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/contacts [get]
func (h *ContactHandlers) GetContacts(c *gin.Context) {}

// GetPersonContacts
// @Summary get all person's contacts
// @Description возвращает все контакты пользователя с указанным person_id
// @ID get-person_contacts
// @Tags persons, contacts
// @Produce json
// @Param id path string true "person_id"
// @Success 200 {object} model.ContactsListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/{person_id}/contacts [get]
func (h *ContactHandlers) GetPersonContacts(c *gin.Context) {}

// GetPersonContact
// @Summary get person's contact by contact_id
// @Description возвращает контакт c указанным contact_id пользователя с указанным person_id
// @ID get-person_contact
// @Tags persons, contacts
// @Produce json
// @Param id path string true "person_id"
// @Param id path string true "contact_id"
// @Success 200 {object} model.Contact
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/{person_id}/contacts/{contact_id} [get]
func (h *ContactHandlers) GetPersonContact(c *gin.Context) {}

// AddPersonContact
// @Summary add new person contact
// @Description создает новый пользовательский контакт
// @ID add-person_contact
// @Tags persons, contacts
// @Produce json
// @Param contact body model.Contact true "contact"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/contacts [post]
func (h *ContactHandlers) AddPersonContact(c *gin.Context) {}

// UpdatePersonContact
// @Summary add new person contact
// @Description изменяет пользовательский контакт
// @ID put-person_contact
// @Tags persons, contacts
// @Produce json
// @Param contact body model.Contact true "contact"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/contacts [put]
func (h *ContactHandlers) UpdatePersonContact(c *gin.Context) {}

// DeletePersonContact
// @Summary delete person's contact by contact_id
// @Description удаляет контакт c указанным contact_id
// @ID delete-person_contact
// @Tags persons, contacts
// @Produce json
// @Param id path string true "person_id"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /persons/contacts/{contact_id} [delete]
func (h *ContactHandlers) DeletePersonContact(c *gin.Context) {}
