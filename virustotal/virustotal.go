package virustotal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type virusTotalResponse struct {
	ResponseCode int                    `json:"response_code"`
	VerboseMsg   string                 `json:"verbose_msg"`
	ScanResults  map[string]interface{} `json:"scans"`
}

func checkVirusTotal(url string) (bool, error) {
	vtApiKey := viper.GetString("virustotal.api_key")
	vtApiUrl := fmt.Sprintf("https://www.virustotal.com/vtapi/v2/url/report?apikey=%s&resource=%s", vtApiKey, url)

	resp, err := http.Get(vtApiUrl)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var vtResponse virusTotalResponse
	err = json.NewDecoder(resp.Body).Decode(&vtResponse)
	if err != nil {
		return false, err
	}

	if vtResponse.ResponseCode != 1 {
		return false, fmt.Errorf("no results found for URL %s", url)
	}

	for _, scanResult := range vtResponse.ScanResults {
		result := scanResult.(map[string]interface{})
		if result["detected"].(bool) {
			return true, nil
		}
	}

	return false, nil
}
