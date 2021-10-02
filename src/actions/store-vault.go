package actions

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go-pass-keeper/src/model"
	"go-pass-keeper/src/utils"
)

func StoreVault(vault *model.Vault, filepath string, password string) error {

	//TODO - validate inputs

	vaultJsonString, err := json.Marshal(vault)
	if err != nil {
		log.Fatal(err)
	}

	encByteVault, _ := utils.Encrypt(string(vaultJsonString), password)
	return utils.SaveFileBytes(encByteVault, filepath)
}
