package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/config"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
	"github.com/lozovoya/GolangUnitedSchool/app/server/handlers"
)

func Run(cfg *config.Model) {
	logger := NewLogger(cfg.Logger.Level,
		cfg.Logger.Encoding)
	ctx := context.Background()
	repository := repository.NewPostgreSQLRepository(ctx, cfg.PGConnectionString, logger)
	personHandler := handlers.NewPersonHandler(*repository, logger)

	r := gin.Default()

	r.GET("/person/:person_id", personHandler.GetPersonById)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
