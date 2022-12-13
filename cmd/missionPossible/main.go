package main

import (
	"github.com/lozovoya/GolangUnitedSchool/app/api/httpserver"
	"github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/grpcserv"
	"github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/grpcserv/api"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
	"github.com/lozovoya/GolangUnitedSchool/app/config"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository/postgres"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title GolangUnitedSchool
// @version 1.0
// @description ### _API for student info collection service._

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {

	executeGrpc()
	//cfg, err := config.Load()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//if err := execute(cfg); err != nil {
	//	log.Println(err)
	//	os.Exit(1)
	//}
}

func executeGrpc() {
	s := grpc.NewServer()
	serve := &grpcserv.StudentServer{}
	__.RegisterStudentServiceServer(s, serve)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed")
	}

	if err := s.Serve(l); err != nil {
		log.Fatal("kek")
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

	// get new postgres pool of connections
	dbCtx := context.Background()
	dbPool, err := postgres.NewDbPool(dbCtx, cfg.PgDsn)
	if err != nil {
		lg.Error(err)
		return err
	}

	// init all layers
	repo := postgres.NewPostgresRepository(lg, dbPool)
	usecases := usecase.InitUsecases(lg, repo)
	handlers := v1.InitHandlers(lg, *usecases)
	router := httpserver.NewRouter(handlers)

	srv := &http.Server{
		Addr:           net.JoinHostPort(cfg.Host, cfg.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// channel to graceful shutdown
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// run server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			lg.Error("server shutdown", err)
			return
		}
	}()

	// waiting for os.Signal
	<-quit

	lg.Info("shutdown server ...")
	// shutdown server
	shCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(shCtx); err != nil {
		lg.Error(err)
	}

	return nil
}
