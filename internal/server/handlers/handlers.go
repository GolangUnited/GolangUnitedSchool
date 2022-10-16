package handlers

import (
	"github.com/lozovoya/GolangUnitedSchool/internal/usecases"
	"go.uber.org/zap"
)

type Handlers struct {
	Auth   *Auth
	Course *Course
	Person *Person
}

func NewHandlers(c *usecases.Cases,
	logger *zap.SugaredLogger) *Handlers {
	return &Handlers{
		Auth:   NewAuth(c.Auth, logger),
		Course: NewCourses(c.Course, logger),
		Person: NewPerson(c.User, logger),
	}
}
