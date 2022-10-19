package app

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/internal/config"
	"github.com/lozovoya/GolangUnitedSchool/internal/logger"
	"github.com/lozovoya/GolangUnitedSchool/internal/repository/postgres"
	"github.com/lozovoya/GolangUnitedSchool/internal/server"
	"github.com/lozovoya/GolangUnitedSchool/internal/server/handlers"
	"github.com/lozovoya/GolangUnitedSchool/internal/usecases"
)

func Run() error {
	cfg := config.Load()

	loger := logger.NewLogger(
		cfg.Logger.Level,
		cfg.Logger.Encoding,
	)
	loger.With("service_name", cfg.ServiceName)

	repo, err := postgres.NewPostgreSQLRepository(
		context.Background(),
		cfg.PgDsn,
	)
	if err != nil {
		loger.Error(err)
		// return err
	}

	cases := usecases.NewUseCases(repo)
	handlers := handlers.NewHandlers(cases, loger)

	srv := server.NewServer(
		cfg, loger, handlers,
	)

	if err := srv.Run(context.Background()); err != nil {
		loger.Error(err)
		return err
	}

	return nil
}
