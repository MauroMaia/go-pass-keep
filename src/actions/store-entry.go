package actions

import (
	"encoding/json"

	"go-pass-keeper/src/model"

	log "github.com/sirupsen/logrus"
)

//import "go-pass-keeper/src/utils"

func StoreEntry(user string) {
	// TODO - this should be like:  func StoreEntry(user string,vault *Vault)

	//Todo validate inputs
	// TODO validate if exist
	// TODO load vault

	entry, err := model.NewEntry(
		"",
		user,
		"1231233123548",
	)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(entry)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(b))

	// TODO - Store the vault again
	//utils.SaveToFile(tt.args.content, tt.args.filepath, tt.args.password)
}
