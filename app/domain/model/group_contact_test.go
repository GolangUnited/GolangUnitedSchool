package model

import (
	"encoding/json"
	"testing"
)

func TestGroupContact_ValidateGroupContact(t *testing.T) {
	type fields struct {
		GroupContactId int64
		GroupId        int64
		ContactTypeId  int64
		IsPrimary      bool
		ContactValue   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "valid data with telegram",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  1,
				IsPrimary:      true,
				ContactValue:   "@kekarj",
			},
			wantErr: false,
		},
		{
			name: "valid data with telegram group",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  2,
				IsPrimary:      true,
				ContactValue:   "t.me/go_united_chat",
			},
			wantErr: false,
		},
		{
			name: "valid data with discord person",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  3,
				IsPrimary:      true,
				ContactValue:   "kek#7546",
			},
			wantErr: false,
		},
		{
			name: "valid data with discord channel",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  4,
				IsPrimary:      true,
				ContactValue:   "https://discord.com/channels/516715744646660106/1012067296744841246",
			},
			wantErr: false,
		},
		{
			name: "valid data with email",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  5,
				IsPrimary:      true,
				ContactValue:   "kek@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "valid data with phone number",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  6,
				IsPrimary:      true,
				ContactValue:   "+79080490909",
			},
			wantErr: false,
		},
		{
			name: "valid data with vk profile",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  7,
				IsPrimary:      true,
				ContactValue:   "vk.com/funeral",
			},
			wantErr: false,
		},
		{
			name: "less than valid contact type id",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  0,
				IsPrimary:      true,
				ContactValue:   "kek@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "more than valid contact type id",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  8,
				IsPrimary:      true,
				ContactValue:   "kek@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "empty group contact id",
			fields: fields{
				//GroupContactId: 1,
				GroupId:       1,
				ContactTypeId: 5,
				IsPrimary:     true,
				ContactValue:  "kek@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "empty group id",
			fields: fields{
				GroupContactId: 1,
				//GroupId:        1,
				ContactTypeId: 5,
				IsPrimary:     true,
				ContactValue:  "kek@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "empty contact type id",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				//ContactTypeId:  7,
				IsPrimary:    true,
				ContactValue: "kek@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "empty group id",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  7,
				//IsPrimary:      true,
				ContactValue: "kek@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "empty contact value",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  7,
				IsPrimary:      true,
				//ContactValue: "kek@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "too short contact value",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  5,
				IsPrimary:      true,
				ContactValue:   "k",
			},
			wantErr: true,
		},
		{
			name: "too long contact value",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  4,
				IsPrimary:      true,
				ContactValue:   "https://discord.com/channels/516715744646660106/1012067296744841246kjhgkjhgkjhgkjhgkjhg",
			},
			wantErr: true,
		},
		{
			name: "valid telegram contact",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  1,
				IsPrimary:      true,
				ContactValue:   "@qwe",
			},
			wantErr: false,
		},
		{
			name: "invalid telegram contact",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  1,
				IsPrimary:      true,
				ContactValue:   "@qwer^^ty",
			},
			wantErr: true,
		},
		{
			name: "valid telegram group",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  2,
				IsPrimary:      true,
				ContactValue:   "t.me/kek_kekek",
			},
			wantErr: false,
		},
		{
			name: "invalid telegram group",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  2,
				IsPrimary:      true,
				ContactValue:   "t.mekek_kekek",
			},
			wantErr: true,
		},
		{
			name: "valid discord username",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  3,
				IsPrimary:      true,
				ContactValue:   "kek#777",
			},
			wantErr: false,
		},
		{
			name: "invalid discord username",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  3,
				IsPrimary:      true,
				ContactValue:   "kek777",
			},
			wantErr: true,
		},
		{
			name: "valid discord channel",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  4,
				IsPrimary:      true,
				ContactValue:   "https://discord.com/channels/516715744646660106/1012067296744841246",
			},
			wantErr: false,
		},
		{
			name: "invalid discord channel",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  4,
				IsPrimary:      true,
				ContactValue:   "https/discord.com/channels/516715744646660106/1012067296744841246",
			},
			wantErr: true,
		},
		{
			name: "too long discord channel",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  4,
				IsPrimary:      true,
				ContactValue:   "https://discord.com/channels/5167157888888888844646660106/1012067296744841246",
			},
			wantErr: true,
		},
		{
			name: "valid email",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  5,
				IsPrimary:      true,
				ContactValue:   "funly@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  5,
				IsPrimary:      true,
				ContactValue:   "funly@gmailcom",
			},
			wantErr: true,
		},
		{
			name: "valid phone",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  6,
				IsPrimary:      true,
				ContactValue:   "+79080490407",
			},
			wantErr: false,
		},
		{
			name: "invalid phone",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  6,
				IsPrimary:      true,
				ContactValue:   "+7908049407",
			},
			wantErr: true,
		},
		{
			name: "valid vk",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  7,
				IsPrimary:      true,
				ContactValue:   "vk.com/kek",
			},
			wantErr: false,
		},
		{
			name: "invalid vk",
			fields: fields{
				GroupContactId: 1,
				GroupId:        1,
				ContactTypeId:  7,
				IsPrimary:      true,
				ContactValue:   "vk.com/k:ek",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &GroupContact{
				GroupContactId: tt.fields.GroupContactId,
				GroupId:        tt.fields.GroupId,
				ContactTypeId:  tt.fields.ContactTypeId,
				IsPrimary:      tt.fields.IsPrimary,
				ContactValue:   tt.fields.ContactValue,
			}
			if err := n.ValidateGroupContact(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateGroupContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGroupContact_nonNumericGroupContactId(t *testing.T) {

	var testStruct GroupContact

	nonNumericGroupContactId := []byte(`{
	"group_contact_id": "2",
	"group_id": 1,
	"contact_type_id": 5,
	"is_primary": true,
	"contact_value": "kek@gmail.com"
}`)

	_ = json.Unmarshal(nonNumericGroupContactId, &testStruct)

	if err := testStruct.ValidateGroupContact(); err == nil {
		t.Error(err)
	}

}

func TestGroupContact_nonNumericGroupId(t *testing.T) {
	var testStruct GroupContact
	nonNumericGroupId := []byte(`{
	"group_contact_id": 2,
	"group_id": "w",
	"contact_type_id": 5,
	"is_primary": true,
	"contact_value": "kek@gmail.com"
}`)

	_ = json.Unmarshal(nonNumericGroupId, &testStruct)

	if err := testStruct.ValidateGroupContact(); err == nil {
		t.Error(err)
	}

}

func TestGroupContact_nonNumericContactTypeId(t *testing.T) {
	var testStruct GroupContact
	nonNumericGroupId := []byte(`{
	"group_contact_id": 2,
	"group_id": 2,
	"contact_type_id": "ww",
	"is_primary": true,
	"contact_value": "kek@gmail.com"
}`)

	_ = json.Unmarshal(nonNumericGroupId, &testStruct)

	if err := testStruct.ValidateGroupContact(); err == nil {
		t.Error(err)
	}

}

func TestGroupContact_nonBooleanIsPrimary(t *testing.T) {
	var testStruct GroupContact
	nonNumericGroupId := []byte(`{
	"group_contact_id": 2,
	"group_id": 2,
	"contact_type_id": 5,
	"is_primary": "yh",
	"contact_value": "kek@gmail.com"
}`)

	_ = json.Unmarshal(nonNumericGroupId, &testStruct)

	if err := testStruct.ValidateGroupContact(); err == nil {
		t.Error(err)
	}

}
