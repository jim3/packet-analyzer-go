package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

	var packets []Packet

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&packets)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	uniqueIPs := make(map[string]bool)

	for i := 0; i < len(packets); i++ {
		uniqueIPs[packets[i].Source.Layers.IP.Src] = true
	}

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
