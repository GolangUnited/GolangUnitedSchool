package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type InterviewHandlers struct {
	lg      logger.Logger
	usecase usecase.InterviewUsecaseInterface
}

func NewInterview(lg logger.Logger, u usecase.InterviewUsecaseInterface) *InterviewHandlers {
	return &InterviewHandlers{
		lg:      lg,
		usecase: u,
	}
}

func GetInterviewById(c *gin.Context)    {}
func GetInterviews(c *gin.Context)       {}
func AddInterview(c *gin.Context)        {}
func UpdateInterviewById(c *gin.Context) {}
func DeleteInterviewById(c *gin.Context) {}
