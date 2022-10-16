package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/config"
	"github.com/lozovoya/GolangUnitedSchool/app/repository/postgres"
	"github.com/lozovoya/GolangUnitedSchool/app/server/handlers"
	"github.com/lozovoya/GolangUnitedSchool/app/usecases"
)

func Run(cfg *config.Config) {
	logger := NewLogger(cfg.Logger.Level,
		cfg.Logger.Encoding)
	ctx := context.Background()
	repo := postgres.NewPostgreSQLRepository(ctx, cfg.PGConnectionString)
	usecases := usecases.NewUseCases(repo)
	handlers := handlers.NewHandlers(usecases, logger)

	r := router(handlers)

	if err := r.Run(); err != nil {
		panic(err)
	}
}

func router(h *handlers.HandlerSt) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", h.Auth.Login)
		auth.GET("/logout", h.Auth.Logout)
	}

	person := r.Group("/person")
	{
		person.GET("/:person_id", h.Person.GetPersonById)
	}

	course := r.Group("/course")
	{
		course.GET("")
		course.GET(":id")
		course.POST("/")
		course.PUT("")
		course.DELETE("/:id")
	}

	return r
}
