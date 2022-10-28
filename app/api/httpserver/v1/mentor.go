package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type MentorHandlers struct {
	lg      logger.Logger
	useCase usecase.MentorUseCaseInterface
}

func NewMentorHandler(
	lg logger.Logger,
	useCase usecase.MentorUseCaseInterface,
) *MentorHandlers {
	return &MentorHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *MentorHandlers) GetAllMentors(c *gin.Context) {}

func (h *MentorHandlers) GetMentorByMentorId(c *gin.Context) {}

func (h *MentorHandlers) AddMentor(c *gin.Context) {}

func (h *MentorHandlers) EditMentor(c *gin.Context) {}

func (h *MentorHandlers) RemoveMentorByMentorId(c *gin.Context) {}
