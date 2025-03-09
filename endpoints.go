package vrmlgo

// Known Discord API Endpoints.
var (
	EndpointAPI     = "https://api.vrmasterleague.com/"
	EndpointToken   = EndpointAPI + "Users/Token"
	EndpointPlayers = EndpointAPI + "Players/"
	EndpointTeams   = EndpointAPI + "Teams/"
	EndpointUsers   = EndpointAPI + "Users/"
	EndpointMe      = EndpointAPI + "Users/@Me"

	EndpointMember             = func(uID string) string { return EndpointUsers + uID }
	EndpointPlayer             = func(pID string) string { return EndpointPlayers + pID }
	EndpointTeam               = func(tID string) string { return EndpointTeams + tID }
	EndpointTeamMatchesHistory = func(tID string) string { return EndpointTeams + tID + "/Matches/History/Detailed" }

	EndpointGame              = func(gName string) string { return EndpointAPI + gName }
	EndpointGameMatch         = func(gName string, mID string) string { return EndpointAPI + gName + "/Matches/" + mID }
	EndpointGameSeasons       = func(gName string) string { return EndpointAPI + gName + "/Seasons" }
	EndpointGamePlayersSearch = func(gName string) string { return EndpointAPI + gName + "/Players/Search" }
)
