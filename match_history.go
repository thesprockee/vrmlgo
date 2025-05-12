package vrmlgo

type MatchHistory struct {
	MapsSet                   []MapsSet   `json:"mapsSet"`
	SeasonName                string      `json:"seasonName"`
	WinningTeamID             string      `json:"winningTeamID"`
	LosingTeamID              string      `json:"losingTeamID"`
	HomeScore                 int         `json:"homeScore"`
	AwayScore                 int         `json:"awayScore"`
	IsTie                     bool        `json:"isTie"`
	IsForfeit                 bool        `json:"isForfeit"`
	MatchID                   string      `json:"matchID"`
	Week                      int         `json:"week"`
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
	HomeHighlights            any         `json:"homeHighlights"`
	AwayHighlights            any         `json:"awayHighlights"`
	PostponeTeamID            *string     `json:"postponeTeamID"`
	ModsReview                bool        `json:"modsReview"`
	ModsReviewNote            any         `json:"modsReviewNote"`
}

type MapsSet struct {
	MapNum    int    `json:"mapNum"`
	MapName   string `json:"mapName"`
	HomeScore int    `json:"homeScore"`
	AwayScore int    `json:"awayScore"`
}
