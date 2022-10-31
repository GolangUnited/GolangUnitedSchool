package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type GroupContactHandlers struct {
	lg      logger.Logger
	useCase usecase.GroupContactUseCaseInterface
}

func NewGroupContactHandler(
	lg logger.Logger,
	useCase usecase.GroupContactUseCaseInterface,
) *GroupContactHandlers {
	return &GroupContactHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *GroupContactHandlers) GetGroupContacts(c *gin.Context)    {}
func (h *GroupContactHandlers) GetGroupContactById(c *gin.Context) {}
func (h *GroupContactHandlers) AddGroupContact(c *gin.Context)     {}
func (h *GroupContactHandlers) UpdateGroupContact(c *gin.Context)  {}
func (h *GroupContactHandlers) DeleteGroupContact(c *gin.Context)  {}
