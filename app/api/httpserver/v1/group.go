package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type GroupHandlers struct {
	lg      logger.Logger
	usecase usecase.GroupUsecaseInterface
}

func NewGroup(lg logger.Logger, u usecase.GroupUsecaseInterface) *GroupHandlers {
	return &GroupHandlers{
		lg:      lg,
		usecase: u,
	}
}

// @Summary add new item to the group list
// @ID create-group
// @Tags groups
// @Produce json
// @Consume json
// @Param group body model.GroupCU true "group"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups [post]
func CreateGroup(c *gin.Context) {}

// @Summary get a group by ID
// @ID get-group-by-id
// @Tags groups
// @Produce json
// @Param id path string true "group id"
// @Success 200 {object} model.Group
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/{id} [get]
func GetGroupById(c *gin.Context) {}

// @Summary get all items in the group list
// @ID get-all-groups
// @Tags groups
// @Produce json
// @Success 200 {object} []model.Group
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups [get]
func GetGroups(c *gin.Context) {}

// @Summary update a group by ID
// @ID update-group-by-id
// @Tags groups
// @Param id path string true "group id"
// @Param group body model.GroupCU true "group"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/{id} [post]
func UpdateGroupById(c *gin.Context) {}

// @Summary put a group by ID
// @ID update-group-by-id
// @Tags groups
// @Param id path string true "group id"
// @Param group body model.GroupCU true "group"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/{id} [put]
func PutGroupById(c *gin.Context) {}

// @Summary delete a group by ID
// @ID delete-group-by-id
// @Tags groups
// @Param id path string true "group id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/{id} [delete]
func DeleteGroupById(c *gin.Context) {}
