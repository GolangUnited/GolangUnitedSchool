package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/cases/user_case"
	"go.uber.org/zap"
)

type Person struct {
	cases  *user_case.UserCases
	logger *zap.SugaredLogger
}

type PersonByIdQuery struct {
	ID int64 `uri:"person_id" binding:"required"`
}

type NewPersonQuery struct {
	FirstName  string    `json:"first_name"  binding:"required"`
	LastName   string    `json:"last_name" binding:"required"`
	Patronymic string    `json:"patronymic,"`
	Login      string    `json:"login" binding:"required"`
	RoleId     int       `json:"role_id" binding:"required"`
	Passwd     string    `json:"passwd" binding:"required"`
	UpdatedAt  time.Time `json:"updated_at"`
	Deleted    bool      `json:"deleted" binding:"boolean"`
}

func (h Person) GetPersonById(ctx *gin.Context) {
	var personQuery PersonByIdQuery
	if err := ctx.ShouldBindUri(&personQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	person, err := h.cases.GetPersonById(ctx, personQuery.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	ctx.JSON(http.StatusOK, person)
}

func (h Person) PostNewPerson(ctx *gin.Context) {
	var newPerson NewPersonQuery
	if err := ctx.ShouldBindJSON(&newPerson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})

	}
	person, err := h.cases.PostNewPerson(ctx, newPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	ctx.JSON(http.StatusOK, person)
	fmt.Println(newPerson)

}

func NewPerson(
	c *user_case.UserCases,
	logger *zap.SugaredLogger) *Person {
	return &Person{
		cases:  c,
		logger: logger,
	}
}
