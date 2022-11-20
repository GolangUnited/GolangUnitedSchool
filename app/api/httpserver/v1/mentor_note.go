package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type MentorNoteHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseUsecaseInterface
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

// GetMentorNotes
// @Summary get all mentor notes
// @Description получить все заметки всех менторов про всех студентов
// @ID get-all-mentor-notes
// @Tags mentorNotes
// @Produce json
// @Success 200 {object} model.MentorNotesListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/notes [get]
func (h *MentorNoteHandlers) GetMentorNotes(c *gin.Context) {}

// GetMentorNotesByMentorId
// @Summary get all notes from concrete mentor
// @Description получить все заметки, оставленные конкретным ментором
// @ID get-all-notes-from-mentor
// @Tags mentorNotes
// @Produce json
// @Param id path string true "mentor_id"
// @Success 200 {object} model.MentorNotesListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/{mentor_id}/notes [get]
func (h *MentorNoteHandlers) GetMentorNotesByMentorId(c *gin.Context) {}

// GetMentorNoteByMentorNoteId
// @Summary get concrete note of mentor
// @Description получить определенную заметку, оставленную ментором
// @ID get-note-from-mentor
// @Tags mentorNotes
// @Produce json
// @Param id path string true "mentor_id"
// @Param id path string true "mentor_note_id"
// @Success 200 {object} model.MentorNotesListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/{mentor_id}/notes/{mentor_note_id} [get]
func (h *MentorNoteHandlers) GetMentorNoteByMentorNoteId(c *gin.Context) {}

// AddMentorNote
// @Summary add new mentor note
// @Description добавить менторскую заметку для студента
// @ID add-mentor-note
// @Tags mentorNotes
// @Produce json
// @Consume json
// @Param mentor_note body model.NewMentorNoteDto true "new_mentor_note"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/notes/ [post]
func (h *MentorNoteHandlers) AddMentorNote(c *gin.Context) {}

// UpdateMentorNoteByMentorNoteId
// @Summary update mentor note
// @Description изменить менторскую заметку
// @ID update-mentor-note
// @Tags mentorNotes
// @Produce json
// @Consume json
// @Param mentor_note body model.MentorNote true "update_mentor_note"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/notes/ [put]
func (h *MentorNoteHandlers) UpdateMentorNoteByMentorNoteId(c *gin.Context) {}

// DeleteMentorNoteByMentorNoteId
// @Summary delete mentor note
// @Description удалить заметку ментора
// @ID delete-note-from-mentor
// @Tags mentorNotes
// @Produce json
// @Param id path string true "mentor_note_id"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/notes/{mentor_note_id} [delete]
func (h *MentorNoteHandlers) DeleteMentorNoteByMentorNoteId(c *gin.Context) {}
