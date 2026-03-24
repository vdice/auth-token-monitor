package providers

import (
	"net/url"
	"regexp"
)

var Github = Provider{
	Name:       "github",
	AuthHeader: "github-authentication-token-expiration",
	BaseURL: &url.URL{
		Scheme: "https",
		Host:   "api.github.com",
	},
	ExpectedStatusCode: 200,
	TokenPatterns: []*regexp.Regexp{
		// classic
		regexp.MustCompile(`^ghp_[a-zA-Z0-9]{36}$`),
		// fine-grained
		regexp.MustCompile(`^github_pat_[a-zA-Z0-9]{22}_[a-zA-Z0-9]{59}$`),
	},
}
