package safebrowsing

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type SafeBrowsingClient struct {
	APIKey string
}

type SafeBrowsingRequest struct {
	Client     SafeBrowsingClient `json:"-"`
	ThreatInfo ThreatInfo         `json:"threatInfo"`
}

type ThreatInfo struct {
	ThreatTypes      []string `json:"threatTypes"`
	PlatformTypes    []string `json:"platformTypes"`
	ThreatEntryTypes []string `json:"threatEntryTypes"`
	ThreatEntries    []struct {
		URL string `json:"url"`
	} `json:"threatEntries"`
}

type SafeBrowsingResponse struct {
	Matches []struct {
		ThreatType          string   `json:"threatType"`
		PlatformType        string   `json:"platformType"`
		ThreatEntryType     string   `json:"threatEntryType"`
		Threat              Threat   `json:"threat"`
		ThreatEntryMetadata Metadata `json:"threatEntryMetadata"`
		CacheDuration       string   `json:"cacheDuration"`
	} `json:"matches"`
}

type Threat struct {
	URL string `json:"url"`
}

type Metadata struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewSafeBrowsingClient() *SafeBrowsingClient {
	return &SafeBrowsingClient{
		APIKey: viper.GetString("safebrowsing.api_key"),
	}
}

func (s *SafeBrowsingClient) CheckURL(urlStr string) (bool, error) {
	req := SafeBrowsingRequest{
		Client: SafeBrowsingClient{
			APIKey: s.APIKey,
		},
		ThreatInfo: ThreatInfo{
			ThreatTypes:      []string{"MALWARE", "SOCIAL_ENGINEERING", "UNWANTED_SOFTWARE", "POTENTIALLY_HARMFUL_APPLICATION"},
			PlatformTypes:    []string{"ANY_PLATFORM"},
			ThreatEntryTypes: []string{"URL"},
			ThreatEntries: []struct {
				URL string `json:"url"`
			}{
				{URL: urlStr},
			},
		},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return false, err
	}

	apiURL := "https://safebrowsing.googleapis.com/v4/threatMatches:find?key=" + s.APIKey

	resp, err := http.Post(apiURL, "application/json", strings.NewReader(string(reqBody)))
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	var response SafeBrowsingResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return false, err
	}

	if len(response.Matches) > 0 {
		return true, nil
	}

	return false, nil
}
