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

func (h *RoleHandlers) GetRoleById(c *gin.Context)     {}
func (h *RoleHandlers) GetRoles(c *gin.Context)        {}
func (h *RoleHandlers) AddRole(c *gin.Context)         {}
func (h *RoleHandlers) UpadateRoleById(c *gin.Context) {}
func (h *RoleHandlers) PutRoleById(c *gin.Context)     {}
func (h *RoleHandlers) DeleteRoleById(c *gin.Context)  {}
