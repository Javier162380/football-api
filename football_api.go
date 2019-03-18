package main

import (
	"encoding/json"
	"football-api/helpers"
	"football-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	PlayerStatsbuffer := helpers.ReadS3Data(helpers.Bucket_name,
		helpers.Players_stats_key)
	PlayertsInformationbuffer := helpers.ReadS3Data(helpers.Bucket_name,
		helpers.Players_information_key)
	GameStatsDatabuffer := helpers.ReadS3Data(helpers.Bucket_name,
		helpers.Games_key)
	PlayerStatsData := GetPlayerStatsData(PlayerStatsbuffer)
	PlayerInformationData := GetPlayerInformationData(PlayertsInformationbuffer)
	GameStatsData := GetGamesStatsData(GameStatsDatabuffer)

	router := mux.NewRouter()
	router.HandleFunc("/{playerid}/{year}",
		GetPlayerYearsStats(PlayerStatsData, PlayerInformationData))

}

func GetGamesStatsData(bytedata []byte) []*models.Games {

	GamesStats := []*models.Games{}
	err := json.Unmarshal(bytedata, &GamesStats)

	if err != nil {
		log.Fatal(err)
	}

	return GamesStats

}

func GetPlayerStatsData(bytedata []byte) []*models.PlayerStats {

	PlayerStats := []*models.PlayerStats{}

	err := json.Unmarshal(bytedata, &PlayerStats)

	if err != nil {
		log.Fatal(err)
	}

	return PlayerStats

}

func GetPlayerInformationData(bytedata []byte) []*models.PlayerInformation {

	PlayerInformation := []*models.PlayerInformation{}

	err := json.Unmarshal(bytedata, &PlayerInformation)

	if err != nil {
		log.Fatal(err)
	}

	return PlayerInformation

}

func GetPlayerYearsStats(PlayerStatsData []*models.PlayerStats,
	PlayerInformationData []*models.PlayerInformation) func(
	w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		for _, player := range PlayerStatsData {
			if player.PlayerID == params["id"] && player.Year == parasm["year"] {
				json.NewEncoder(w).Encode(player)
				return
			}

		}
		json.NewEncoder(w).Encode(models.PlayerStats)
		return
	}
}
