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

// GetRoleById
// @Summary get a role by ID
// @Description возвращает роль с указанным id
// @ID get-role-by-id
// @Tags roles
// @Produce json
// @Param id path string true "role id"
// @Success 200 {object} model.Role
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /roles/{id} [get]
func (h *RoleHandlers) GetRoleById(c *gin.Context) {}

// GetRoles
// @Summary get all items in the role list
// @Description возвращает все роли
// @ID get-all-roles
// @Tags roles
// @Produce json
// @Success 200 {object} []model.Role
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /roles [get]
func (h *RoleHandlers) GetRoles(c *gin.Context) {}

// AddRole
// @Summary add new role to the course list
// @Description добавляет новую роль
// @ID create-role
// @Tags roles
// @Produce json
// @Consume json
// @Param role body model.RoleCU true "role"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /roles [post]
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
// @Router /roles/{id} [post]
func (h *RoleHandlers) UpadateRoleById(c *gin.Context) {}

// @Summary put a role by ID
// @Description изменяет роль с указанным id
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
// @Router /roles/{id} [put]
func (h *RoleHandlers) PutRoleById(c *gin.Context) {}

// DeleteRoleById
// @Summary delete a role by ID
// @Description удаляет роль с указанным id
// @ID delete-role-by-id
// @Tags roles
// @Param id path string true "role id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /roles/{id} [delete]
func (h *RoleHandlers) DeleteRoleById(c *gin.Context) {}
