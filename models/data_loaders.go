package models

import (
	"encoding/json"
	"log"
)

func GetGamesStatsData(bytedata []byte) []*Games {

	GamesStats := []*Games{}
	err := json.Unmarshal(bytedata, &GamesStats)

	if err != nil {
		log.Fatal(err)
	}

	return GamesStats

}

func GetPlayerStatsData(bytedata []byte) []*PlayerStats {

	PlayerStats := []*PlayerStats{}

	err := json.Unmarshal(bytedata, &PlayerStats)

	if err != nil {
		log.Fatal(err)
	}

	return PlayerStats

}

func GetPlayerInformationData(bytedata []byte) []*PlayerInformation {

	PlayerInformation := []*PlayerInformation{}

	err := json.Unmarshal(bytedata, &PlayerInformation)

	if err != nil {
		log.Fatal(err)
	}

	return PlayerInformation
}
