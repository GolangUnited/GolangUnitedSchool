package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/cases/user_case"
	"github.com/lozovoya/GolangUnitedSchool/app/model"
	"go.uber.org/zap"
)

type Person struct {
	cases  *user_case.UserCases
	logger *zap.SugaredLogger
}

type PersonByIdQuery struct {
	ID int64 `uri:"person_id" binding:"required"`
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

func (h Person) UpdatePerson(c *gin.Context) {
	var personId PersonByIdQuery
	if err := c.ShouldBindUri(&personId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if err := person.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if err := h.cases.UpdatePerson(c, personId.ID, &person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Success")
}

func NewPerson(
	c *user_case.UserCases,
	logger *zap.SugaredLogger) *Person {
	return &Person{
		cases:  c,
		logger: logger,
	}
}
