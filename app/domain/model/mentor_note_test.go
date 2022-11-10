package model

import (
	"testing"
	"time"
)

func TestMentorNote_ValidateMentorNote(t *testing.T) {
	type fields struct {
		MentorNoteId int64
		StudentId    int64
		MentorId     int64
		Note         string
		CreatedAt    time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid data",
			fields: fields{
				MentorNoteId: 1,
				StudentId:    1,
				MentorId:     1,
				Note:         "hello world!",
				CreatedAt:    time.Now(),
			},
			wantErr: false,
		},
		{
			name: "zero values",
			fields: fields{
				MentorNoteId: 0,
				StudentId:    0,
				MentorId:     0,
				Note:         "hello world!",
				CreatedAt:    time.Now(),
			},
			wantErr: true,
		},
		{
			name: "empty field",
			fields: fields{
				//MentorNoteId: 1,
				StudentId: 1,
				MentorId:  1,
				Note:      "hello world!",
				CreatedAt: time.Now(),
			},
			wantErr: true,
		},
		{
			name: "empty field2",
			fields: fields{
				MentorNoteId: 1,
				//StudentId: 1,
				//MentorId:  1,
				Note:      "hello world!",
				CreatedAt: time.Now(),
			},
			wantErr: true,
		},
		{
			name: "short note",
			fields: fields{
				MentorNoteId: 1,
				StudentId:    1,
				MentorId:     1,
				Note:         "!",
				CreatedAt:    time.Now(),
			},
			wantErr: true,
		},
		{
			name: "long note",
			fields: fields{
				MentorNoteId: 1,
				StudentId:    1,
				MentorId:     1,
				Note: "qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop" +
					"qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop" +
					"qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop" +
					"qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop" +
					"qwertyuiopqwertyuiopqwertyuiop",
				CreatedAt: time.Now(),
			},
			wantErr: true,
		},
		{
			name: "omit epty",
			fields: fields{
				MentorNoteId: 1,
				StudentId:    1,
				MentorId:     1,
				Note:         "!1111",
				//CreatedAt:    time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &MentorNote{
				MentorNoteId: tt.fields.MentorNoteId,
				StudentId:    tt.fields.StudentId,
				MentorId:     tt.fields.MentorId,
				Note:         tt.fields.Note,
				CreatedAt:    tt.fields.CreatedAt,
			}
			if err := n.ValidateMentorNote(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateMentorNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
