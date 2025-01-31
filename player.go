package vrmlgo

type GameParticipant struct {
	User                User          `json:"user"`
	ThisGame            UserGames     `json:"thisGame"`
	ConnoisseurOverview interface{}   `json:"connoisseurOverview"`
	ConnoisseurHistory  []interface{} `json:"connoisseurHistory"`
}

type Player struct {
	ID                            string      `json:"playerID"`
	IsCooldown                    bool        `json:"isCooldown"`
	CooldownNote                  interface{} `json:"cooldownNote"`
	CooldownDateExpiresUTC        interface{} `json:"cooldownDateExpiresUTC"`
	HonoursMentionNote            interface{} `json:"honoursMentionNote"`
	HonoursMentionLogo            interface{} `json:"honoursMentionLogo"`
	DiscordID                     float64     `json:"discordID"`
	DiscordTag                    string      `json:"discordTag"`
	DiscordTeamRole               *int64      `json:"discordTeamRole"`
	PlayerName                    string      `json:"playerName"`
	UserID                        string      `json:"userID"`
	UserLogo                      string      `json:"userLogo"`
	Country                       string      `json:"country"`
	Nationality                   string      `json:"nationality"`
	StreamURL                     *string     `json:"streamURL"`
	TeamID                        string      `json:"teamID"`
	TeamName                      string      `json:"teamName"`
	RoleID                        string      `json:"roleID"`
	Role                          string      `json:"role"`
	IsTeamOwner                   bool        `json:"isTeamOwner"`
	IsTeamOwnerCaptainOrCoCaptain bool        `json:"isTeamOwnerCaptainOrCoCaptain"`
	IsTeamStarter                 bool        `json:"isTeamStarter"`
}
