package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type GroupContactHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseStatusUseCaseInterface
}

func NewGroupContactHandler(
	lg logger.Logger,
	useCase usecase.CourseStatusUseCaseInterface,
) *GroupContactHandlers {
	return &GroupContactHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// GetGroupContacts
// @Summary get all group contacts
// @ID get-all-groups-contacts
// @Description возвращает все контакты всех группы
// @Tags groupContacts
// @Produce json
// @Consume json
// @Success 201 {object} model.GroupContactsListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/contacts [get]
func (h *GroupContactHandlers) GetGroupContacts(c *gin.Context) {}

// GetGroupContactsByGroupId
// @Summary get group contacts by  group id
// @Description возвращает все контакты группы с указанным group_id
// @ID get-group-contacts-by-id
// @Tags groupContacts
// @Produce json
// @Param id path string true "group_id"
// @Success 200 {object} model.GroupContactsListDto
// @Failure 404 {object} model.ResponseMessage
// @Router /groups/{group_id}/contacts [get]
func (h *GroupContactHandlers) GetGroupContactsByGroupId(c *gin.Context) {}

// AddGroupContact
// @Summary add new group contact
// @ID add-group-contact
// @Description добавляет новый контакт группы
// @Tags groupContacts
// @Produce json
// @Consume json
// @Param group_contact body model.GroupContactsAddDto true "group contact"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/contacts [post]
func (h *GroupContactHandlers) AddGroupContact(c *gin.Context) {}

// UpdateGroupContact
// @Summary update group contact
// @Description изменяет контакт группы
// @Tags groupContacts
// @ID update-group-contact-by-id
// @Param id path string true "group_contact_id"
// @Param group_contact body model.GroupContactsAddDto true "group contact"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /groups/contacts/{group_contact_id} [put]
func (h *GroupContactHandlers) UpdateGroupContact(c *gin.Context) {}

// DeleteGroupContact
// @Summary delete group contact by id
// @Description удаляет контакт группы
// @Tags groupContacts
// @ID delete-group-contact-by-id
// @Param id path string true "group_contact_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/contacts/{group_contact_id} [delete]
func (h *GroupContactHandlers) DeleteGroupContact(c *gin.Context) {}

// GetGroupContact
// @Summary get group contact
// @Description возвращает контакт group_contact_id группы с group_id
// @ID get-group-contact
// @Tags groupContacts
// @Produce json
// @Param id path string true "group_id"
// @Param id path string true "group_contact_id"
// @Success 200 {object} model.GroupContact
// @Failure 404 {object} model.ResponseMessage
// @Router /groups/{group_id}/contacts/{group_contact_id} [get]
func (h *GroupContactHandlers) GetGroupContact(c *gin.Context) {}
