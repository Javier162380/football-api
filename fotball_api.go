package main

import (
	"encoding/json"
	"football-api/helpers"
	"football-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetGamesStatsData(bytedata []byte) []*models.Games {

	GamesStats := []*models.Games{}
	err := json.Unmarshal(bytedata, &GamesStats)

	if err != nil {
		log.Fatal(err)
	}

	return GamesStats

}

func GetPlayerInformationData(bytedata []byte) []*models.PlayerStats {

	PlayerStats := []*models.PlayerStats{}

	err := json.Unmarshal(bytedata, &PlayerStats)

	if err != nil {
		log.Fatal(err)
	}

	return PlayerStats

}

func GetPlayerYearStats(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, player := range PlayerIndormationData {

	}


}

var PlayerInformationbuffer = helpers.ReadS3Data(helpers.Bucket_name,
	helpers.Players_information_key)
var GameStatsDatabuffer = helpers.ReadS3Data(helpers.Bucket_name,
	helpers.Games_key)
var PlayerInformationData = GetPlayerInformationData(PlayerInformationbuffer)
var GameStatsData = GetGamesStatsData(GameStatsDatabuffer)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/{playerid}/{year}", GetPlayerYearStats).Methods("GET")

}
