package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type CourseStatusHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseStatusUseCaseInterface
}

func NewCourseStatusHandler(
	lg logger.Logger,
	useCase usecase.CourseStatusUseCaseInterface,
) *CourseStatusHandlers {
	return &CourseStatusHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *CourseStatusHandlers) GetCourseStatuses(c *gin.Context)      {}
func (h *CourseStatusHandlers) GetCourseStatusById(c *gin.Context)    {}
func (h *CourseStatusHandlers) AddCourseStatus(c *gin.Context)        {}
func (h *CourseStatusHandlers) UpdateCourseStatusById(c *gin.Context) {}
func (h *CourseStatusHandlers) DeleteCourseStatusById(c *gin.Context) {}
