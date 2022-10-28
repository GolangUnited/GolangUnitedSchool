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

	// init course repository, useCase and handlers
	courseRepo := postgres.NewCourse(lg, dbPool)
	courseUseCase := usecase.NewCourse(lg, courseRepo)
	courseHandler := v1.NewCourseHandler(lg, courseUseCase)
	// init person rep, useCase and handlers
	personRepo := postgres.NewPerson(lg, dbPool)
	personUseCase := usecase.NewPerson(lg, personRepo)
	personHandler := v1.NewPersonHandler(lg, personUseCase)
	// init student rep, useCase and handlers
	studentRepo := postgres.NewStudent(lg, dbPool)
	studentUseCase := usecase.NewStudent(lg, studentRepo)
	studentHandler := v1.NewStudentHandler(lg, studentUseCase)
	// init mentor rep, useCase and handlers
	mentorRepo := postgres.NewMentor(lg, dbPool)
	mentorUseCase := usecase.NewMentor(lg, mentorRepo)
	mentorHandler := v1.NewMentorHandler(lg, mentorUseCase)
	// init
	mentorNoteRepo := postgres.NewMentorNote(lg, dbPool)
	mentorNoteUseCase := usecase.NewMentorNote(lg, mentorNoteRepo)
	mentorNoteHandler := v1.NewMentorNoteHandler(lg, mentorNoteUseCase)
	// init
	studentNoteRepo := postgres.NewStudentNote(lg, dbPool)
	studentNoteUseCase := usecase.NewStudentNote(lg, studentNoteRepo)
	studentNoteHandler := v1.NewStudentNoteHandler(lg, studentNoteUseCase)
	// init
	studentNoteTypeRepo := postgres.NewStudentNoteType(lg, dbPool)
	studentNoteTypeUseCase := usecase.NewStudentNoteType(lg, studentNoteTypeRepo)
	studentNoteTypeHandler := v1.NewStudentNoteTypeHandler(lg, studentNoteTypeUseCase)

	router := httpserver.NewRouter(
		courseHandler,
		personHandler,
		studentHandler,
		mentorHandler,
		mentorNoteHandler,
		studentNoteHandler,
		studentNoteTypeHandler,
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
