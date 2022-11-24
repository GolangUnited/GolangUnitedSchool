package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentNoteHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseUsecaseInterface
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

// GetStudentNotesByStudentId
// @Summary get all student notes by student id
// @Description возвращает все заметки по студенту с указанным id
// @ID get-student-notes
// @Tags studentNotes
// @Produce json
// @Param id path string true "student_id"
// @Success 200 {object} model.StudentNotesListDto
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/notes [get]
func (h *StudentNoteHandlers) GetStudentNotesByStudentId(c *gin.Context) {}

// GetStudentNoteByStudentNoteId
// @Summary get concrete student's note from concrete student
// @Description возвращает заметку с указанным id для студента с указанным id
// @ID get-student-note-by-student-note-id
// @Tags studentNotes
// @Produce json
// @Param id path string true "student_id"
// @Param id path string true "student_note_id"
// @Success 200 {object} model.StudentNote
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/notes/{student_note_id} [get]
func (h *StudentNoteHandlers) GetStudentNoteByStudentNoteId() {}

// AddStudentNote
// @Summary add new note to student
// @Description добавляет заметку студента
// @ID add-new-student-note
// @Tags studentNotes
// @Produce json
// @Consume json
// @Param student_note body model.NewStudentNoteDto true "student_note"
// @Param id path string true "student_id"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/notes [post]
func (h *StudentNoteHandlers) AddStudentNote(c *gin.Context) {}

// PutStudentNoteByStudentNoteId
// @Summary update student note
// @Description изменяет заметку студента с указанным id студента
// @ID update-student-note
// @Tags studentNotes
// @Param id path string true "student_id"
// @Param id path string true "student_note_id"
// @Param student_note body model.StudentNote true "student_note"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/{student_id}/notes/{student_note_id} [put]
func (h *StudentNoteHandlers) PutStudentNoteByStudentNoteId(c *gin.Context) {}

// UpdateStudentNoteByStudentNoteId
// @Summary update student note by Id
// @ID update-student-note-by-id
// @Tags studentNotes
// @Param id path string true "student_note_id"
// @Param student_note body model.UpdateStudentNoteDto true "student_note"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/notes/{student_id} [post]
func (h *StudentGroupHandlers) UpdateStudentNoteByStudentNoteId(c *gin.Context) {}

// DeleteStudentNote
// @Summary delete student note
// @Description удаляет заметку студента с указанным id студента
// @ID delete-student-note
// @Tags studentNotes
// @Param id path string true "student_note_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/notes/{student_note_id} [delete]
func (h *StudentNoteHandlers) DeleteStudentNote(c *gin.Context) {}
