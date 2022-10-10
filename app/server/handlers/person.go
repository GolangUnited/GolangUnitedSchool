package handlers

import (
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
	firstName  string    `form:"first_name" binding:"required"`
	lastName   string    `form:"last_name" binding:"required"`
	patronymic string    `form:"patronymic"`
	login      string    `form:"login" binding:"required"`
	roleId     int       `form:"role_id" binding:"required"`
	passwd     string    `form:"passwd" binding:"required"`
	updatedAt  time.Time `form:"updated_at"`
	deleted    bool      `form:"deleted"`
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

func NewPerson(
	c *user_case.UserCases,
	logger *zap.SugaredLogger) *Person {
	return &Person{
		cases:  c,
		logger: logger,
	}
}
