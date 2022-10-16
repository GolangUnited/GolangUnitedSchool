package handlers

import (
	"github.com/lozovoya/GolangUnitedSchool/app/usecases"
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
