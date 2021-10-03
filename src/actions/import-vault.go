package actions

import (
	log "github.com/sirupsen/logrus"
	"go-pass-keeper/src/model"
	"go-pass-keeper/src/utils"
)

func ReadCSVFileToEntryList(filepath string) ([]*model.Entry, error) {
	//TODO -validate input

	mapEntries, err := utils.ReadCSVFileToMap(filepath)
	mapEntriesValues := mapEntries[1:]

	var result []*model.Entry
	for _, mapEntry := range mapEntriesValues {
		entry, err := model.NewEntry(
			// mapEntry[0], //path
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
					mapEntries[0][0]: mapEntry[0],
					mapEntries[0][1]: mapEntry[1],
					mapEntries[0][2]: mapEntry[2],
					mapEntries[0][3]: password,
					mapEntries[0][4]: mapEntry[4],
					mapEntries[0][5]: mapEntry[5],
				}).
				Errorf("Failed to generate entry. Reason: %s", err)
			continue
		}
		result = append(result, entry)
	}

	return result, err
}
