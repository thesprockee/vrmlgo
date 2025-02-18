package vrmlgo

// Known Discord API Endpoints.
var (
	EndpointAPI     = "https://api.vrmasterleague.com/"
	EndpointToken   = EndpointAPI + "Users/Token"
	EndpointUsers   = EndpointAPI + "Users/"
	EndpointTeams   = EndpointAPI + "Teams/"
	EndpointPlayers = EndpointAPI + "Players/"

	EndpointMe                 = EndpointAPI + "Users/@Me"
	EndpointMember             = func(uID string) string { return EndpointUsers + uID }
	EndpointPlayer             = func(pID string) string { return EndpointPlayers + pID }
	EndpointGame               = func(gShortName string) string { return EndpointAPI + gShortName }
	EndpointTeam               = func(gID string) string { return EndpointTeams + gID }
	EndpointTeamMatchesHistory = func(tID string) string { return EndpointTeams + tID + "/Matches/History/Detailed" }
	EndpointMatch              = func(gID string, mID string) string { return EndpointAPI + gID + "/Matches/" + mID }
	EndpointSeasons            = func(gID string) string { return EndpointAPI + gID + "/Seasons" }
)
