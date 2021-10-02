package utils

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"testing"
	"time"
)

var randomText string
var randomPassword string

const FILE_PATH = "./teste.txt"

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

func Test_SaveToFile(t *testing.T) {
	type args struct {
		content  string
		filepath string
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Success test case", args{content: randomText, filepath: FILE_PATH, password: randomPassword}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SaveToFile(tt.args.content, tt.args.filepath, tt.args.password)
			if _, err := os.Stat(FILE_PATH); !os.IsNotExist(err) {
				os.Remove(FILE_PATH)
			}
			if err != nil {
				log.Error(err)
				t.FailNow()
			}
		})
	}
}

func Test_LoadFromFile(t *testing.T) {
	type args struct {
		filepath string
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Success test case", args{FILE_PATH, randomPassword}, randomText},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadFromFile(tt.args.filepath, tt.args.password)
			if err != nil {
				log.Error(err)
				t.FailNow()
			}
			if got != tt.want {
				t.Errorf("LoadFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
