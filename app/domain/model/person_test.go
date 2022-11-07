package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestPerson_Name(t *testing.T) {
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

	LongLastNameJson := []byte(`{
	"first_name": "vasil",
	"last_name": "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
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
