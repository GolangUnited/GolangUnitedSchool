package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lozovoya/GolangUnitedSchool/app/config"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
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

	srv := &http.Server{
		Addr: cfg.Host,
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
