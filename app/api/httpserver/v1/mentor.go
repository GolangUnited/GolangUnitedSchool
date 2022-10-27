package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type MentorHandlers struct {
	lg      logger.Logger
	usecase usecase.MentorUsecaseInterface
}

func NewMentorHandler(
	lg logger.Logger,
	usecase usecase.MentorUsecaseInterface,
) *MentorHandlers {
	return &MentorHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *MentorHandlers) GetAllMentors(c *gin.Context) {
}
func (h *MentorHandlers) GetMentorById(c *gin.Context) {
}
func (h *MentorHandlers) SearchMentor(c *gin.Context) {
}
func (h *MentorHandlers) AddMentorByName(c *gin.Context) {
}
func (h *MentorHandlers) RemoveMentorByName(c *gin.Context) {
}
func (h *MentorHandlers) RemoveMentorById(c *gin.Context) {
}
