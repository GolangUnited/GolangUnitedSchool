package handlers

import (
	"github.com/lozovoya/GolangUnitedSchool/app/usecases"
	"go.uber.org/zap"
)

type HandlerSt struct {
	Auth   *Auth
	Course *Course
	Person *Person
}

func NewHandlers(c *usecases.Cases,
	logger *zap.SugaredLogger) *HandlerSt {
	return &HandlerSt{
		Auth:   NewAuth(c.Auth, logger),
		Course: NewCourses(c.Course, logger),
		Person: NewPerson(c.User, logger),
	}
}
