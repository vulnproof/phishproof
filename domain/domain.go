package domain

import (
	"net/url"
	"strings"
)

// Domain represents a domain name
type Domain struct {
	Name string
}

// NewDomain creates a new Domain from a string
func NewDomain(domain string) (*Domain, error) {
	u, err := url.Parse(domain)
	if err == nil && u.Scheme != "" {
		domain = u.Hostname()
	}

	domain = strings.ToLower(domain)

	if strings.HasPrefix(domain, "www.") {
		domain = domain[4:]
	}

	return &Domain{
		Name: domain,
	}, nil
}
