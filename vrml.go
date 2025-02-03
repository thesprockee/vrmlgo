package vrmlgo

import (
	"net/http"
	"time"
)

// This is modeled after github.com/bwmarrin/discordgo/discord.go

const VERSION = "1.0.2"

func New(token string) (s *Session) {

	// Create an empty Session interface.
	return &Session{
		Token:                  token,
		Ratelimiter:            NewRatelimiter(),
		Cache:                  NewLocalCache(),
		ShouldRetryOnRateLimit: true,
		MaxRestRetries:         3,
		Client:                 &http.Client{Timeout: (20 * time.Second)},
		UserAgent:              "vrmlgo (https://github.com/echotools/vrmlgo, v" + VERSION + ")",
	}
}
