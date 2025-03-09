package vrmlgo

type Player struct {
	User                UserCompact   `json:"user"`
	ThisGame            UserGame      `json:"thisGame"`
	ConnoisseurOverview interface{}   `json:"connoisseurOverview"`
	ConnoisseurHistory  []interface{} `json:"connoisseurHistory"`
}

type BioPastSeason struct {
	TeamMode               string      `json:"teamMode"`
	GameURL                string      `json:"gameUrl"`
	SeasonID               string      `json:"seasonID"`
	SeasonName             string      `json:"seasonName"`
	DivisionLogo           string      `json:"divisionLogo"`
	DivisionName           string      `json:"divisionName"`
	Mmr                    string      `json:"mmr"`
	PlayerID               string      `json:"playerID"`
	UserID                 string      `json:"userID"`
	PlayerName             string      `json:"playerName"`
	UserLogo               string      `json:"userLogo"`
	Country                string      `json:"country"`
	Nationality            string      `json:"nationality"`
	RoleID                 string      `json:"roleID"`
	Role                   string      `json:"role"`
	IsTeamOwner            bool        `json:"isTeamOwner"`
	IsTeamStarter          bool        `json:"isTeamStarter"`
	TeamID                 string      `json:"teamID"`
	TeamName               string      `json:"teamName"`
	TeamNameFull           string      `json:"teamNameFull"`
	TeamLogo               string      `json:"teamLogo"`
	HonoursMention         interface{} `json:"honoursMention"`
	HonoursMentionLogo     interface{} `json:"honoursMentionLogo"`
	CooldownID             interface{} `json:"cooldownID"`
	CooldownNote           interface{} `json:"cooldownNote"`
	CooldownDateExpiresUTC interface{} `json:"cooldownDateExpiresUTC"`
}

type GameCompact struct {
	GameID         string `json:"gameID"`
	GameName       string `json:"gameName"`
	TeamMode       string `json:"teamMode"`
	MatchMode      string `json:"matchMode"`
	URL            string `json:"url"`
	URLShort       string `json:"urlShort"`
	URLComplete    string `json:"urlComplete"`
	HasSubstitutes bool   `json:"hasSubstitutes"`
	HasTies        bool   `json:"hasTies"`
	HasCasters     bool   `json:"hasCasters"`
	HasCameraman   bool   `json:"hasCameraman"`
}

type UserCompact struct {
	UserID        string      `json:"userID"`
	UserName      string      `json:"userName"`
	UserLogo      string      `json:"userLogo"`
	Country       string      `json:"country"`
	Nationality   string      `json:"nationality"`
	DateJoinedUTC string      `json:"dateJoinedUTC"`
	StreamURL     interface{} `json:"streamUrl"`
	DiscordID     float64     `json:"discordID"`
	DiscordTag    string      `json:"discordTag"`
	SteamID       interface{} `json:"steamID"`
	IsTerminated  bool        `json:"isTerminated"`
}

type PlayerSearchResult struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
