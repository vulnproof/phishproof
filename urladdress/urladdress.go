package urladdress

import (
	"net/http"
	"net/url"
)

// Check if a given string is a valid URL
func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// Extract the domain from a given URL string
func extractDomain(u string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return parsed.Hostname(), nil
}

// Check if a given URL is redirecting to another URL
// If it is redirecting, return the redirected URL
func checkURLRedirect(url string) (string, bool) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", false
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	if resp.Request.URL.String() != url {
		return resp.Request.URL.String(), true
	}

	return url, false
}

// Check if a given URL is a shortened URL
func isShortenedURL(urlToCheck string) bool {
	shortenedHosts := []string{
		"bit.ly",
		"goo.gl",
		"t.co",
		"tinyurl.com",
		"ow.ly",
		"is.gd",
		"buff.ly",
		"adf.ly",
		"j.mp",
		"amzn.to",
		"fb.me",
		"mzl.la",
		"n.pr",
		"nyti.ms",
		"tcrn.ch",
	}

	u, err := url.Parse(urlToCheck)
	if err != nil {
		return false
	}

	for _, host := range shortenedHosts {
		if u.Host == host {
			return true
		}
	}
	return false
}

// Check if a given URL is safe using Google Safe Browsing API
func isURLSafeUsingGoogle(url string, apiKey string) (bool, error) {
	// TODO: Implement the logic to check URL safety using the Google Safe Browsing API
	return false, nil
}

// Check if a given URL is safe using VirusTotal API
func isURLSafeUsingVirusTotal(url string, apiKey string) (bool, error) {
	// TODO: Implement the logic to check URL safety using the VirusTotal API
	return false, nil
}
