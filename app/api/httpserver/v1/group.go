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

func GetGroupById(c *gin.Context)    {}
func GetGroups(c *gin.Context)       {}
func CreateGroup(c *gin.Context)     {}
func UpdateGroupById(c *gin.Context) {}
func PutGroupById(c *gin.Context)    {}
func DeleteGroupById(c *gin.Context) {}
