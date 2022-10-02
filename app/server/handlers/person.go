package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
	"go.uber.org/zap"
)

type PersonHandler struct {
	repository repository.PostgreSQLRepository
	logger     *zap.SugaredLogger
}

type PersonByIdQuery struct {
	ID int64 `uri:"person_id" binding:"required"`
}

func (h PersonHandler) GetPersonById(ctx *gin.Context) {
	var personQuery PersonByIdQuery
	if err := ctx.ShouldBindUri(&personQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	person, err := h.repository.GetPersonById(ctx, personQuery.ID)
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	ctx.JSON(http.StatusOK, person)
}

func NewPersonHandler(r repository.PostgreSQLRepository,
	logger *zap.SugaredLogger) *PersonHandler {
	return &PersonHandler{
		repository: r,
		logger:     logger,
	}
}
