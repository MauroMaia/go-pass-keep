package utils

import (
	"io/ioutil"
	"os"
)

func saveFileBytes(text []byte, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(text)
	return err
}

func readFile(filepath string) ([]byte, error) {
	f, _ := os.Open(filepath)
	defer f.Close()

	return ioutil.ReadAll(f)
}
