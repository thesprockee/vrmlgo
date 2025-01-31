package vrmlgo

type MatchDetails struct {
	Match       Match          `json:"match"`
	PlayersHome []*MatchPlayer `json:"playersHome"`
	PlayersAway []*MatchPlayer `json:"playersAway"`
	//Season               Season            `json:"season"` // This (incorrectly?) returns the latest season
	CurrentUserTeamID    string            `json:"currentUserTeamID"`
	HomeTeamRank         int64             `json:"homeTeamRank"`
	AwayTeamRank         int64             `json:"awayTeamRank"`
	PastMatchesCommon    []Match           `json:"pastMatchesCommon"`
	StatsMapsCommon      []StatsMapsCommon `json:"statsMapsCommon"`
	StatsMapsHome        []StatsMaps       `json:"statsMapsHome"`
	StatsMapsAway        []StatsMaps       `json:"statsMapsAway"`
	ModsMatchInformation interface{}       `json:"modsMatchInformation"`
	Streamers            interface{}       `json:"streamers"`
	Streamer             interface{}       `json:"streamer"`
}

func (m *MatchDetails) Players() []*MatchPlayer {
	return append(m.PlayersHome, m.PlayersAway...)
}

type Match struct {
	ID                        string      `json:"matchID"`
	SeasonName                string      `json:"seasonName"`
	WinningTeamID             string      `json:"winningTeamID"`
	LosingTeamID              string      `json:"losingTeamID"`
	HomeScore                 int64       `json:"homeScore"`
	AwayScore                 int64       `json:"awayScore"`
	IsTie                     bool        `json:"isTie"`
	IsForfeit                 bool        `json:"isForfeit"`
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
	VODURL                    string      `json:"vodUrl"`
	HomeHighlights            interface{} `json:"homeHighlights"`
	AwayHighlights            interface{} `json:"awayHighlights"`
	PostponeTeamID            interface{} `json:"postponeTeamID"`
	ModsReview                bool        `json:"modsReview"`
	ModsReviewNote            interface{} `json:"modsReviewNote"`
}

type MatchTeam struct {
	SubmittedScores bool   `json:"submittedScores"`
	DivisionName    string `json:"divisionName"`
	DivisionLogo    string `json:"divisionLogo"`
	TeamID          string `json:"teamID"`
	TeamName        string `json:"teamName"`
	TeamLogo        string `json:"teamLogo"`
	RegionID        string `json:"regionID"`
	RegionName      string `json:"regionName"`
}

type CastingInfo struct {
	ChannelType           *int64      `json:"channelType"`
	ChannelID             *string     `json:"channelID"`
	ChannelURL            *string     `json:"channelURL"`
	CasterID              *string     `json:"casterID"`
	Caster                *string     `json:"caster"`
	CasterLogo            *string     `json:"casterLogo"`
	CoCasterID            *string     `json:"coCasterID"`
	CoCaster              *string     `json:"coCaster"`
	CoCasterLogo          *string     `json:"coCasterLogo"`
	PostGameInterviewID   interface{} `json:"postGameInterviewID"`
	PostGameInterview     interface{} `json:"postGameInterview"`
	PostGameInterviewLogo interface{} `json:"postGameInterviewLogo"`
	CameramanID           *string     `json:"cameramanID"`
	Cameraman             *string     `json:"cameraman"`
	CameramanLogo         *string     `json:"cameramanLogo"`
}

type MatchPlayer struct {
	ID                            string  `json:"playerID"`
	PlayerName                    string  `json:"playerName"`
	UserID                        string  `json:"userID"`
	UserLogo                      string  `json:"userLogo"`
	Country                       string  `json:"country"`
	Nationality                   *string `json:"nationality"`
	StreamURL                     *string `json:"streamURL"`
	TeamID                        string  `json:"teamID"`
	TeamName                      string  `json:"teamName"`
	RoleID                        string  `json:"roleID"`
	Role                          string  `json:"role"`
	IsTeamOwner                   bool    `json:"isTeamOwner"`
	IsTeamOwnerCaptainOrCoCaptain bool    `json:"isTeamOwnerCaptainOrCoCaptain"`
	IsTeamStarter                 bool    `json:"isTeamStarter"`
}

type StatsMaps struct {
	MapName             string `json:"mapName"`
	Played              int64  `json:"played"`
	Win                 int64  `json:"win"`
	WinPercentage       int64  `json:"winPercentage"`
	RoundsPlayed        int64  `json:"roundsPlayed"`
	RoundsWin           int64  `json:"roundsWin"`
	RoundsWinPercentage int64  `json:"roundsWinPercentage"`
}

type StatsMapsCommon struct {
	MapName                  string `json:"mapName"`
	Played                   int64  `json:"played"`
	Team1Win                 int64  `json:"team1Win"`
	Team1WinPercentage       int64  `json:"team1WinPercentage"`
	Team1RoundsWin           int64  `json:"team1RoundsWin"`
	Team1RoundsWinPercentage int64  `json:"team1RoundsWinPercentage"`
	Team2Win                 int64  `json:"team2Win"`
	Team2WinPercentage       int64  `json:"team2WinPercentage"`
	Team2RoundsWin           int64  `json:"team2RoundsWin"`
	Team2RoundsWinPercentage int64  `json:"team2RoundsWinPercentage"`
}
