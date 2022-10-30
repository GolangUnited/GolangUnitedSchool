package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentNoteHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseUseCaseInterface
}

func NewStudentNoteHandler(
	lg logger.Logger,
	useCase usecase.StudentNoteUseCaseInterface,
) *StudentNoteHandlers {
	return &StudentNoteHandlers{
		lg:      lg,
		useCase: useCase,
	}
}
func (h *StudentNoteHandlers) GetStudentNotes(c *gin.Context)                {}
func (h *StudentNoteHandlers) GetStudentNotesByStudentId(c *gin.Context)     {}
func (h *StudentNoteHandlers) GetStudentNoteByStudentNoteId(c *gin.Context)  {}
func (h *StudentNoteHandlers) AddStudentNote(c *gin.Context)                 {}
func (h *StudentNoteHandlers) EditStudentNoteByStudentNoteId(c *gin.Context) {}
func (h *StudentNoteHandlers) DeleteStudentNote(c *gin.Context)              {}
