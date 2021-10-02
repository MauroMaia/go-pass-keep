package actions

import (
	log "github.com/sirupsen/logrus"
	"go-pass-keeper/src/model"
)

func CreateVault(filepath string, vaultTitle string, password string) (*model.Vault, error) {

	//TODO - validate inputs

	vault, err := model.NewVault(vaultTitle)
	if err != nil {
		log.Fatal(err)
	}

	err = StoreVault(vault, filepath, password)
	return vault, err
}
