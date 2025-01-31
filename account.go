package vrmlgo

import "slices"

type Member struct {
	Metadata                      User        `json:"user"`
	Games                         []UserGames `json:"allGames"`
	AllPlayersSameDiscord         interface{} `json:"allPlayersSameDiscord"`
	AllPlayersSameIP              interface{} `json:"allPlayersSameIP"`
	PenaltyPoints                 interface{} `json:"penaltyPoints"`
	BanInfo                       interface{} `json:"banInfo"`
	LocationFromIP                interface{} `json:"locationFromIP"`
	CurrentUserIsModInRelatedGame bool        `json:"currentUserIsModInRelatedGame"`
}

func (a *Member) PlayerID(gameUrlShort string) string {
	for _, g := range a.Games {
		if g.Game.ShortName == gameUrlShort {
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
	}
	return ""
}

// Teams returns a list of team IDs that the user is a member of for a given game.
func (a *Member) Teams(urlShort string) []string {
	teamIDs := make([]string, 0)

	for _, g := range a.Games {
		if g.Game.ShortName == urlShort {

			teamIDs = append(teamIDs, g.BioCurrent.TeamID)

			for _, t := range g.BioCurrentSeasonPastTeams {
				teamIDs = append(teamIDs, t.TeamID)
			}

			for _, t := range g.BioPastSeasons {
				teamIDs = append(teamIDs, t.TeamID)
			}
		}
	}

	for i := 0; i < len(teamIDs); i++ {
		if teamIDs[i] == "" {
			teamIDs = append(teamIDs[:i], teamIDs[i+1:]...)
			i--
		}
	}

	slices.Sort(teamIDs)
	teamIDs = slices.Compact(teamIDs)

	return teamIDs
}
