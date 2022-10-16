package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/config"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository/postgres"
	"github.com/lozovoya/GolangUnitedSchool/app/server/handlers"
	"github.com/lozovoya/GolangUnitedSchool/app/usecases"
)

func Run(cfg *config.Config) {
	log := logger.NewLogger(
		cfg.Logger.Level,
		cfg.Logger.Encoding,
	)

	ctx := context.Background()
	repo, err := postgres.NewPostgreSQLRepository(ctx, cfg.PGConnectionString)
	if err != nil {
		log.Error(err)
		return
	}
	
	usecases := usecases.NewUseCases(repo)
	handlers := handlers.NewHandlers(usecases, log)

	srv := &http.Server{
		Addr:    cfg.Host,
		Handler: router(handlers),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	// gracfull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Info("shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Warn("Server Sutdown", err)
		return
	}
}

func router(h *handlers.HandlerSt) *gin.Engine {
	// *gin.Engine with recovery and loging middlewares
	r := gin.Default()

	// TODO: get cors config from app configurations
	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Origin", "Content-Length", "Content-Type"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			MaxAge: 12 * time.Hour,
		},
	))

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
