package vrmlgo

type MatchHistory struct {
	MapsSet                   []MapsSet   `json:"mapsSet"`
	SeasonName                string      `json:"seasonName"`
	WinningTeamID             string      `json:"winningTeamID"`
	LosingTeamID              string      `json:"losingTeamID"`
	HomeScore                 int64       `json:"homeScore"`
	AwayScore                 int64       `json:"awayScore"`
	IsTie                     bool        `json:"isTie"`
	IsForfeit                 bool        `json:"isForfeit"`
	MatchID                   string      `json:"matchID"`
	Week                      int64       `json:"week"`
	IsScheduled               bool        `json:"isScheduled"`
	IsSpecificDivision        bool        `json:"isSpecificDivision"`
	IsChallenge               bool        `json:"isChallenge"`
	IsCup                     bool        `json:"isCup"`
	DateScheduledUTC          string      `json:"dateScheduledUTC"`
	DateScheduledUser         string      `json:"dateScheduledUser"`
	DateScheduledUserTimezone string      `json:"dateScheduledUserTimezone"`
	HomeTeam                  MatchTeam   `json:"homeTeam"`
	AwayTeam                  MatchTeam   `json:"awayTeam"`
	CastingInfo               CastingInfo `json:"castingInfo"`
	VODURL                    *string     `json:"vodUrl"`
	HomeHighlights            interface{} `json:"homeHighlights"`
	AwayHighlights            interface{} `json:"awayHighlights"`
	PostponeTeamID            *string     `json:"postponeTeamID"`
	ModsReview                bool        `json:"modsReview"`
	ModsReviewNote            interface{} `json:"modsReviewNote"`
}

type MapsSet struct {
	MapNum    int64  `json:"mapNum"`
	MapName   string `json:"mapName"`
	HomeScore int64  `json:"homeScore"`
	AwayScore int64  `json:"awayScore"`
}
