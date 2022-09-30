package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
	"github.com/lozovoya/GolangUnitedSchool/app/server/handlers"
)

func Run() {
	
	repository := repository.NewPostgreSQLRepository("pg_connection_string")
	personHandler := handlers.NewPersonHandler(*repository)


	r := gin.Default()

	r.GET("/person/:person_id", personHandler.GetPersonById)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
