package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type RoleHandlers struct {
	lg      logger.Logger
	usecase usecase.RoleUsecaseInterface
}

func NewRoleHandlers(lg logger.Logger,
	usecase usecase.RoleUsecaseInterface) *RoleHandlers {
	return &RoleHandlers{lg: lg, usecase: usecase}
}

// @Summary get a role by ID
// @ID get-role-by-id
// @Tags roles
// @Produce json
// @Param id path string true "role id"
// @Success 200 {object} model.Role
// @Failure 404 {object} model.ResponseMessage
// @Router /role/{id} [get]
func (h *RoleHandlers) GetRoleById(c *gin.Context) {}

// @Summary get all items in the role list
// @ID get-all-roles
// @Tags roles
// @Produce json
// @Success 200 {object} []model.Role
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /role [get]
func (h *RoleHandlers) GetRoles(c *gin.Context) {}

// @Summary add new role to the course list
// @ID create-role
// @Tags roles
// @Produce json
// @Consume json
// @Param role body model.RoleCU true "role"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /role [post]
func (h *RoleHandlers) AddRole(c *gin.Context) {}

// @Summary update a role by ID
// @ID update-role-by-id
// @Tags roles
// @Param id path string true "role id"
// @Param role body model.RoleCU true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /role/{id} [post]
func (h *RoleHandlers) UpadateRoleById(c *gin.Context) {}

// @Summary put a role by ID
// @ID put-role-by-id
// @Tags roles
// @Param id path string true "role id"
// @Param role body model.Role true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /role/{id} [put]
func (h *RoleHandlers) PutRoleById(c *gin.Context) {}

// @Summary delete a role by ID
// @ID delete-role-by-id
// @Tags roles
// @Param id path string true "role id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /role/{id} [delete]
func (h *RoleHandlers) DeleteRoleById(c *gin.Context) {}
