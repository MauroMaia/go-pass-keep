package model

import (
	"reflect"
	"testing"
)

func TestPassword_GetPassword(t *testing.T) {
	type fields struct {
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Success test case", fields{password: "Mauro"}, "Mauro", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Password{
				password: tt.fields.password,
			}
			got, err := p.GetPassword()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword_SetPassword(t *testing.T) {
	type fields struct {
		password string
	}
	type args struct {
		newPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Password
		wantErr bool
	}{
		{"Fail test case. Password < 8 ", fields{password: "875621"}, args{newPassword: "123"}, &Password{password: "875621"}, true},
		{"Success test case", fields{password: "875621"}, args{newPassword: "875621149"}, &Password{password: "875621149"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Password{
				password: tt.fields.password,
			}
			got, err := p.SetPassword(tt.args.newPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}
