package controllers

import (
	"encoding/json"
	"football-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPlayerYearsStats(PlayerStatsData []*models.PlayerStats, PlayerInformationData []*models.PlayerInformation) func(w http.ResponseWriter, req *http.Request) {
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

func GetTeamStats(PlayerStatsData []*models.PlayerStats, PlayerInformationData *[]models.PlayerInformation) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		team := params["team"]
		year := params["year"]

		var TeamData []models.PlayerStatsProfile

		for _, player := range PlayerStatsData {

			if strconv.Itoa(player.Year) == year &&
				player.Team == team {
				player_id := player.PlayerID
				var TeamMember models.PlayerStatsProfile
				TeamMember.PlayerStats = player

				for _, player := range PlayerInformationData {
					if player.PlayerID == player_id &&
						strconv.Itoa(player.Year) == year {
						TeamMember.PlayerInformation = player
						TeamData = append(TeamData, TeamMember)
					}
				}
			}

		}
		json.NewEncoder(w).Encode(TeamData)
		return

	}
}
