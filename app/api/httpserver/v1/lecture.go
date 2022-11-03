package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type LectureHandlers struct {
	lg      logger.Logger
	usecase usecase.LectureUsecaseInterface
}

func NewLectureHandler(
	lg logger.Logger,
	usecase usecase.LectureUsecaseInterface,
) *LectureHandlers {
	return &LectureHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *LectureHandlers) GetLectures(c *gin.Context) {
}

func (h *LectureHandlers) GetLectureById(c *gin.Context) {
}

func (h *LectureHandlers) AddLecture(c *gin.Context) {
}

func (h *LectureHandlers) UpdateLecture(c *gin.Context) {
}

func (h *LectureHandlers) DeleteLecture(c *gin.Context) {
}
