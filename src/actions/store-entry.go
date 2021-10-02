package actions

import (
	"go-pass-keeper/src/model"

	log "github.com/sirupsen/logrus"
)

func StoreEntry(user string, title string, password string, vault *model.Vault) *model.Vault {

	//Todo validate inputs

	entry, err := model.NewEntry(
		user,
		title,
		password,
	)
	if err != nil {
		log.Fatal(err)
	}

	return vault.PutEntry(entry)
}
