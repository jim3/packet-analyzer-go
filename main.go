package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Packet is a struct to hold the JSON data
type Packet struct {
	Source struct {
		Layers struct {
			IP struct {
				Src string `json:"ip.src"`
			} `json:"ip"`
		} `json:"layers"`
	} `json:"_source"`
}

func ip_address(file string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	fmt.Println("Successfully Opened packets.json")
	defer jsonFile.Close()

	// create a variable to store the decoded JSON
	var packets []Packet

	// decode JSON into packets variable.
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&packets)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Create a map to store unique IP addresses as keys
	uniqueIPs := make(map[string]bool)

	// Loop through packets and add the IP addresses to the map
	for i := 0; i < len(packets); i++ {
		uniqueIPs[packets[i].Source.Layers.IP.Src] = true
	}

	fmt.Println("Number of Packets:", len(packets))

	fmt.Println("City locations for the first 10 IP addresses:")
	for ip := range uniqueIPs {
		// If it is a local IP address, skip it and continue to the next IP address https://ipinfo.io/bogon
		if ip == "bogon" {
			continue
		}
		getCityLocation(ip)
	}
}

// ------------------------------------------

// `getCityLocation` gets the IP addresses and calls ipinfo.io to get the location of the IP addresses
func getCityLocation(ipAddress string) {
	token := "YOUR_TOKEN_HERE"

	// Make an HTTP GET request to ipinfo.io to get the city information for the IP address
	resp, err := http.Get("http://ipinfo.io/" + ipAddress + "/city?token=" + token)
	if err != nil {
		fmt.Println("Error fetching city location:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("City location for IP %s: %s\n", ipAddress, string(body))
}

// ------------------------------------------

func main() {
	ip_address("packet.json")
}
