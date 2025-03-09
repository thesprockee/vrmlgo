package vrmlgo

import "slices"

type Member struct {
	User                          *User       `json:"user"`
	Games                         []UserGame  `json:"allGames"`
	AllPlayersSameDiscord         interface{} `json:"allPlayersSameDiscord"`
	AllPlayersSameIP              interface{} `json:"allPlayersSameIP"`
	PenaltyPoints                 interface{} `json:"penaltyPoints"`
	BanInfo                       interface{} `json:"banInfo"`
	LocationFromIP                interface{} `json:"locationFromIP"`
	CurrentUserIsModInRelatedGame bool        `json:"currentUserIsModInRelatedGame"`
}

// PlayerID returns the player ID of the user for a given game.
func (a *Member) PlayerID(gameName string) string {
	for _, g := range a.Games {

		if g.Game.ShortName != gameName {
			continue
		}

		if g.BioCurrent.PlayerID != "" {
			return g.BioCurrent.PlayerID
		}

		for _, t := range g.BioCurrentSeasonPastTeams {
			if t.PlayerID != "" {
				return t.PlayerID
			}
		}

		for _, t := range g.BioPastSeasons {
			if t.PlayerID != "" {
				return t.PlayerID
			}
		}

	}

	return ""
}

// Teams returns a list of team IDs that the user is a member of for a given game.
func (a *Member) TeamIDs(gameName string) []string {
	teamIDs := make([]string, 0)

	for _, g := range a.Games {
		if g.Game.ShortName == gameName {
			teamIDs = append(teamIDs, g.TeamIDs()...)
		}
	}

	slices.Sort(teamIDs)
	teamIDs = slices.Compact(teamIDs)

	return teamIDs
}
