package model

import "testing"

func TestStudentNoteType_ValidateStudentNoteType(t *testing.T) {
	type fields struct {
		StudentNoteTypeId int64
		Title             string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid data",
			fields: fields{
				StudentNoteTypeId: 1,
				Title:             "work exp",
			},
			wantErr: false,
		},
		{
			name:   "empty fields",
			fields: fields{
				//StudentNoteTypeId: 1,
				//Title:             "work exp",
			},
			wantErr: true,
		},
		{
			name: "non valid value",
			fields: fields{
				StudentNoteTypeId: 0,
				Title:             "work exp",
			},
			wantErr: true,
		},
		{
			name: "short title",
			fields: fields{
				StudentNoteTypeId: 1,
				Title:             "w",
			},
			wantErr: true,
		},
		{
			name: "long title",
			fields: fields{
				StudentNoteTypeId: 1,
				Title:             "qwertyuiopp",
			},
			wantErr: true,
		},
		{
			name: "empty title",
			fields: fields{
				StudentNoteTypeId: 1,
				Title:             "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &StudentNoteType{
				StudentNoteTypeId: tt.fields.StudentNoteTypeId,
				Title:             tt.fields.Title,
			}
			if err := n.ValidateStudentNoteType(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateStudentNoteType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
