package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentHomeworkHandlers struct {
	lg      logger.Logger
	usecase usecase.StudentHomeworkUsecaseInterface
}

func NewStudentHomeworkHandler(
	lg logger.Logger,
	usecase usecase.StudentHomeworkUsecaseInterface,
) *StudentHomeworkHandlers {
	return &StudentHomeworkHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *StudentHomeworkHandlers) GetStudentHomeworks(c *gin.Context) {
}

func (h *StudentHomeworkHandlers) GetStudentHomeworksByStudentId(c *gin.Context) {
}

func (h *StudentHomeworkHandlers) GetStudentHomeworkById(c *gin.Context) {
}

func (h *StudentHomeworkHandlers) AddStudentHomework(c *gin.Context) {
}

func (h *StudentHomeworkHandlers) UpdateStudentHomework(c *gin.Context) {
}

func (h *StudentHomeworkHandlers) DeleteStudentHomework(c *gin.Context) {
}
