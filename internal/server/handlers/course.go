package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/internal/usecases"
	"go.uber.org/zap"
)

type Course struct {
	cases  usecases.Course
	logger *zap.SugaredLogger
}

func NewCourses(c usecases.Course,
	logger *zap.SugaredLogger) *Course {
	return &Course{
		cases:  c,
		logger: logger,
	}
}

func (s *Course) CreateCourse(c *gin.Context) {

}

func (s *Course) GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
}

func (s *Course) GetCoursesByFilter(c *gin.Context) {

}

func (s *Course) UpdateCourse(c *gin.Context) {

}

func (s *Course) DeleteCourseByID(c *gin.Context) {

}
