package utils

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"testing"
	"time"
)

var randomText string
var randomPassword string

func init() {
	rand.Seed(time.Now().UnixNano())
	randomText = RandStringRunes(300000)
	randomPassword = RandStringRunes(16)
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Test_Encrypt(t *testing.T) {
	type args struct {
		content  string
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Success test case", args{content: randomText, password: randomPassword}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Encrypt(tt.args.content, tt.args.password)
			if err != nil {
				log.Error(err)
				t.FailNow()
			}
		})
	}
}

func Test_Decrypt(t *testing.T) {
	type args struct {
		ciphertext string
		password   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Success test case",
			args{
				func() string {
					t, _ := Encrypt(randomText, randomPassword)
					return string(t)
				}(), randomPassword,
			},
			randomText,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.ciphertext, tt.args.password)
			if err != nil {
				log.Error(err)
				t.FailNow()
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
