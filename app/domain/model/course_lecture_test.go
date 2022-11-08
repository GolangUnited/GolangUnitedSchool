package model

import (
	"encoding/json"
	"testing"
)

var testStruct CourseLecture

func TestCourseLecture_ValidateCourseLecture(t *testing.T) {
	NonNumericCourseIdJson := []byte(`{
	"course_id": "w",
	"lecture_id": "w"
		}`)
	_ = json.Unmarshal(NonNumericCourseIdJson, &testStruct)

	if err := testStruct.ValidateCourseLecture(); err == nil {
		t.Error("not passed", err)
	}

}
