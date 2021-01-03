package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// adds room codes from room_codes.txt into an array
func getCodes() []string {
	file, err := os.Open("room_codes.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	roomCodes := []string{}

	for _, code := range txtlines {
		roomCodes = append(roomCodes, code)
	}

	return roomCodes
}

// makes api calls to endpoint using every code in roomCodes[]
func makeApiCall(roomCodes []string) {
	// base endpoint
	url := "https://blobcast.jackboxgames.com/room/"

	for _, code := range roomCodes {

		fullUrl := url + code

		resp, err := http.Get(fullUrl)
		if err != nil {
			log.Fatalln(err)
		}

		// if 404, continue
		if resp.Status == "404 Not Found" {
			continue
		}

		// read json response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// convert to string and log
		sb := string(body)
		log.Printf(sb)
	}
}

func main() {
	roomCodes := getCodes()

	// concurrency to speed up the process of finding rooms
	go makeApiCall(roomCodes[10001:20000])
	go makeApiCall(roomCodes[20001:30000])
	go makeApiCall(roomCodes[30001:40000])
	go makeApiCall(roomCodes[40001:50000])
	go makeApiCall(roomCodes[50001:60000])
	makeApiCall(roomCodes)
}
