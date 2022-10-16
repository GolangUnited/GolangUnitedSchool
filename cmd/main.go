package main

import (
	"github.com/lozovoya/GolangUnitedSchool/internal/config"
	"github.com/lozovoya/GolangUnitedSchool/internal/server"
)

func main() {
	server.Run(config.Load())
}
