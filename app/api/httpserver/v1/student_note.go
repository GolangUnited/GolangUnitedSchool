package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentNoteHandlers struct {
	lg      logger.Logger
	usecase usecase.CourseUsecaseInterface
}

func NewStudentNoteHandler(
	lg logger.Logger,
	usecase usecase.StudentNoteUsecaseInterface,
) *StudentNoteHandlers {
	return &StudentNoteHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *StudentNoteHandlers) GetStudentNotesByStudentId(c *gin.Context) {

}

func (h *StudentNoteHandlers) GetStudentNoteById(c *gin.Context) {

}

func (h *StudentNoteHandlers) AddStudentNote(c *gin.Context) {

}

func (h *StudentNoteHandlers) UpdateStudentNote(c *gin.Context) {

}

func (h *StudentNoteHandlers) DeleteStudentNote(c *gin.Context) {

}
