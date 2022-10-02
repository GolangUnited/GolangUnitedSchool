package main

import (
	"github.com/lozovoya/GolangUnitedSchool/app/config"
	"github.com/lozovoya/GolangUnitedSchool/app/server"
)

func main() {
	server.Run(config.Load())
}
