package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type HomeworkHandlers struct {
	lg      logger.Logger
	usecase usecase.HomeworkUsecaseInterface
}

func NewHomeworkHandler(
	lg logger.Logger,
	usecase usecase.HomeworkUsecaseInterface,
) *HomeworkHandlers {
	return &HomeworkHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *HomeworkHandlers) GetHomeworks(c *gin.Context) {
}

func (h *HomeworkHandlers) GetHomeworksByLectureId(c *gin.Context) {
}

func (h *HomeworkHandlers) GetHomeworkById(c *gin.Context) {
}

func (h *HomeworkHandlers) AddHomework(c *gin.Context) {
}

func (h *HomeworkHandlers) UpdateHomework(c *gin.Context) {
}

func (h *HomeworkHandlers) DeleteHomework(c *gin.Context) {
}
