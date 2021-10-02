package actions

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go-pass-keeper/src/model"
	"go-pass-keeper/src/utils"
)

func LoadVault(filepath string, password string) (*model.Vault, error) {

	//TODO - validate inputs

	encVault, err := utils.ReadFileToBytes(filepath)
	if err != nil {
		return nil, err
	}

	vaultJsonString, err := utils.Decrypt(string(encVault), password)

	var vault model.Vault
	err = json.Unmarshal([]byte(vaultJsonString), &vault)
	if err != nil {
		log.Fatal(err)
	}

	return &vault, nil
}
