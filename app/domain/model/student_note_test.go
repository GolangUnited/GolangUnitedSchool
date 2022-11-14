package model

import (
	"testing"
	"time"
)

func TestStudentNote_ValidateStudentNote(t *testing.T) {
	type fields struct {
		StudentNoteId     int64
		StudentId         int64
		StudentNoteTypeId int64
		Note              string
		CreatedAt         time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid data",
			fields: fields{
				StudentNoteId:     1,
				StudentId:         1,
				StudentNoteTypeId: 1,
				Note:              "qwerty",
				CreatedAt:         time.Now(),
			},
			wantErr: false,
		},
		{
			name: "short note",
			fields: fields{
				StudentNoteId:     1,
				StudentId:         1,
				StudentNoteTypeId: 1,
				Note:              "q",
				CreatedAt:         time.Now(),
			},
			wantErr: true,
		},
		{
			name: "long note",
			fields: fields{
				StudentNoteId:     1,
				StudentId:         1,
				StudentNoteTypeId: 1,
				Note: "qwertyuiopqwertyuiopqwertyuiopqwertyuiop" +
					"qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop" +
					"qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop" +
					"qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop" +
					"qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop" +
					"qwertyuiopqwertyuiop",
				CreatedAt: time.Now(),
			},
			wantErr: true,
		},
		{
			name: "omit empty",
			fields: fields{
				StudentNoteId:     1,
				StudentId:         1,
				StudentNoteTypeId: 1,
				Note:              "qkl",
				//CreatedAt:         time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &StudentNote{
				StudentNoteId:     tt.fields.StudentNoteId,
				StudentId:         tt.fields.StudentId,
				StudentNoteTypeId: tt.fields.StudentNoteTypeId,
				Note:              tt.fields.Note,
				CreatedAt:         tt.fields.CreatedAt,
			}
			if err := n.ValidateStudentNote(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateStudentNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
