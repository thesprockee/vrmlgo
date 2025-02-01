package vrmlgo

import "strconv"

type User struct {
	ID               string `json:"userID"`
	Email            string `json:"email,omitempty"`
	UserName         string `json:"userName"`
	UserLogo         string `json:"userLogo"`
	Country          string `json:"country"`
	Nationality      string `json:"nationality"`
	DateJoinedUTC    string `json:"dateJoinedUTC"`
	Theme            string `json:"theme"`
	TimezoneID       string `json:"timezoneID"`
	DiscordID        uint64 `json:"discordID"`
	DiscordTag       string `json:"discordTag"`
	SteamID          string `json:"steamID"`
	OculusHomeID     string `json:"oculusHomeID"`
	OculusHomeAlias  string `json:"oculusHomeAlias"`
	StreamURL        string `json:"streamURL"`
	IsUsingDarkTheme bool   `json:"isUsingDarkTheme"`
}

func (u User) GetDiscordID() string {
	return strconv.FormatUint(u.DiscordID, 10)
}

type UserGames struct {
	PlayerID                  string                            `json:"playerID"`
	PlayerName                string                            `json:"playerName"`
	UserLogo                  string                            `json:"userLogo"`
	Game                      Game                              `json:"game"`
	BioCurrent                BioCurrent                        `json:"bioCurrent"`
	BioCurrentDef             BioCurrent                        `json:"bioCurrentDef"`
	BioCurrentSeasonPastTeams []BioCurrentSeasonPastTeamElement `json:"bioCurrentSeasonPastTeams"`
	BioPastSeasons            []BioCurrentSeasonPastTeamElement `json:"bioPastSeasons"`
}

type UserTeams struct {
	Teams           []TeamMetadata `json:"teams"`
	Substitutes     []interface{}  `json:"substitutes"`
	PendingRecruits []interface{}  `json:"pendingRecruits"`
}

type Region struct {
	ID         string `json:"regionID"`
	RegionName string `json:"regionName"`
	RegionLogo string `json:"regionLogo"`
}

type BioCurrent struct {
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

type BioCurrentSeasonPastTeamElement struct {
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
