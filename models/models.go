package models

type Games struct {
	RequestsUrl                string
	Year                       int
	MatchDate                  string
	Stadium                    string
	Assitance                  string
	LocalTeam                  string
	LocalTeamYellowCards       int
	LocalTeamRedCards          int
	LocalTeamLeaguePosition    string
	LocalTeamballPosition      string
	LocalTeamScoreGoals        int
	VisitingTeam               string
	VisitingTeamYellowCards    int
	VisitingTeamRedCards       int
	VisitingTeamLeaguePosition string
	VisitingTeamballPosition   string
	VisitingTeamScoreGoals     int
	MatchResult                string
}

type PlayerStats struct {
	PlayerID        int
	Name            string
	Year            int
	Competition     string
	Team            string
	GamesPlayed     int
	GamesStarting   int
	GamesCompleted  int
	GamesSustituted int
	MinutePlayed    int
	YellowCards     int
	RedCards        int
	Assits          int
	GoalsScore      int
}

type PlayerInformation struct {
	PlayerID       int
	YearUrl        int
	PlayerName     string
	PlayerFullName string
	BirthPlace     string
	BirthDate      string
	BirthCountry   string
	Nationality    string
	Position       string
	Heigth         string
	Weigth         string
}

type PlayerStatsProfile struct {
	*PlayerStats
	PlayerInformation.PlayerName
	PlayerInformation.PlayerFullName
	PlayerInformation.BirthPlace
	PlayerInformation.BirthDate
	PlayerInformation.BirthCountry
	PlayerInformation.Nationality
	PlayerInformation.Position
	PlayerInformation.Heigth
	PlayerInformation.Weigth
}
