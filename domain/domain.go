package domain

import (
	"net/url"
)

// Domain represents a domain name
type Domain struct {
	Name string
}

// GetDomainFromURL extracts and returns the domain from a given URL string.
// It returns an empty string if the URL is invalid or cannot be parsed.
func GetDomainFromURL(u string) (*Domain, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return &Domain{
			Name: "",
		}, nil
	}

	return &Domain{
		Name: parsed.Hostname(),
	}, nil
}
