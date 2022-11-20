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

// CreateGroup
// @Summary add new item to the group list
// @Description создает новую группу
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

// GetGroupById
// @Summary get a group by ID
// @Description возвращает группу с указанным id
// @ID get-group-by-id
// @Tags groups
// @Produce json
// @Param id path string true "group id"
// @Success 200 {object} model.Group
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/{id} [get]
func GetGroupById(c *gin.Context) {}

// GetGroups
// @Summary get all items in the group list
// @Description возвращает все группы
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
// @Description изменяет группу с указанным id
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

// DeleteGroupById
// @Summary delete a group by ID
// @Description удаляет группу с указанным id
// @ID delete-group-by-id
// @Tags groups
// @Param id path string true "group id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/{id} [delete]
func DeleteGroupById(c *gin.Context) {}
