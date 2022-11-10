package model

type StudentNoteType struct {
	StudentNoteTypeId int64  `json:"student_note_type_id" validate:"required,gte=0"`
	Title             string `json:"title" validate:"required,gte=2,lte=10"`
}
