package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/cases/user_case"
	"github.com/lozovoya/GolangUnitedSchool/app/domain"
	"go.uber.org/zap"
)

type Person struct {
	cases  *user_case.UserCases
	logger *zap.SugaredLogger
}

type PersonByIdQuery struct {
	ID int64 `uri:"person_id" binding:"required"`
}

func NewPerson(
	c *user_case.UserCases,
	logger *zap.SugaredLogger) *Person {
	return &Person{
		cases:  c,
		logger: logger,
	}
}

func (h Person) GetPersonById(ctx *gin.Context) {
	var personQuery PersonByIdQuery
	if err := ctx.ShouldBindUri(&personQuery); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{IsSuccess: false, Message: "Incorrect input"})
		return
	}

	person, err := h.cases.GetPersonById(ctx, personQuery.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{IsSuccess: false, Message: "Server error", Data: err})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{IsSuccess: true, Data: person})
}

func (h Person) DeletePerson(ctx *gin.Context) {
	var personQuery PersonByIdQuery
	if err := ctx.ShouldBindUri(&personQuery); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.Response{IsSuccess: false, Message: "Incorrect input", Data: err})
		return
	}

	if err := h.cases.DeletePerson(ctx, personQuery.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, domain.Response{IsSuccess: false, Message: "Server error", Data: err})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{IsSuccess: true, Message: "Person deleted"})
}
