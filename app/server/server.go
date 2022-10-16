package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    cfg.Host,
		Handler: router(handlers),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Info("shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Warn("Server Sutdown", err)
		return
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
