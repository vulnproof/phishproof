package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/vulnproof/phishproof/domain"
	"github.com/vulnproof/phishproof/urladdress"
	"github.com/vulnproof/phishproof/virustotal"
)

func main() {
	// Load the configuration file
	viper.SetConfigFile("config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Get the API key for VirusTotal
	vtAPIKey := viper.GetString("virustotal.api_key")

	// Create a new VirusTotal client
	vtClient, err := virustotal.NewClient(vtAPIKey)
	if err != nil {
		log.Fatalf("Failed to create VirusTotal client: %v", err)
	}

	// Create a new URL scanner
	urlScanner := url.NewScanner()

	// Create a new domain checker
	domainChecker := domain.NewChecker()

	// Get the URL to scan from the command line arguments
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Please provide a URL to scan")
	}
	urlToScan := args[0]

	// Check if the URL is valid
	if !url.IsValidURL(urlToScan) {
		log.Fatal("Invalid URL")
	}

	// Extract the domain from the URL
	domainName := url.ExtractDomain(urlToScan)

	// Check the domain against the list of known phishing domains
	if domainChecker.IsPhishingDomain(domainName) {
		log.Printf("Phishing domain: %s", domainName)
		return
	}

	// Check the URL for redirects
	redirectURL, err := urlScanner.CheckForRedirects(urlToScan)
	if err != nil {
		log.Fatalf("Error checking URL for redirects: %v", err)
	}
	if redirectURL != "" {
		log.Printf("Redirected URL: %s", redirectURL)
	}

	// Check the URL against the VirusTotal API
	vtResult, err := vtClient.ScanURL(urlToScan)
	if err != nil {
		log.Fatalf("Error scanning URL with VirusTotal: %v", err)
	}

	// Print the results
	fmt.Printf("URL: %s\n", urlToScan)
	fmt.Printf("Domain: %s\n", domainName)
	fmt.Printf("Redirect URL: %s\n", redirectURL)
	fmt.Printf("VirusTotal Result: %s\n", vtResult)
}
