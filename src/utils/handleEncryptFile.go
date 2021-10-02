package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io/ioutil"
	"os"
)

func SaveToFile(plainText string, filepath string, password string) error {

	if err := validatePasswordAndContent(plainText, password); err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return err
	}

	var KEY = sha256.Sum256([]byte(password))

	// Create a new AES block cipher.
	block, err := aes.NewCipher(KEY[:])
	if err != nil {
		return err
	}

	// CBC mode always works in whole blocks.
	originData := _PKCS5Padding([]byte(plainText), aes.BlockSize)

	// Create a new CBC mode encrypter using our AES block cipher, and use it
	// to encrypt our text.
	ciphertext := make([]byte, aes.BlockSize+len(originData))
	enc := cipher.NewCBCEncrypter(block, iv)
	enc.CryptBlocks(ciphertext, append(iv[:], originData...))

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(ciphertext)
	return err
}

func LoadFromFile(filepath string, password string) (string, error) {

	if err := validatePassword(password); err != nil {
		return "", err
	}

	f, _ := os.Open(filepath)
	defer f.Close()

	ciphertext, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	var KEY = sha256.Sum256([]byte(password))

	block, err := aes.NewCipher(KEY[:])
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)
	plainText := _PKCS5UnPadding(ciphertext)
	return string(plainText), nil
}

func validatePasswordAndContent(content string, password string) error {

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(content) < aes.BlockSize {
		return errors.New("ciphertext too short")
	}

	return validatePassword(password)
}
func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password too short")
	}
	return nil
}

func _PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func _PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
