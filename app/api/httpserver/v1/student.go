package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentHandlers struct {
	lg      logger.Logger
	useCase usecase.StudentUseCaseInterface
}

func NewStudentHandler(
	lg logger.Logger,
	useCase usecase.StudentUseCaseInterface,
) *StudentHandlers {
	return &StudentHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *StudentHandlers) AddStudent(c *gin.Context) {}

func (h *StudentHandlers) DeleteStudentByStudentId(c *gin.Context) {}

func (h *StudentHandlers) EditStudentByStudentId(c *gin.Context) {}

func (h *StudentHandlers) GetStudentByStudentId(c *gin.Context) {}

func (h *StudentHandlers) GetStudents(c *gin.Context) {}
