package utils

import "testing"

func TestValidatePassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Pin code 4d - sucessfull", args{password: "8529"}, false},
		{"Pin code 6d - sucessfull", args{password: "984560"}, false},
		{"StrongPassword - sucessfull", args{password: "correctho&rsebatterystaple"}, false},
		{"Fail - to small", args{password: "123"}, true},
		{"Check for sequences 1 - Fail", args{password: "123456789"}, true},
		{"Check for sequences 2 - sucessfull", args{password: "125486234"}, false},
		{"Check for sequences 3 - Fail", args{password: "qwertyqwerty"}, true},
		{"Check for sequences 3.1 - Fail", args{password: "Qwertyqwerty"}, true},
		{"Check for sequences 4 - sucessfull", args{password: "987462158"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePassword(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsPasswordWeak(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"StrongPassword - sucessfull", args{password: "correctho&rsebatterystaple"}, false},
		{"Fail - to small", args{password: "123"}, true},
		{"Check for sequences 1 - Fail", args{password: "123456789"}, true},
		{"Check for sequences 2 - Fail", args{password: "125486234"}, true},
		{"Check for sequences 3 - Fail", args{password: "qwertyqwerty"}, true},
		{"Check for sequences 3.1 - Fail", args{password: "Qwertyqwerty"}, true},
		{"Check for sequences 4 - Fail", args{password: "987462158"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPasswordWeak(tt.args.password); got != tt.want {
				t.Errorf("IsPasswordWeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
