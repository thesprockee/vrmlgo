package vrmlgo

type TeamDetails struct {
	Team            *Team    `json:"team"`
	Season          *Season  `json:"season"`
	SeasonTeam      any      `json:"seasonTeam"`
	SeasonPlayers   any      `json:"seasonPlayers"`
	SeasonStatsMaps []any    `json:"seasonStatsMaps"`
	SeasonMatches   []any    `json:"seasonMatches"`
	ExMembers       []*User  `json:"exMembers"`
	Context         *Context `json:"context"`
}

func (t TeamDetails) ID() string {
	return t.Team.ID
}

type Context struct {
	UserTeamID          string `json:"userTeamID"`
	UserTeamIsTeamOwner bool   `json:"userTeamIsTeamOwner"`
}

type Team struct {
	ID                                       string    `json:"teamID"`
	Regions                                  []*Region `json:"regions"`
	CurrentUserPlayerID                      string    `json:"currentUserPlayerID"`
	CurrentUserIsTeamOwner                   bool      `json:"currentUserIsTeamOwner"`
	CurrentUserIsTeamOwnerCaptainOrCoCaptain bool      `json:"currentUserIsTeamOwnerCaptainOrCoCaptain"`
	CurrentUserIsTeamStarter                 bool      `json:"currentUserIsTeamStarter"`
	CurrentUserHasScrimNotifications         bool      `json:"currentUserHasScrimNotifications"`
	TournamentRegistrationFlag               int64     `json:"tournamentRegistrationFlag"`
	TournamentRegistered                     int64     `json:"tournamentRegistered"`
	GameName                                 string    `json:"gameName"`
	GameURL                                  string    `json:"gameUrl"`
	LogoUpdate                               string    `json:"logoUpdate"`
	TeamMode                                 string    `json:"teamMode"`
	IsActive                                 bool      `json:"isActive"`
	IsRetired                                bool      `json:"isRetired"`
	IsDeleted                                bool      `json:"isDeleted"`
	IsRecruiting                             bool      `json:"isRecruiting"`
	IsBlockingRecruiting                     bool      `json:"isBlockingRecruiting"`
	ScheduleInfoWeekdaysFromUTC              string    `json:"scheduleInfoWeekdaysFromUTC"`
	ScheduleInfoWeekdaysToUTC                string    `json:"scheduleInfoWeekdaysToUTC"`
	ScheduleInfoWeekendsFromUTC              string    `json:"scheduleInfoWeekendsFromUTC"`
	ScheduleInfoWeekendsToUTC                string    `json:"scheduleInfoWeekendsToUTC"`
	Players                                  []*Player `json:"players"`
	TeamName                                 string    `json:"teamName"`
	TeamLogo                                 string    `json:"teamLogo"`
	RegionID                                 string    `json:"regionID"`
	RegionName                               string    `json:"regionName"`
}
