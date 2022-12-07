package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentNoteTypeHandlers struct {
	lg      logger.Logger
	useCase usecase.StudentNoteTypeUseCaseInterface
}

func NewStudentNoteTypeHandler(
	lg logger.Logger,
	useCase usecase.StudentNoteTypeUseCaseInterface,
) *StudentNoteTypeHandlers {
	return &StudentNoteTypeHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// GetStudentNoteTypes
// @Summary get list of student note types
// @Description возвращает все типы студенческих заметок
// @ID get-all-student-note-types
// @Tags studentNoteTypes
// @Produce json
// @Success 200 {object} []model.StudentNoteType
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/notes/types [get]
func (h *StudentNoteTypeHandlers) GetStudentNoteTypes(c *gin.Context) {}

// GetStudentNoteTypeById
// @Summary get student note type by id
// @Description возвращает тип студенческой заметки с указанным id
// @ID get-student-student-note-type-by-id
// @Tags studentNoteTypes
// @Produce json
// @Param id path string true "Student_note_type_id"
// @Success 200 {object} model.StudentNoteType
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/notes/types/{student_note_type_id} [get]
func (h *StudentNoteTypeHandlers) GetStudentNoteTypeById(c *gin.Context) {}

// AddStudentNoteType
// @Summary add new student type note
// @Description добавляет новый тип студенческой заметки
// @ID create-note-types
// @Tags studentNoteTypes
// @Produce json
// @Consume json
// @Param new_student_note_type body model.NewStudentNoteType true "student note type"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/notes/types [post]
func (h *StudentNoteTypeHandlers) AddStudentNoteType(c *gin.Context) {}

// PutStudentNoteTypeById
// @Summary update student note type
// @Description изменяет тип студенческой заметки с указанным id
// @ID update-student-note-type
// @Tags studentNoteTypes
// @Param id path string true "student_note_type_id"
// @Param student_note_type body model.UpdateStudentNoteType true "student note type"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/notes/types/{student_note_type_id} [put]
func (h *StudentNoteTypeHandlers) PutStudentNoteTypeById(c *gin.Context) {}

// UpdateStudentNoteTypeById
// @Summary update student note type bu id
// @ID update-student-note-type-by-id
// @Tags studentNoteTypes
// @Param id path string true "student_note_type_id"
// @Param student_note_type body model.UpdateStudentNoteType true "student_note_type"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/notes/types/{id} [post]
func (h *StudentNoteTypeHandlers) UpdateStudentNoteTypeById(c *gin.Context) {}

// DeleteStudentNoteTypeById
// @Summary delete student note type by id
// @Description удаляет тип заметки студента по id
// @ID delete-student-note-type-by-id
// @Tags studentNoteTypes
// @Param id path string true "student_note_type_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /students/notes/types/{student_note_type_id} [delete]
func (h *StudentNoteTypeHandlers) DeleteStudentNoteTypeById(c *gin.Context) {}
