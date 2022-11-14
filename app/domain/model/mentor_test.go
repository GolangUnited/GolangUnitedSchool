package model

import (
	"encoding/json"
	"testing"
)

func TestMentor_ValidateMentor(t *testing.T) {
	type fields struct {
		MentorId int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:   "empty field mentor id",
			fields: fields{
				//MentorId: 0,
			},
			wantErr: true,
		},
		{
			name: "valid data",
			fields: fields{
				MentorId: 1,
			},
			wantErr: false,
		},
		{
			name: "zero mentor",
			fields: fields{
				MentorId: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Mentor{
				MentorId: tt.fields.MentorId,
			}
			if err := n.ValidateMentor(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateMentor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMentor_ValidateNonNumeric(t *testing.T) {
	var testStruct Mentor
	nonNumericJson := []byte(`{
		"mentor_id" : "2"
}`)
	_ = json.Unmarshal(nonNumericJson, &testStruct)

	if err := testStruct.ValidateMentor(); err == nil {
		t.Error("not passed", err)
	}

}
