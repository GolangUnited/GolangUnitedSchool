package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type CourseHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseUseCaseInterface
}

func NewCourseHandler(
	lg logger.Logger,
	useCase usecase.CourseUseCaseInterface,
) *CourseHandlers {
	return &CourseHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *CourseHandlers) AddCourse(c *gin.Context) {

}

func (h *CourseHandlers) CreateCourse(c *gin.Context) {

}

func (h *CourseHandlers) GetCourseById(c *gin.Context) {

}

func (h *CourseHandlers) SearchCourse(c *gin.Context) {

}

func (h *CourseHandlers) EditCourse(c *gin.Context) {

}

func (h *CourseHandlers) EditCourseById(c *gin.Context) {

}

func (h *CourseHandlers) DeleteCourse(c *gin.Context) {

}

func (h *CourseHandlers) DeleteCourseById(c *gin.Context) {

}
