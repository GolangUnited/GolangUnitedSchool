package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/usecases"

	"go.uber.org/zap"
)

type Auth struct {
	cases  usecases.Auth
	logger *zap.SugaredLogger
}

func NewAuth(cases usecases.Auth,
	logger *zap.SugaredLogger) *Auth {
	return &Auth{
		cases:  cases,
		logger: logger,
	}
}

func (a *Auth) Login(ctx *gin.Context) {

}

func (a *Auth) Logout(ctx *gin.Context) {

}
