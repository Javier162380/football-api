package main

import (
	"football-api/controllers"
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
	PlayerStatsData := models.GetPlayerStatsData(PlayerStatsbuffer)
	PlayerInformationData := models.GetPlayerInformationData(PlayertsInformationbuffer)

	router := mux.NewRouter()
	router.HandleFunc("/{playerid}/{year}",
		controllers.GetPlayerYearsStats(PlayerStatsData, PlayerInformationData)).Methods("GET")
	router.HandleFunc("/{team}/{year}",
		controllers.GetTeamStats(PlayerStatsData, PlayerInformationData)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}
