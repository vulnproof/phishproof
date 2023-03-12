package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/vulnproof/phishproof/virustotal"
)

func main() {
	// Load the configuration file
	viper.SetConfigFile("config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	servTheApp()
	urltocheck := getFormData()
	print(virustotal.ScanURLWithVT(urltocheck, viper.GetString("virustotal.apikey")))

}
