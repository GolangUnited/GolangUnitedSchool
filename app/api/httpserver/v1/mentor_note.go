package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type MentorNoteHandlers struct {
	lg      logger.Logger
	usecase usecase.CourseUsecaseInterface
}

func NewMentorNoteHandler(
	lg logger.Logger,
	usecase usecase.MentorUsecaseInterface,
) *MentorNoteHandlers {
	return &MentorNoteHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *MentorNoteHandlers) GetMentorNotes(c *gin.Context) {

}
func (h *MentorNoteHandlers) GetMentorNoteById(c *gin.Context) {

}
func (h *MentorNoteHandlers) AddNewMentorNote(c *gin.Context) {

}
func (h *MentorNoteHandlers) UpdateMentorNotebyId(c *gin.Context) {

}

func (h *MentorNoteHandlers) DeleteMentorNoteById(c *gin.Context) {

}
