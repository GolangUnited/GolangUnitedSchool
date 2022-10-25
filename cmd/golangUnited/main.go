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

	lg, err := zap.NewLogger(
		cfg.Logger.Level,
		cfg.Logger.Encoding,
	)
	if err != nil {
		log.Println(err)
		return
	}
	lg.With("service name", cfg.ServiceName)

	dbCtx := context.Background()
	dbPool, err := postgres.NewDbPool(dbCtx, cfg.PgDsn)
	if err != nil {
		lg.Error(err)
		return
	}

	courseRepo := postgres.NewCourse(lg, dbPool)
	courseUseCase := usecase.NewCourse(lg, courseRepo)
	courseHandler := v1.NewCourseHandler(lg, courseUseCase)
	router := httpserver.NewRouter()
	srv := &http.Server{
		Addr:    cfg.Host,
		Handler: router,
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

}

func execute() {

}
