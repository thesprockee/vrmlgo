package vrmlgo

import (
	"math"
	"net/http"
	"sync"
	"time"
)

// A Session represents a connection to the Discord API.
type Session struct {
	sync.RWMutex

	// General configurable settings.

	// Authentication token for this session
	Token string

	// Debug for printing JSON request/responses
	LogLevel int
	Debug    bool

	// Should the session retry requests when rate limited.
	ShouldRetryOnRateLimit bool

	// Should use cache
	CacheEnabled bool

	// Max number of REST API retries
	MaxRestRetries int

	// The http client used for REST requests
	Client *http.Client

	// The user agent used for REST APIs
	UserAgent string

	// Used to deal with rate limits
	Ratelimiter *RateLimiter

	// Used to cache data
	Cache Cache
}

// A TooManyRequests struct holds information received from VRML
// when receiving a HTTP 429 response.
type TooManyRequests struct {
	Bucket     string        `json:"bucket"`
	Message    string        `json:"message"`
	RetryAfter time.Duration `json:"retry_after"`
}

// UnmarshalJSON helps support translation of a milliseconds-based float
// into a time.Duration on TooManyRequests.
func (t *TooManyRequests) UnmarshalJSON(b []byte) error {
	u := struct {
		Bucket     string  `json:"bucket"`
		Message    string  `json:"message"`
		RetryAfter float64 `json:"retry_after"`
	}{}
	err := Unmarshal(b, &u)
	if err != nil {
		return err
	}

	t.Bucket = u.Bucket
	t.Message = u.Message
	whole, frac := math.Modf(u.RetryAfter)
	t.RetryAfter = time.Duration(whole)*time.Second + time.Duration(frac*1000)*time.Millisecond
	return nil
}

// An APIErrorMessage is an api error message returned from VRML
type APIErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// A UserToken is the token information returned from VRML
type UserToken struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}
