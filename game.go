package vrmlgo

type GameDetails struct {
	Game        Game          `json:"game"`
	NewsPosts   []NewsPost    `json:"newsPosts"`
	NextMatches []interface{} `json:"nextMatches"`
	Season      GameSeason    `json:"season"`
}

type Game struct {
	ID              string      `json:"gameID"`
	RoundsPerMap    int64       `json:"roundsPerMap"`
	GameByURL       string      `json:"gameByUrl"`
	GameByImage     string      `json:"gameByImage"`
	HeaderImage     string      `json:"headerImage"`
	HomeTitleImage  interface{} `json:"homeTitleImage"`
	DiscordInvite   string      `json:"discordInvite"`
	DiscordServerID string      `json:"discordServerID"`
	Youtube         string      `json:"youtube"`
	Twitter         string      `json:"twitter"`
	Reddit          string      `json:"reddit"`
	Facebook        string      `json:"facebook"`
	Name            string      `json:"gameName"`
	TeamMode        string      `json:"teamMode"`
	MatchMode       string      `json:"matchMode"`
	URL             string      `json:"url"`
	ShortName       string      `json:"urlShort"`
	URLComplete     string      `json:"urlComplete"`
	HasSubstitutes  bool        `json:"hasSubstitutes"`
	HasTies         bool        `json:"hasTies"`
	HasCasters      bool        `json:"hasCasters"`
	HasCameraman    bool        `json:"hasCameraman"`
}

type NewsPost struct {
	ID               string  `json:"newsID"`
	UserID           string  `json:"userID"`
	UserLogo         string  `json:"userLogo"`
	UserName         string  `json:"userName"`
	DateSubmittedUTC string  `json:"dateSubmittedUTC"`
	DateEditedUTC    *string `json:"dateEditedUTC"`
	Title            string  `json:"title"`
	HTML             string  `json:"html"`
}

type GameSeason struct {
	ID                       string      `json:"seasonID"`
	DateStartUTC             string      `json:"dateStartUTC"`
	DateEndUTC               string      `json:"dateEndUTC"`
	DateRosterLockUTC        string      `json:"dateRosterLockUTC"`
	DateChampionshipStartUTC interface{} `json:"dateChampionshipStartUTC"`
	ChampionshipURL          string      `json:"championshipUrl"`
	GameName                 string      `json:"gameName"`
	GameURLShort             string      `json:"gameUrlShort"`
	GameActive               bool        `json:"gameActive"`
	SeasonName               string      `json:"seasonName"`
	IsCurrent                bool        `json:"isCurrent"`
}
