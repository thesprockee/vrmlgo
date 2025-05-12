package vrmlgo

import (
	"slices"
	"strconv"
	"time"
)

type User struct {
	ID               string `json:"userID"`
	Email            string `json:"email,omitempty"`
	Name             string `json:"userName"`
	Logo             string `json:"userLogo"`
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

func (u User) GetDateJoined() time.Time {
	dateJoined, _ := time.Parse("2006-01-02 15:04", u.DateJoinedUTC)
	return dateJoined
}

type UserGame struct {
	PlayerID                  string                            `json:"playerID"`
	PlayerName                string                            `json:"playerName"`
	UserLogo                  string                            `json:"userLogo"`
	Game                      Game                              `json:"game"`
	BioCurrent                BioCurrent                        `json:"bioCurrent"`
	BioCurrentDef             BioCurrent                        `json:"bioCurrentDef"`
	BioCurrentSeasonPastTeams []BioCurrentSeasonPastTeamElement `json:"bioCurrentSeasonPastTeams"`
	BioPastSeasons            []BioCurrentSeasonPastTeamElement `json:"bioPastSeasons"`
}

// Teams returns a list of team IDs that the user is a member of for a given game.
func (g *UserGame) TeamIDs() []string {
	teamIDs := make([]string, 0)

	teamIDs = append(teamIDs, g.BioCurrent.TeamID)

	for _, t := range g.BioCurrentSeasonPastTeams {
		teamIDs = append(teamIDs, t.TeamID)
	}

	for _, t := range g.BioPastSeasons {
		teamIDs = append(teamIDs, t.TeamID)
	}
	for i := 0; i < len(teamIDs); i++ {
		if teamIDs[i] == "" {
			teamIDs = slices.Delete(teamIDs, i, i+1)
			i--
		}
	}

	slices.Sort(teamIDs)
	teamIDs = slices.Compact(teamIDs)

	return teamIDs
}

type UserTeams struct {
	Teams           []Team `json:"teams"`
	Substitutes     []any  `json:"substitutes"`
	PendingRecruits []any  `json:"pendingRecruits"`
}

type Region struct {
	ID         string `json:"regionID"`
	RegionName string `json:"regionName"`
	RegionLogo string `json:"regionLogo"`
}

type BioCurrent struct {
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

type BioCurrentSeasonPastTeamElement struct {
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
