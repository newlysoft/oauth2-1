// +build !appengine

package google

import (
	"strings"

	"github.com/golang/oauth2"
	"google.golang.org/appengine"
)

// AppEngineConfig represents a configuration for an
// App Engine application's Google service account.
type AppEngineConfig struct {
	context appengine.Context
	scopes  []string
	cache   oauth2.Cache
}

// NewAppEngineConfig creates a new AppEngineConfig for the
// provided auth scopes.
func NewAppEngineConfig(context appengine.Context, scopes []string) *AppEngineConfig {
	return &AppEngineConfig{context: context, scopes: scopes}
}

// NewTransport returns a transport that authorizes
// the requests with the application's service account.
func (c *AppEngineConfig) NewTransport() oauth2.Transport {
	return oauth2.NewAuthorizedTransport(c, nil)
}

// NewTransport returns a token-caching transport that authorizes
// the requests with the application's service account.
func (c *AppEngineConfig) NewTransportWithCache(cache oauth2.Cache) (oauth2.Transport, error) {
	token, err := cache.Read()
	if err != nil {
		return nil, err
	}
	c.cache = cache
	return oauth2.NewAuthorizedTransport(c, token), nil
}

// FetchToken fetches a new access token for the provided scopes.
func (c *AppEngineConfig) FetchToken(existing *oauth2.Token) (*oauth2.Token, error) {
	token, expiry, err := appengine.AccessToken(c.context, strings.Join(c.scopes, " "))
	if err != nil {
		return nil, err
	}
	return &oauth2.Token{
		AccessToken: token,
		Expiry:      expiry,
	}, nil
}

func (c *AppEngineConfig) Cache() oauth2.Cache {
	return c.cache
}