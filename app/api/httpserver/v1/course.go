package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type CourseHandlers struct {
	lg      logger.Logger
	usecase usecase.CourseUsecaseInterface
}

func NewCourseHandler(
	lg logger.Logger,
	usecase usecase.CourseUsecaseInterface,
) *CourseHandlers {
	return &CourseHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *CourseHandlers) AddCourse(c *gin.Context) {

}

func (h *CourseHandlers) CreateCourse(c *gin.Context) {

}

func (h *CourseHandlers) GetCourseByID(c *gin.Context) {

}

func (h *CourseHandlers) SearchCourse(c *gin.Context) {

}

func (h *CourseHandlers) EditCourse(c *gin.Context) {

}

func (h *CourseHandlers) EditCourseByID(c *gin.Context) {

}

func (h *CourseHandlers) DeleteCourse(c *gin.Context) {

}

func (h *CourseHandlers) DeleteCoursebyID(c *gin.Context) {

}
