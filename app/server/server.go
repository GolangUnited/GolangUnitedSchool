package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/cases/user_case"
	"github.com/lozovoya/GolangUnitedSchool/app/config"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
	"github.com/lozovoya/GolangUnitedSchool/app/repository/postgres"
	"github.com/lozovoya/GolangUnitedSchool/app/server/handlers"
)

var repo repository.Repository

func Run(cfg *config.Config) {
	logger := NewLogger(cfg.Logger.Level,
		cfg.Logger.Encoding)
	ctx := context.Background()
	repo = postgres.NewPostgreSQLRepository(ctx, cfg.PGConnectionString)
	userCases := user_case.NewUserCases(repo)
	personHandler := handlers.NewPerson(userCases, logger)

	r := gin.Default()

	r.GET("/person/:person_id", personHandler.GetPersonById)
	r.POST("/person", personHandler.PostNewPerson)
	if err := r.Run(); err != nil {
		panic(err)
	}

}
