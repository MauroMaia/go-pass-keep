package utils

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func SaveFileBytes(text []byte, filepath string) error {
	log.Infof("Saving byte[] to file %s \n", filepath)
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(text)
	return err
}

func ReadFileToBytes(filepath string) ([]byte, error) {
	log.Infof("Reading file %s \n", filepath)
	f, _ := os.Open(filepath)
	defer f.Close()

	return ioutil.ReadAll(f)
}
