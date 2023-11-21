package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
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

// makes api calls to endpoint using every code in roomCodes[]
func findRooms(roomCodes []string) {
	// base endpoint
	url := "https://blobcast.jackboxgames.com/room/"

	// infinite loop
	for {
		// getting random room code
		randomIndex := rand.Intn(len(roomCodes))
		code := roomCodes[randomIndex]

		fullUrl := url + code

		resp, err := http.Get(fullUrl)
		if err != nil {
			log.Fatalln(err)
		}

		// if 403, continue
		if resp.Status == "403 Forbidden" {
			fmt.Println("'403 Forbidden' status code. Your IP is likely blocked.")
			continue
		}

		// if 404, continue
		if resp.Status == "404 Not Found" {
			continue
		}

		// if 500, continue
		if resp.Status == "500 Internal Server Error" {
			continue
		}

		// get json response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// if body is empty, continue
		if len(body) == 0 {
			fmt.Println("Empty response body. Your IP might be throttled/blacklisted")
			continue
		}

		// deserialize json into struct
		var room Room
		err = json.Unmarshal([]byte(body), &room)
		if err != nil {
			fmt.Println(fullUrl)
			fmt.Println(resp.Status)
			fmt.Println(body)
			log.Fatalln(err)
		}

		// you can't participate in these games
		if room.JoinAs == "audience" {
			continue
		}

		// no point in trying to join a password-protected game
		if room.RequiresPassword {
			continue
		}

		// sometimes the Room ID is empty, ignore and continue
		if room.RoomId == "" {
			continue
		}

		fmt.Printf("Room Code: %s, Game: %s\n", color.GreenString(room.RoomId), color.GreenString(room.AppTag))
	}
}

func main() {
	fmt.Println("Finding open rooms...")

	// seed for randomizer. used for getting random room codes from RoomCodes slice
	rand.Seed(time.Now().UnixNano())

	// concurrency to speed up the process of finding rooms
	go findRooms(RoomCodes)
	go findRooms(RoomCodes)
	go findRooms(RoomCodes)
	go findRooms(RoomCodes)
	go findRooms(RoomCodes)
	go findRooms(RoomCodes)
	findRooms(RoomCodes)
}
