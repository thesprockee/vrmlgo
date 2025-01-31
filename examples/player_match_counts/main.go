package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/echotools/vrmlgo"
)

func main() {

	// Get the API token from environment or .env file
	token := os.Getenv("VRML_API_TOKEN")
	if token == "" {
		token = read_token_from_dotenv()
	}
	if token == "" {
		fmt.Println("Please set the VRML_API_TOKEN environment variable or create a .env file with the token.")
		os.Exit(1)
	}

	// Create a new VRML API session
	vg := vrmlgo.New(token)

	// Get the user's account information
	me, err := vg.Me()
	if err != nil {
		panic(err)
	}

	// Get the account information
	account, err := vg.Member(me.ID, vrmlgo.WithUseCache(false))
	if err != nil {
		panic(err)
	}

	// Get the game details
	gameDetails, err := vg.GameSearch("EchoArena")
	if err != nil {
		panic(err)
	}

	// Get the seasons for the game
	seasons, err := vg.Seasons(gameDetails.Game.ShortName)
	if err != nil {
		panic(err)
	}

	pp(seasons)
	// Create a map of seasons
	seasonMap := make(map[string]*vrmlgo.Season)
	for _, s := range seasons {
		seasonMap[s.ID] = s
	}

	// Get the player ID for this game
	playerID := account.PlayerID(gameDetails.Game.ShortName)

	// Get the match history for each team
	matchesBySeason := make(map[string][]string)
	for _, t := range account.Teams(gameDetails.Game.ShortName) {

		history, err := vg.TeamMatchesHistory(t)
		if err != nil {
			panic(err)
		}

		// Create a map of matches by season
		for _, h := range history {
			matchesBySeason[h.SeasonName] = append(matchesBySeason[h.SeasonName], h.MatchID)
		}
	}

	// Get the match details for the first two matches of each season
	matchCountBySeason := make(map[string]int)

	for _, season := range seasons {

		for _, mID := range matchesBySeason[season.Name] {

			// Get the match details
			matchDetails, err := vg.Match(gameDetails.Game.ShortName, mID)
			if err != nil {
				panic(err)
			}

			// Skip forfeits
			if matchDetails.Match.IsForfeit {
				continue
			}

			// Count the number of matches the player is in
			for _, p := range matchDetails.Players() {

				// Check if the player is in the match
				if p.ID == playerID {
					matchCountBySeason[matchDetails.Match.SeasonName]++
				}
			}
		}
	}

	// Get the season names
	seasonNames := make([]string, 0, len(matchCountBySeason))
	for sName := range matchCountBySeason {
		seasonNames = append(seasonNames, sName)
	}

	// Sort the season names
	slices.Sort(seasonNames)

	// Print the number of matches, by season, the player was part of the team
	for _, sName := range seasonNames {
		fmt.Printf("Season: %s, Matches: %d\n", sName, matchCountBySeason[sName])
	}

}

func read_token_from_dotenv() string {

	// Recursively look for a .env file in parent directories
	// Start at the current directory
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	// Keep looking in parent directories until we find a .env file or hit the root
	for {
		if _, err := os.Stat(dir + "/.env"); err == nil {
			// Found .env file, change to this directory
			os.Chdir(dir)
			break
		}

		// Move up one directory
		parent := dir[:strings.LastIndex(dir, "/")]
		if parent == dir {
			// We've hit the root directory
			return ""
		}
		dir = parent
	}

	dotenvfile, err := os.Open(dir + "/.env")
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(dotenvfile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "VRML_API_TOKEN=") {
			s := strings.TrimPrefix(line, "VRML_API_TOKEN=")
			s = strings.Trim(s, "\"")
			return s
		}
	}
	return ""
}

func pp(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
	fmt.Println()
}
