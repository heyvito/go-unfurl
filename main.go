package unfurl

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// Client represents a unfurl client instance
type Client struct {
	httpClient *http.Client
	options    Options
}

// Options exposes internal settings to change the Client behaviour
type Options struct {
	// MaxHops defines how many redirects the client can suffer before returning
	// an error
	MaxHops int
}

// ErrTooManyRedirects indicates that the unfurl client has archieved the
// maximum allowed request count defined by the Options struct
var ErrTooManyRedirects = errors.New("Too many redirects")

// NewClientWithOptions returns a new instance of a Client using the provided
// Options values
func NewClientWithOptions(opts Options) Client {
	return Client{
		options: opts,
		httpClient: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// NewClient returns a new instance of a Client using the default values for
// the settings parameters.
// (MaxHops: 20)
func NewClient() Client {
	return NewClientWithOptions(Options{MaxHops: 20})
}

// Process attempts to follow all possible redirects of a given URL
func (c *Client) Process(in string) (string, error) {
	jar, _ := cookiejar.New(nil)
	c.httpClient.Jar = jar
	hops := 0
	for {
		resp, err := c.httpClient.Get(in)
		if err != nil {
			return "", err
		}
		if _, ok := resp.Header["Location"]; ok && resp.StatusCode/100.0 == 3 {
			if hops >= c.options.MaxHops {
				return "", ErrTooManyRedirects
			}
			hops++
			base, err := url.Parse(in)
			if err != nil {
				return "", err
			}
			next, err := url.Parse(resp.Header["Location"][0])
			if err != nil {
				return "", err
			}
			in = base.ResolveReference(next).String()
			continue
		} else {
			return resp.Request.URL.String(), nil
		}
	}
}
