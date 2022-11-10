package model

import (
	"encoding/json"
	"testing"
	"time"
)

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
		wantErr bool
	}{
		{
			name: "valid_data",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: false,
		},
		{
			name: "empty_firstName",
			fields: fields{
				FirstName:  "",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "too short firstName",
			fields: fields{
				FirstName:  "v",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "too long firstName",
			fields: fields{
				FirstName:  "vasilvasilvasilvasilvasilvasilvasilvasilvasilvasilvasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "cyrillic firstName",
			fields: fields{
				FirstName:  "вася",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: false,
		},
		{
			name: "empty lastName",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "",
				Patronymic: "vasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "too short lastName",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "v",
				Patronymic: "vasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "too long lastName",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasilievvasilievvasilievvasilievvasilievvasilievvasiliev",
				Patronymic: "vasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "cyrillic lastName",
			fields: fields{
				FirstName:  "vasya",
				LastName:   "васильев",
				Patronymic: "vasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: false,
		},
		{
			name: "empty patronymic",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: false,
		},
		{
			name: "short patronymic",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "j",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "long patronymic",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievichvasilievichvasilievichvasilievichvasilievich",
				Login:      "vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "empty login",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "short login",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "v",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    false,
			},
			wantErr: true,
		},
		{
			name: "long login",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "vasya12345vasya12345vasya12345vasya12345vasya12345vasya12345",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    true,
			},
			wantErr: true,
		},
		{
			name: "not ascii login",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "лорплорп",
				RoleId:     2,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    true,
			},
			wantErr: true,
		},
		{
			name: "zero roleId",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "qwerty12345",
				RoleId:     0,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    true,
			},
			wantErr: true,
		},
		{
			name: "12 roleId",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "qwerty12345",
				RoleId:     12,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    true,
			},
			wantErr: true,
		},
		{
			name: "0 roleId",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "qwerty12345",
				RoleId:     0,
				Passwd:     "kek21212121",
				UpdatedAt:  time.Now(),
				Deleted:    true,
			},
			wantErr: true,
		},
		{
			name: "empty passwd",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "qwerty12345",
				RoleId:     2,
				Passwd:     "",
				UpdatedAt:  time.Now(),
				Deleted:    true,
			},
			wantErr: true,
		},
		{
			name: "short passwd",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "qwerty12345",
				RoleId:     2,
				Passwd:     "jj",
				UpdatedAt:  time.Now(),
				Deleted:    true,
			},
			wantErr: true,
		},
		{
			name: "long passwd",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "qwerty12345",
				RoleId:     2,
				Passwd:     "12345678901234567890123456789012345678901234567890",
				UpdatedAt:  time.Now(),
				Deleted:    true,
			},
			wantErr: true,
		},
		{
			name: "non ascii passwd",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "qwerty12345",
				RoleId:     2,
				Passwd:     "фыва",
				UpdatedAt:  time.Now(),
				Deleted:    true,
			},

			wantErr: true,
		},
		{
			name: "omitempty time",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "qwerty12345",
				RoleId:     2,
				Passwd:     "qwertrtddd",
				Deleted:    true,
			},

			wantErr: false,
		},
		{
			name: "omitempty deleted",
			fields: fields{
				FirstName:  "vasil",
				LastName:   "vasiliev",
				Patronymic: "vasilievich",
				Login:      "qwerty12345",
				RoleId:     2,
				Passwd:     "qwertrtddd",
				UpdatedAt:  time.Now(),
			},

			wantErr: false,
		},
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
			if err := n.ValidatePerson(); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePerson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// // исправить
var testPerson Person

func TestPerson_nonAsciiLogin(t *testing.T) {

	LongLastNameJson := []byte(`{
	"first_name": "vasil",
	"last_name": "qqqqqqqq",
	"patronymic": "vasiliveich",
	"login": "12лlj2",
	"role_id": 1,
	"passwd": "12345555",
	"updated_at": "",
	"deleted": false
	}`)

	_ = json.Unmarshal(LongLastNameJson, &testPerson)
	if err := testPerson.ValidatePerson(); err == nil {
		t.Error("not passed")
	}

}

// Интересный момент,  при использовании конструкции assert тесты падают в панику
// без использования ассертов именно на этом тесте с передачей в жсон невалидныъ
// данных
func TestPerson_nonNumericRoleId(t *testing.T) {
	NameJson := []byte(`{
	"first_name": "vasil",
	"last_name": "qqqqqqqq",
	"patronymic": "vasiliveich",
	"login": "qwelrty",
	"role_id": 0,
	"passwd": "123kk456",
	"updated_at": "2006-01-02T15:04:05Z",
	"deleted": false
	}`)

	_ = json.Unmarshal(NameJson, &testPerson)
	if err := testPerson.ValidatePerson(); err == nil {
		t.Error("not passed", err)
	}

}

/*возможно, стоило бы расписать кейсы по валидации неверных полей
для времени и удаления, но я не буду заморачиваться*/
