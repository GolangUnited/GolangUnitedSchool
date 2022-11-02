package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type CourseLectureHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseStatusUseCaseInterface
}

func NewCourseLectureHandler(
	lg logger.Logger,
	useCase usecase.CourseLectureUseCaseInterface,
) *CourseLectureHandlers {
	return &CourseLectureHandlers{
		lg:      lg,
		useCase: useCase,
	}
}
func (h *CourseLectureHandlers) GetCourseLectures(c *gin.Context)       {}
func (h *CourseLectureHandlers) GetCourseLectureById(c *gin.Context)    {}
func (h *CourseLectureHandlers) AddCourseLecture(c *gin.Context)        {}
func (h *CourseLectureHandlers) UpdateCourseLectureById(c *gin.Context) {}
func (h *CourseLectureHandlers) DeleteCourseLectureById(c *gin.Context) {}
