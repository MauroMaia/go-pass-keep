package utils

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

const REGEX_ONLY_NUMBER = "^[0-9]+$"
const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTYVWXYZ0123456789"

// '(,._-*~"<>/|!@#$%^&)+='

var ForbidenWords = []string{
	// TODO - change this to load a well know dictionary
	"qwerty",
}

var WeakWords []string

func init() {
	// Generate alphabet 4 chars sequences
	for i := 0; i < len(ALPHABET)-3; i++ {
		ForbidenWords = append(ForbidenWords,
			string(ALPHABET[i])+string(ALPHABET[i+1])+string(ALPHABET[i+2])+string(ALPHABET[i+3]))
		ForbidenWords = append(ForbidenWords,
			string(ALPHABET[i+3])+string(ALPHABET[i+2])+string(ALPHABET[i+1])+string(ALPHABET[i]))
	}

	WeakWords = append(ForbidenWords, WeakWords...)
}

func ValidatePassword(password string) error {
	passwordSize := len(password)
	if passwordSize == 4 || passwordSize == 6 {
		result, _ := regexp.Match(REGEX_ONLY_NUMBER, []byte(password))
		if !result {
			return errors.New("Invalid pin content. Must contain only numbers.")
		}
		return nil
	}

	if passwordSize < 8 {
		return errors.New(fmt.Sprintf("Invalid Password len(%d). Must be >= 8. If it's a pin must have 4 or 6 digits.", passwordSize))
	}

	for _, word := range ForbidenWords {
		if strings.Contains(password, word) {
			return errors.New(fmt.Sprintf("Invalid password content. Must not contain this word: %s", word))
		}
	}

	return nil
}

func IsPasswordWeak(password string) bool {
	passwordSize := len(password)
	var test = false
	if passwordSize < 16 {
		log.Warn(fmt.Sprintf("Password to short(%d). Should be >= 16. ", passwordSize))
		test = true
	}

	for _, word := range WeakWords {
		if strings.Contains(password, word) {
			log.Warnf("Invalid password content. Must not contain this word: %s", word)
			test = true
		}
	}
	return test
}
