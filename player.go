package vrmlgo

import "strconv"

type Player struct {
	User                UserCompact `json:"user"`
	ThisGame            UserGame    `json:"thisGame"`
	ConnoisseurOverview any         `json:"connoisseurOverview"`
	ConnoisseurHistory  []any       `json:"connoisseurHistory"`
}

type BioPastSeason struct {
	TeamMode               string `json:"teamMode"`
	GameURL                string `json:"gameUrl"`
	SeasonID               string `json:"seasonID"`
	SeasonName             string `json:"seasonName"`
	DivisionLogo           string `json:"divisionLogo"`
	DivisionName           string `json:"divisionName"`
	Mmr                    string `json:"mmr"`
	PlayerID               string `json:"playerID"`
	UserID                 string `json:"userID"`
	PlayerName             string `json:"playerName"`
	UserLogo               string `json:"userLogo"`
	Country                string `json:"country"`
	Nationality            string `json:"nationality"`
	RoleID                 string `json:"roleID"`
	Role                   string `json:"role"`
	IsTeamOwner            bool   `json:"isTeamOwner"`
	IsTeamStarter          bool   `json:"isTeamStarter"`
	TeamID                 string `json:"teamID"`
	TeamName               string `json:"teamName"`
	TeamNameFull           string `json:"teamNameFull"`
	TeamLogo               string `json:"teamLogo"`
	HonoursMention         any    `json:"honoursMention"`
	HonoursMentionLogo     any    `json:"honoursMentionLogo"`
	CooldownID             any    `json:"cooldownID"`
	CooldownNote           any    `json:"cooldownNote"`
	CooldownDateExpiresUTC any    `json:"cooldownDateExpiresUTC"`
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
	ID            string `json:"userID"`
	Name          string `json:"userName"`
	Logo          string `json:"userLogo"`
	Country       string `json:"country"`
	Nationality   string `json:"nationality"`
	DateJoinedUTC string `json:"dateJoinedUTC"`
	StreamURL     string `json:"streamUrl"`
	DiscordID     uint64 `json:"discordID"`
	DiscordTag    string `json:"discordTag"`
	SteamID       string `json:"steamID"`
	IsTerminated  bool   `json:"isTerminated"`
}

func (u UserCompact) GetDiscordID() string {
	return strconv.FormatUint(u.DiscordID, 10)
}

type PlayerSearchResult struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
