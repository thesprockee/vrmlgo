package vrmlgo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// All error constants
var (
	ErrJSONUnmarshal   = errors.New("json unmarshal")
	ErrPruneDaysBounds = errors.New("the number of days should be more than or equal to 1")
	ErrUnauthorized    = errors.New("HTTP request was unauthorized. This could be because the provided token was not a bot token. Please add \"Bot \" to the start of your token. https://discord.com/developers/docs/reference#authentication-example-bot-token-authorization-header")
)

var (
	// Marshal defines function used to encode JSON payloads
	Marshal func(v interface{}) ([]byte, error) = json.Marshal
	// Unmarshal defines function used to decode JSON payloads
	Unmarshal func(src []byte, v interface{}) error = json.Unmarshal
)

// RESTError stores error information about a request with a bad response code.
// Message is not always present, there are cases where api calls can fail
// without returning a json message.
type RESTError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte

	Message *APIErrorMessage // Message may be nil.
}

// newRestError returns a new REST API error.
func newRestError(req *http.Request, resp *http.Response, body []byte) *RESTError {
	restErr := &RESTError{
		Request:      req,
		Response:     resp,
		ResponseBody: body,
	}

	// Attempt to decode the error and assume no message was provided if it fails
	var msg *APIErrorMessage
	err := Unmarshal(body, &msg)
	if err == nil {
		restErr.Message = msg
	}

	return restErr
}

// Error returns a Rest API Error with its status code and body.
func (r RESTError) Error() string {
	return "HTTP " + r.Response.Status + ", " + string(r.ResponseBody)
}

// RateLimitError is returned when a request exceeds a rate limit
// and ShouldRetryOnRateLimit is false. The request may be manually
// retried after waiting the duration specified by RetryAfter.
type RateLimitError struct {
	*TooManyRequests
	URL string
}

// Error returns a rate limit error with rate limited endpoint and retry time.
func (e RateLimitError) Error() string {
	return "Rate limit exceeded on " + e.URL + ", retry after " + e.RetryAfter.String()
}

// RequestConfig is an HTTP request configuration.
type RequestConfig struct {
	Request                *http.Request
	ShouldRetryOnRateLimit bool
	ShouldIgnoreCacheFail  bool
	ShouldUseCache         bool
	MaxRestRetries         int
	Client                 *http.Client
}

// newRequestConfig returns a new HTTP request configuration based on parameters in Session.
func newRequestConfig(s *Session, req *http.Request) *RequestConfig {
	return &RequestConfig{
		ShouldRetryOnRateLimit: s.ShouldRetryOnRateLimit,
		ShouldUseCache:         s.CacheEnabled,
		MaxRestRetries:         s.MaxRestRetries,
		Client:                 s.Client,
		Request:                req,
	}
}

// RequestOption is a function which mutates request configuration.
// It can be supplied as an argument to any REST method.
type RequestOption func(cfg *RequestConfig)

// WithClient changes the HTTP client used for the request.
func WithClient(client *http.Client) RequestOption {
	return func(cfg *RequestConfig) {
		if client != nil {
			cfg.Client = client
		}
	}
}

// WithRetryOnRatelimit controls whether session will retry the request on rate limit.
func WithRetryOnRatelimit(retry bool) RequestOption {
	return func(cfg *RequestConfig) {
		cfg.ShouldRetryOnRateLimit = retry
	}
}

// WithRestRetries changes maximum amount of retries if request fails.
func WithRestRetries(max int) RequestOption {
	return func(cfg *RequestConfig) {
		cfg.MaxRestRetries = max
	}
}

// WithHeader sets a header in the request.
func WithHeader(key, value string) RequestOption {
	return func(cfg *RequestConfig) {
		cfg.Request.Header.Set(key, value)
	}
}

// WithContext changes context of the request.
func WithContext(ctx context.Context) RequestOption {
	return func(cfg *RequestConfig) {
		cfg.Request = cfg.Request.WithContext(ctx)
	}
}

// WithRetryOnRatelimit controls whether session will retry the request on rate limit.
func WithIgnoreCacheFailure(ignore bool) RequestOption {
	return func(cfg *RequestConfig) {
		cfg.ShouldIgnoreCacheFail = ignore
	}
}

func WithUseCache(use bool) RequestOption {
	return func(cfg *RequestConfig) {
		cfg.ShouldUseCache = use
	}
}

// Request is the same as RequestWithBucketID but the bucket id is the same as the urlStr
func (s *Session) Request(method, urlStr string, data interface{}, options ...RequestOption) (response []byte, err error) {
	return s.RequestWithBucketID(method, urlStr, data, strings.SplitN(urlStr, "?", 2)[0], options...)
}

// RequestWithBucketID makes a (GET/POST/...) Requests to Discord REST API with JSON data.
func (s *Session) RequestWithBucketID(method, urlStr string, data interface{}, bucketID string, options ...RequestOption) (response []byte, err error) {
	var body []byte
	if data != nil {
		body, err = Marshal(data)
		if err != nil {
			return
		}
	}

	return s.request(method, urlStr, "application/json", body, bucketID, 0, options...)
}

// request makes a (GET/POST/...) Requests to Discord REST API.
// Sequence is the sequence number, if it fails with a 502 it will
// retry with sequence+1 until it either succeeds or sequence >= session.MaxRestRetries
func (s *Session) request(method, urlStr, contentType string, b []byte, bucketID string, sequence int, options ...RequestOption) (response []byte, err error) {
	if bucketID == "" {
		bucketID = strings.SplitN(urlStr, "?", 2)[0]
	}
	return s.RequestWithLockedBucket(method, urlStr, contentType, b, s.Ratelimiter.LockBucket(bucketID), sequence, options...)
}

// RequestWithLockedBucket makes a request using a bucket that's already been locked
func (s *Session) RequestWithLockedBucket(method, urlStr, contentType string, b []byte, bucket *Bucket, sequence int, options ...RequestOption) (response []byte, err error) {
	if s.Debug {
		log.Printf("API REQUEST %8s :: %s\n", method, urlStr)
		log.Printf("API REQUEST  PAYLOAD :: [%s]\n", string(b))
	}

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(b))
	if err != nil {
		bucket.Release(nil)
		return
	}

	// Not used on initial login..
	// TODO: Verify if a login, otherwise complain about no-token
	if s.Token != "" {
		req.Header.Set("authorization", "Bearer "+s.Token)
	}

	// Discord's API returns a 400 Bad Request is Content-Type is set, but the
	// request body is empty.
	if b != nil {
		req.Header.Set("Content-Type", contentType)
	}

	// TODO: Make a configurable static variable.
	req.Header.Set("User-Agent", s.UserAgent)

	cfg := newRequestConfig(s, req)
	for _, opt := range options {
		opt(cfg)
	}
	req = cfg.Request

	if s.Debug {
		for k, v := range req.Header {
			log.Printf("API REQUEST   HEADER :: [%s] = %+v\n", k, v)
		}
	}

	if cfg.ShouldUseCache && method == "GET" {
		if data, found, err := s.Cache.Get(urlStr); err != nil {
			bucket.Release(nil)
			return nil, fmt.Errorf("error getting cached data: %w", err)
		} else if found {
			bucket.Release(nil)
			return []byte(data), nil
		}
	}

	resp, err := cfg.Client.Do(req)
	if err != nil {
		bucket.Release(nil)
		return
	}
	defer func() {
		err2 := resp.Body.Close()
		if s.Debug && err2 != nil {
			log.Println("error closing resp body")
		}
	}()

	err = bucket.Release(resp.Header)
	if err != nil {
		return
	}

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if s.Debug {

		log.Printf("API RESPONSE  STATUS :: %s\n", resp.Status)
		for k, v := range resp.Header {
			log.Printf("API RESPONSE  HEADER :: [%s] = %+v\n", k, v)
		}
		log.Printf("API RESPONSE    BODY :: [%s]\n\n\n", response)
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case http.StatusBadGateway:
		// Retry sending request if possible
		if sequence < cfg.MaxRestRetries {

			s.log(LogInformational, "%s Failed (%s), Retrying...", urlStr, resp.Status)
			response, err = s.RequestWithLockedBucket(method, urlStr, contentType, b, s.Ratelimiter.LockBucketObject(bucket), sequence+1, options...)
		} else {
			err = fmt.Errorf("Exceeded Max retries HTTP %s, %s", resp.Status, response)
		}
	case 429: // TOO MANY REQUESTS - Rate limiting
		rl := TooManyRequests{}
		err = Unmarshal(response, &rl)
		if err != nil {
			s.log(LogError, "rate limit unmarshal error, %s", err)
			return
		}

		if cfg.ShouldRetryOnRateLimit {
			s.log(LogInformational, "Rate Limiting %s, retry in %v", urlStr, rl.RetryAfter)

			time.Sleep(rl.RetryAfter)
			// we can make the above smarter
			// this method can cause longer delays than required

			response, err = s.RequestWithLockedBucket(method, urlStr, contentType, b, s.Ratelimiter.LockBucketObject(bucket), sequence, options...)
		} else {
			err = &RateLimitError{TooManyRequests: &rl, URL: urlStr}
		}
	case http.StatusUnauthorized:
		fallthrough
	default: // Error condition
		err = newRestError(req, resp, response)
	}

	if cfg.ShouldUseCache && s.Cache != nil {
		if method == "GET" {
			if err := s.Cache.Set(urlStr, string(response)); err != nil && !cfg.ShouldIgnoreCacheFail {
				return nil, fmt.Errorf("error setting cached data: %w", err)
			}
		}
	}
	return
}

func unmarshal(data []byte, v interface{}) error {
	err := Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}

	return nil
}

func (s *Session) Me(options ...RequestOption) (st *User, err error) {
	// Do not cache @Me requests
	options = append(options, WithUseCache(false))

	body, err := s.RequestWithBucketID("GET", EndpointMe, nil, EndpointAPI, options...)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

func (s *Session) Member(userID string, options ...RequestOption) (st *Member, err error) {

	body, err := s.RequestWithBucketID("GET", EndpointMember(userID), nil, EndpointAPI)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

func (s *Session) GameSearch(gameName string, options ...RequestOption) (st *GameDetails, err error) {

	body, err := s.RequestWithBucketID("GET", EndpointGame(gameName), nil, EndpointAPI)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

func (s *Session) Seasons(gameName string, options ...RequestOption) (st []*Season, err error) {
	body, err := s.RequestWithBucketID("GET", EndpointSeasons(gameName), nil, EndpointAPI)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

func (s *Session) TeamMatchesHistory(teamID string, options ...RequestOption) (st []*MatchHistory, err error) {

	body, err := s.RequestWithBucketID("GET", EndpointTeamMatchesHistory(teamID), nil, EndpointAPI)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

func (s *Session) Match(gName, matchID string, options ...RequestOption) (st *MatchDetails, err error) {

	body, err := s.RequestWithBucketID("GET", EndpointMatch(gName, matchID), nil, EndpointAPI)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

func (s *Session) Team(teamID string, options ...RequestOption) (st *TeamDetails, err error) {

	body, err := s.RequestWithBucketID("GET", EndpointTeam(teamID), nil, EndpointAPI)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

func (s *Session) Player(playerID string, options ...RequestOption) (st *Player, err error) {

	body, err := s.RequestWithBucketID("GET", EndpointPlayer(playerID), nil, EndpointAPI)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}
