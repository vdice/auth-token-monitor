package providers

import "net/url"

// The Auto provider automatically determines which provider to use
// for each token.
var Auto = Provider{
	Name:    "auto",
	BaseURL: &url.URL{},
}
