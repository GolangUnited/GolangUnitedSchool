package model

import "testing"

func TestCourseStatus_ValidateCourseStatus(t *testing.T) {
	type fields struct {
		CourseStatusId int64
		Title          string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid data",
			fields: fields{
				CourseStatusId: 1,
				Title:          "qwerty",
			},
			wantErr: false,
		},
		{
			name: "short title",
			fields: fields{
				CourseStatusId: 1,
				Title:          "w",
			},
			wantErr: true,
		},
		{
			name: "long title",
			fields: fields{
				CourseStatusId: 1,
				Title:          "qwertqwertqwertqwertqwert",
			},
			wantErr: true,
		},
		{
			name: "non alphaunicode title",
			fields: fields{
				CourseStatusId: 1,
				Title:          "qw:::",
			},
			wantErr: true,
		},
		{
			name: "empty courseStatusId",
			fields: fields{

				Title: "qwqrt",
			},
			wantErr: true,
		},
		{
			name: "empty title",
			fields: fields{
				CourseStatusId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &CourseStatus{
				CourseStatusId: tt.fields.CourseStatusId,
				Title:          tt.fields.Title,
			}
			if err := n.ValidateCourseStatus(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCourseStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

/*можно сделать проверку интов, но я пока не буду*/
