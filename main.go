package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Room struct {
	RoomId           string
	Server           string
	AppTag           string
	AppId            string
	NumAudience      int
	AudienceEnabled  bool
	JoinAs           string
	RequiresPassword bool
}

// adds room codes from room_codes.txt into array roomCodes[]
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
func findRooms(roomCodes []string) {
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

		// get json response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// unmarshall json into struct
		var room Room
		json.Unmarshal([]byte(body), &room)

		if room.JoinAs == "audience" {
			continue
		}

		if room.RequiresPassword {
			continue
		}

		fmt.Printf("Room Code: %s, Game: %s \n", color.GreenString(room.RoomId), color.GreenString(room.AppTag))
	}
}

func main() {
	fmt.Println("Finding open rooms...\n")

	roomCodes := getCodes()

	// concurrency to speed up the process of finding rooms
	go findRooms(roomCodes[100000:200000])
	go findRooms(roomCodes[200001:300000])
	go findRooms(roomCodes[300001:400000])
	go findRooms(roomCodes[400001:456976])
	findRooms(roomCodes) // 0 - 100000
}
