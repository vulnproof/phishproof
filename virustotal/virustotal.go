package virustotal

import (
	"github.com/VirusTotal/vt-go"
)

func ScanURLWithVT(url string, apiKey string) ([]byte, error) {
	// Set up the VirusTotal API v3 client.
	client := vt.NewClient(apiKey)

	// Create a URL scanner.
	scanner := client.NewURLScanner()

	// Scan the URL.
	res, err := scanner.Scan(url)
	if err != nil {
		return nil, err
	}

	return res.MarshalJSON()
}
