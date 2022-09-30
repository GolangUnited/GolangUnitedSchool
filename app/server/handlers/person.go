package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type PersonHandler struct {
	repository repository.PostgreSQLRepository
}

type PersonQuery struct {
	ID int
	Name string
}

func (h PersonHandler) GetPersonById(ctx *gin.Context) {
	personID := ctx.GetInt64("person_id")

	person := h.repository.GetPersonById(personID)
	ctx.JSON(http.StatusOK, person)
}

func NewPersonHandler(r repository.PostgreSQLRepository) *PersonHandler {
	return &PersonHandler{
		repository: r,
	}
}
