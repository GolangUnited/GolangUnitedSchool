package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CourseHandlers struct {
	lg *zap.SugaredLogger
}

func NewCourseHandler(lg *zap.SugaredLogger) *CourseHandlers {
	return &CourseHandlers{
		lg: lg,
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
