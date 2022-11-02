package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type MentorNoteHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseUseCaseInterface
}

func NewMentorNoteHandler(
	lg logger.Logger,
	useCase usecase.MentorUseCaseInterface,
) *MentorNoteHandlers {
	return &MentorNoteHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *MentorNoteHandlers) GetMentorNotes(c *gin.Context)                 {}
func (h *MentorNoteHandlers) GetMentorNotesByMentorId(c *gin.Context)       {}
func (h *MentorNoteHandlers) GetMentorNoteByMentorNoteId(c *gin.Context)    {}
func (h *MentorNoteHandlers) AddMentorNote(c *gin.Context)                  {}
func (h *MentorNoteHandlers) EditMentorNoteByMentorNoteId(c *gin.Context)   {}
func (h *MentorNoteHandlers) DeleteMentorNoteByMentorNoteId(c *gin.Context) {}
