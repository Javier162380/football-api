package main

import (
	"encoding/json"
	"football-api/helpers"
	"football-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	PlayerStatsbuffer := helpers.ReadS3Data(helpers.Bucket_name,
		helpers.Players_stats_key)
	PlayertsInformationbuffer := helpers.ReadS3Data(helpers.Bucket_name,
		helpers.Players_information_key)
	PlayerStatsData := GetPlayerStatsData(PlayerStatsbuffer)
	PlayerInformationData := GetPlayerInformationData(PlayertsInformationbuffer)

	router := mux.NewRouter()
	router.HandleFunc("/{playerid}/{year}",
		GetPlayerYearsStats(PlayerStatsData, PlayerInformationData)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

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
		player_id := params["playerid"]
		year := params["year"]
		var PlayerData models.PlayerStatsProfile
		for _, player := range PlayerStatsData {
			if strconv.Itoa(player.PlayerID) == player_id &&
				strconv.Itoa(player.Year) == year {
				PlayerData.PlayerStats = player
				for _, player := range PlayerInformationData {
					if strconv.Itoa(player.PlayerID) == player_id &&
						strconv.Itoa(player.YearUrl) == year {
						PlayerData.PlayerInformation = player
					}

				}
				json.NewEncoder(w).Encode(PlayerData)
				return
			}

		}
		json.NewEncoder(w).Encode(PlayerData)
		return
	}
}
