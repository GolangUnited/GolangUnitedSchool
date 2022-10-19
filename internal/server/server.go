package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/internal/config"
	"github.com/lozovoya/GolangUnitedSchool/internal/server/handlers"

	"go.uber.org/zap"
)

// @title Golang United School
// @version 1.0
// @description This is a sample server.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
type ServerSt struct {
	log *zap.SugaredLogger
	h   *handlers.Handlers
	srv *http.Server
}

func NewServer(
	cfg *config.Config,
	log *zap.SugaredLogger,
	h *handlers.Handlers,
) *ServerSt {
	return &ServerSt{
		log: log,
		h:   h,
		srv: &http.Server{
			Addr:           fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:        router(h),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *ServerSt) Run(ctx context.Context) error {

	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			s.log.Error(err)
		}
	}()

	// gracfull shutdown
	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
	case <-ctx.Done():
	}

	s.log.Info("shutdown Server ...")
	ctxSh, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := s.srv.Shutdown(ctxSh); err != nil {
		s.log.Warn("Server Sutdown", err)
		return err
	}

	return nil
}

func router(h *handlers.Handlers) *gin.Engine {
	// *gin.Engine with recovery and loging middlewares
	r := gin.Default()

	// TODO: get CORS configs from app configurations
	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
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
		course.GET("/:id", h.Course.GetCourseByID)
		course.POST("")
		course.POST("/:id")
		course.PUT("/:id")
		course.DELETE("/:id")
	}

	return r
}
