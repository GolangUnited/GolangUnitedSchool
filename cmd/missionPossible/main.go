package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lozovoya/GolangUnitedSchool/app/api/httpserver"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
	"github.com/lozovoya/GolangUnitedSchool/app/config"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository/postgres"
	"golang.org/x/net/context"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println(err)
		return
	}

	if err := execute(cfg); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(cfg *config.Config) error {
	lg, err := zap.NewLogger(
		cfg.Logger.Level,
		cfg.Logger.Encoding,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	// add service name to logger
	lg.With("service name", cfg.ServiceName)

	dbCtx := context.Background()
	dbPool, err := postgres.NewDbPool(dbCtx, cfg.PgDsn)
	if err != nil {
		lg.Error(err)
		return err
	}

	// init course repository, usecase and handlers
	courseRepo := postgres.NewCourse(lg, dbPool)
	courseUseCase := usecase.NewCourse(lg, courseRepo)
	courseHandler := v1.NewCourseHandler(lg, courseUseCase)

	projectRepo := postgres.NewProject(lg, dbPool)
	projectUseCase := usecase.NewProject(lg, projectRepo)
	projectHandler := v1.NewProjectHandler(lg, projectUseCase)

	operationLogRepo := postgres.NewOperationLog(lg, dbPool)
	operationLogUseCase := usecase.NewOperationLog(lg, operationLogRepo)
	operationLogHandler := v1.NewOperationLogHandler(lg, operationLogUseCase)

	operationRepo := postgres.NewOperation(lg, dbPool)
	operationUseCase := usecase.NewOperation(lg, operationRepo)
	operationHandler := v1.NewOperationHandler(lg, operationUseCase)

	operationTypeRepo := postgres.NewOperationType(lg, dbPool)
	operationTypeUseCase := usecase.NewOperationType(lg, operationTypeRepo)
	operationTypeHandler := v1.NewOperationTypeHandler(lg, operationTypeUseCase)

	contactTypeRepo := postgres.NewContactType(lg, dbPool)
	contactTypeUseCase := usecase.NewContactType(lg, contactTypeRepo)
	contactTypeHandler := v1.NewContactTypeHandler(lg, contactTypeUseCase)

	router := httpserver.NewRouter(
		courseHandler, projectHandler, operationLogHandler,
		operationHandler, operationTypeHandler, contactTypeHandler,
	)
	srv := &http.Server{
		Addr:           cfg.Host,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			lg.Error("server shut down")
			return
		}
	}()

	quit := make(chan os.Signal, 1)

	<-quit

	shCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(shCtx); err != nil {
		lg.Error(err)
	}

	return nil
}
