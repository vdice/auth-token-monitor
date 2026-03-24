package providers

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/alecthomas/kong"
)

func (p *Provider) Decode(ctx *kong.DecodeContext) error {
	value := ctx.Scan.Pop().String()
	provider, ok := Providers[value]
	if !ok {
		return fmt.Errorf("unsupported provider: %q", value)
	}
	*p = provider
	return nil
}

func (p *Provider) Equal(p2 *Provider) bool {
	if p == p2 {
		return true
	}
	if p == nil || p2 == nil {
		return false
	}

	if p.Name != p2.Name ||
		p.AuthHeader != p2.AuthHeader ||
		p.BaseURL != p2.BaseURL ||
		p.Path != p2.Path ||
		p.ExpectedStatusCode != p2.ExpectedStatusCode {
		return false
	}

	// Compare the TokenPatterns slices manually
	if len(p.TokenPatterns) != len(p2.TokenPatterns) {
		return false
	}
	for i := range p.TokenPatterns {
		if p.TokenPatterns[i].String() != p2.TokenPatterns[i].String() {
			return false
		}
	}

	return true
}

type Provider struct {
	Name               string
	AuthHeader         string
	BaseURL            *url.URL
	Path               string
	ExpectedStatusCode int
	TokenPatterns      []*regexp.Regexp
}

var Providers = map[string]Provider{
	"github": Github,
	"fwf":    Fwf,
	"linode": Linode,
	"auto":   Auto,
}
