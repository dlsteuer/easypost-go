package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/EasyPost/easypost-go"
)

func main() {
	apiKey := os.Getenv("EASYPOST_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "missing API key")
		os.Exit(1)
		return
	}
	client := easypost.New(apiKey)

	// Rerate a shipment
	rates, err := client.RerateShipment("shp_123")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error retrieving rates:", err)
		os.Exit(1)
		return
	}

	prettyJSON, err := json.MarshalIndent(rates, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating JSON:", err)
		os.Exit(1)
		return
	}
	fmt.Println(string(prettyJSON))
}
