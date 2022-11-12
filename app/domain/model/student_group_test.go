package model

import (
	"encoding/json"
	"log"
	"testing"
)

func TestStudentGroupValidate(t *testing.T) {
	var testStruct StudentGroup

	validJson := []byte(`{
	"student_id": 1,
	"group_id": 2
	
	}`)

	if err := json.Unmarshal(validJson, &testStruct); err != nil {
		log.Println(err)
	}
	if err := testStruct.ValidateStudentGroup(); err != nil {
		t.Error("not passed", err)
	}

}

func TestStudentGroupStudentId(t *testing.T) {
	wrongCourseIdJson := []byte(`{
		"student_id": "s",
		"group_id": 2
}`)

	if err := json.Unmarshal(wrongCourseIdJson, &testStruct); err != nil {

	}
	if err := testStruct.ValidateCourseLecture(); err == nil {
		t.Error("not passed", err)
	}

}

func TestStudentGroupId(t *testing.T) {
	wrongCourseIdJson := []byte(`{
		"student_id": 2,
		"group_id": null
}`)

	if err := json.Unmarshal(wrongCourseIdJson, &testStruct); err != nil {

	}
	if err := testStruct.ValidateCourseLecture(); err == nil {
		t.Error("not passed", err)
	}

}
