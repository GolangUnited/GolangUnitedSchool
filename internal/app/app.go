package app

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/internal/config"
	"github.com/lozovoya/GolangUnitedSchool/internal/logger"
	"github.com/lozovoya/GolangUnitedSchool/internal/repository"
	"github.com/lozovoya/GolangUnitedSchool/internal/repository/postgres"
	"github.com/lozovoya/GolangUnitedSchool/internal/server"
	"github.com/lozovoya/GolangUnitedSchool/internal/server/handlers"
	"github.com/lozovoya/GolangUnitedSchool/internal/usecases"
	"go.uber.org/zap"
)

type App struct {
	cfg      *config.Config
	log      *zap.SugaredLogger
	repo     repository.Repository
	cases    *usecases.Cases
	handlers *handlers.Handlers
}

func Run() {
	var app App
	app.cfg = config.Load()

	app.log = logger.NewLogger(
		app.cfg.Logger.Level,
		app.cfg.Logger.Encoding,
	)
	app.log.With("service_name", app.cfg.ServiceName)

	var err error
	app.repo, err = postgres.NewPostgreSQLRepository(
		context.Background(),
		app.cfg.PgDsn,
	)
	if err != nil {
		app.log.Error(err)
		return
	}

	app.cases = usecases.NewUseCases(app.repo)
	app.handlers = handlers.NewHandlers(app.cases, app.log)

	srv := server.NewServer(
		app.cfg, app.log, app.handlers,
	)

	if err := srv.Run(context.Background()); err != nil {
		app.log.Error(err)
		return
	}
}
