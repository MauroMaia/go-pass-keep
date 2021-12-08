package actions

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"go-pass-keeper/src/model"
	"go-pass-keeper/src/utils"
	"strings"
)

func ImportCsvFileToVault(filepath string, vaultInMem model.Vault) (*model.Vault, error) {
	//TODO -validate input

	mapEntries, err := utils.ReadCSVFileToMap(filepath)
	if len(mapEntries) < 2 {
		return nil, errors.New(
			"file must have more than 2 lines with the headers in the firts and the data starting in the 2+",
		)
	}

	headers := mapEntries[0]
	mapEntries = mapEntries[1:]

	for _, mapEntry := range mapEntries {
		entry, err := model.NewEntry(
			mapEntry[1],
			mapEntry[2],
			mapEntry[3],
		)
		if err != nil {
			var password string
			if len(mapEntry[3]) > 3 {
				password = mapEntry[3][0:3] + "..."
			} else {
				password = mapEntry[3]
			}
			log.WithFields(
				log.Fields{
					headers[0]: mapEntry[0],
					headers[1]: mapEntry[1],
					headers[2]: mapEntry[2],
					headers[3]: password,
					headers[4]: mapEntry[4],
					headers[5]: mapEntry[5],
				}).
				Errorf("Failed to generate entry. Reason: %s", err)
			continue
		}
		groupPath := strings.Split(mapEntry[0], "/")
		if contains := vaultInMem.ContainsEntry(entry.GetUsername(), entry.GetTitle(), groupPath); contains {
			var password string
			if len(mapEntry[3]) > 3 {
				password = mapEntry[3][0:3] + "..."
			} else {
				password = mapEntry[3]
			}
			log.WithFields(
				log.Fields{
					headers[0]: mapEntry[0],
					headers[1]: mapEntry[1],
					headers[2]: mapEntry[2],
					headers[3]: password,
					headers[4]: mapEntry[4],
					headers[5]: mapEntry[5],
				}).Warn("Some entry already exist with the same title and username")
			continue
		}

		vaultInMem, err = vaultInMem.PutEntryInVault(entry, groupPath)
		if err != nil {
			return nil, err
		}
	}

	return &vaultInMem, err
}
