package models

type Games []struct {
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
