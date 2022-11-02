package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentGroupHandlers struct {
	lg      logger.Logger
	useCase usecase.StudentGroupUseCaseInterface
}

func NewStudentGroupHandler(
	lg logger.Logger,
	useCase usecase.StudentGroupUseCaseInterface,
) *StudentGroupHandlers {
	return &StudentGroupHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *StudentGroupHandlers) GetStudentGroups(c *gin.Context)       {}
func (h *StudentGroupHandlers) GetStudentGroupById(c *gin.Context)    {}
func (h *StudentGroupHandlers) AddStudentGroup(c *gin.Context)        {}
func (h *StudentGroupHandlers) UpdateStudentGroupbyId(c *gin.Context) {}
func (h *StudentGroupHandlers) DeleteStudentGroup(c *gin.Context)     {}
