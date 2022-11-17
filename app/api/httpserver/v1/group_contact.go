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
// @ID get-all-group-contacts
// @Tags group_contact
// @Produce json
// @Consume json
// @Success 201 {object} model.GroupContactListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /group/contact [get]
func (h *GroupContactHandlers) GetGroupContacts(c *gin.Context) {}

// GetGroupContactById
// @Summary get group contact by id
// @ID get-group-contact-by-id
// @Tags group_contact
// @Produce json
// @Param id path string true "group_contact_id"
// @Success 200 {object} model.GroupContact
// @Failure 404 {object} model.ResponseMessage
// @Router /group/contact/{group_contact_id} [get]
func (h *GroupContactHandlers) GetGroupContactById(c *gin.Context) {}

// AddGroupContact
// @Summary add new group contact
// @ID add-group-contact
// @Tags group_contact
// @Produce json
// @Consume json
// @Param course body model.GroupContactAddDto true "course"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /course [post]
func (h *GroupContactHandlers) AddGroupContact(c *gin.Context) {}

// UpdateGroupContact
// @Summary update group contact
// @ID update-group-contact-by-id
// @Tags group_contact
// @Param id path string true "group_contact_id"
// @Param course body model.GroupContactAddDto true "group contact"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /group/contact/{group_contact_id} [put]
func (h *GroupContactHandlers) UpdateGroupContact(c *gin.Context) {}

// DeleteGroupContact
// @Summary delete group contact by id
// @ID delete-group-contact-by-id
// @Tags group_contact
// @Param id path string true "group_contact_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /course/contact/{group_contact_id} [delete]
func (h *GroupContactHandlers) DeleteGroupContact(c *gin.Context) {}
