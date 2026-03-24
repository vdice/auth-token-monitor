package providers

import (
	"net/url"
	"regexp"
)

var Linode = Provider{
	Name: "linode",
	BaseURL: &url.URL{
		Scheme: "https",
		Host:   "api.linode.com",
	},
	TokenPatterns: []*regexp.Regexp{
		regexp.MustCompile(`^[a-f0-9]{64}$`),
	},
}
