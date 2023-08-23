package main

import (
	"encoding/json"
	"fmt"
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

	// decode JSON into packets variable
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&packets)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Create a map to store unique IP addresses
	uniqueIPs := make(map[string]bool)

	// Loop through packets and add the IP addresses to the map
	for i := 0; i < len(packets); i++ {
		uniqueIPs[packets[i].Source.Layers.IP.Src] = true
	}

	// Print out the unique IP addresses
	fmt.Println("Unique IP addresses:")

	for ip := range uniqueIPs {
		fmt.Println(ip)
	}

	fmt.Println("Number of Packets:", len(packets))
}

// ------------------------------------------

func main() {
	ip_address("packets.json")
}
