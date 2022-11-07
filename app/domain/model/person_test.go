package model

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testPerson Person

func TestPerson_ValidPerson(t *testing.T) {

	validJson := []byte(`{
	"first_name": "vasil",
	"last_name": "vasiliev",
	"patronymic": "vasiliveich",
	"login": "12lkj2",
	"role_id": 1,
	"passwd": "123kk456",
	"updated_at": "2006-01-02T15:04:05Z",
	"deleted": false
	}`)

	_ = json.Unmarshal(validJson, &testPerson)
	assert.Nil(t, testPerson.ValidatePerson())

}

func TestPerson_EmptyName(t *testing.T) {
	emptyNameJson := []byte(`{
	"first_name": "",
	"last_name": "vasiliev",
	"patronymic": "vasiliveich",
	"login": "12lkj2",
	"role_id": 1,
	"passwd": "123kk456",
	"updated_at": "2006-01-02T15:04:05Z",
	"deleted": false
	}`)

	_ = json.Unmarshal(emptyNameJson, &testPerson)
	assert.Equal(t,
		"[FirstName is a required field]",
		testPerson.ValidatePerson().Error())
}

func TestPerson_ShortName(t *testing.T) {
	shortNameJson := []byte(`{
	"first_name": "v",
	"last_name": "vasiliev",
	"patronymic": "vasiliveich",
	"login": "12lkj2",
	"role_id": 1,
	"passwd": "123kk456",
	"updated_at": "2006-01-02T15:04:05Z",
	"deleted": false
	}`)

	_ = json.Unmarshal(shortNameJson, &testPerson)
	assert.Equal(t,
		"[FirstName must be at least 2 characters in length]",
		testPerson.ValidatePerson().Error())

}

func TestPerson_LongName(t *testing.T) {

	longNameJson := []byte(`{
	"first_name": "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh",
	"last_name": "vasiliev",
	"patronymic": "vasiliveich",
	"login": "12lkj2",
	"role_id": 1,
	"passwd": "123kk456",
	"updated_at": "2006-01-02T15:04:05Z",
	"deleted": false
	}`)

	_ = json.Unmarshal(longNameJson, &testPerson)
	assert.Equal(t,
		"[FirstName must be a maximum of 50 characters in length]",
		testPerson.ValidatePerson().Error())

}

func TestPerson_EmptyLastName(t *testing.T) {

	EmptyLastNameJson := []byte(`{
	"first_name": "vasil",
	"last_name": "",
	"patronymic": "vasiliveich",
	"login": "12lkj2",
	"role_id": 1,
	"passwd": "123kk456",
	"updated_at": "2006-01-02T15:04:05Z",
	"deleted": false
	}`)

	_ = json.Unmarshal(EmptyLastNameJson, &testPerson)
	assert.Equal(t,
		"[LastName is a required field]",
		testPerson.ValidatePerson().Error())

}

func TestPerson_longShortLastName(t *testing.T) {

	ShortLastNameJson := []byte(`{
	"first_name": "vasil",
	"last_name": "q",
	"patronymic": "vasiliveich",
	"login": "12lkj2",
	"role_id": 1,
	"passwd": "123kk456",
	"updated_at": "2006-01-02T15:04:05Z",
	"deleted": false
	}`)

	_ = json.Unmarshal(ShortLastNameJson, &testPerson)
	assert.Equal(t,
		"[LastName must be at least 2 characters in length]",
		testPerson.ValidatePerson().Error())
}

func TestPerson_LongLastName(t *testing.T) {

	LongLastNameJson := []byte(`{
	"first_name": "vasil",
	"last_name": "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
	"patronymic": "vasiliveich",
	"login": "12lkj2",
	"role_id": 1,
	"passwd": "123kk456",
	"updated_at": "2006-01-02T15:04:05Z",
	"deleted": false
	}`)

	_ = json.Unmarshal(LongLastNameJson, &testPerson)
	assert.Equal(t,
		"[LastName must be a maximum of 50 characters in length]",
		testPerson.ValidatePerson().Error())

}

func TestPerson_ValidatePerson(t *testing.T) {
	type fields struct {
		PersonId   int64
		FirstName  string
		LastName   string
		Patronymic string
		Login      string
		RoleId     int
		Passwd     string
		UpdatedAt  time.Time
		Deleted    bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Person{
				PersonId:   tt.fields.PersonId,
				FirstName:  tt.fields.FirstName,
				LastName:   tt.fields.LastName,
				Patronymic: tt.fields.Patronymic,
				Login:      tt.fields.Login,
				RoleId:     tt.fields.RoleId,
				Passwd:     tt.fields.Passwd,
				UpdatedAt:  tt.fields.UpdatedAt,
				Deleted:    tt.fields.Deleted,
			}
			tt.wantErr(t, n.ValidatePerson(), fmt.Sprintf("ValidatePerson()"))
		})
	}
}
