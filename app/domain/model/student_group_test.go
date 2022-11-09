package model

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCourseLecture_Validate(t *testing.T) {
	var testStruct CourseLecture

	validJson := []byte(`{
	"course_id": 1,
	"lecture_id": 2
	
	}`)

	if err := json.Unmarshal(validJson, &testStruct); err != nil {
		log.Println(err)
	}
	if err := testStruct.ValidateCourseLecture(); err != nil {
		t.Error("not passed", err)
	}

}

func TestCourseLecture_nonValidValues(t *testing.T) {
	wrongCourseIdJson := []byte(`{
		"course_id": "s",
		"lecture_id": 2
}`)

	if err := json.Unmarshal(wrongCourseIdJson, &testStruct); err != nil {

	}
	if err := testStruct.ValidateCourseLecture(); err == nil {
		t.Error("not passed", err)
	}

}

func TestCourseLecture_emptyValues(t *testing.T) {
	wrongCourseIdJson := []byte(`{
		"course_id": 2,
		"lecture_id": null
}`)

	if err := json.Unmarshal(wrongCourseIdJson, &testStruct); err != nil {

	}
	if err := testStruct.ValidateCourseLecture(); err != nil {
		t.Error("not passed", err)
	}

}
