# JackboxTV Room Finder

Go script that finds available rooms on [jackbox.tv](https://jackbox.tv/). You can join random ongoing games with strangers.

The code takes a bruteforce approach to finding empty rooms. The code hits the jackbox.tv API with different room codes and prints out open rooms.

**2023-01-25 Update**: It looks like jackbox.tv has set up some sort of throttling. This program will only work for around ~50 API requests before the IP address is throttled. At this time I have not figured out a way around this.

## Included Files

- `main.go` - the main code.
- `roomcodes.go` - contains an array with all possible room code permutations. Used in `main.go`.

## Running

You'll need Go installed. 

Clone the repo and build the binary:

```
go build
```

Then run it:

```
./jackbox-tv-room-finder
```

Your terminal will fill up with open rooms along with the room code, name of the game, and a convenient URL that auto-fills the room code for you:

```
Finding open rooms...

Room Code: OFVZ, Game: rapbattle, URL: https://jackbox.tv/OFVZ
Room Code: IOBE, Game: drawful2international, URL: https://jackbox.tv/IOBE
Room Code: RBZN, Game: monstermingle, URL: https://jackbox.tv/RBZN
Room Code: WTXS, Game: blanky-blank, URL: https://jackbox.tv/WTXS
```

## Joining a Room

To join a room, copy and paste the room code of your choosing into the "ROOM CODE" field on [jackbox.tv](https://jackbox.tv/) (or use the URL that shows up in the terminal output). Write down any name you want in the "NAME" field, then click "PLAY"

## FAQ 

**Won't joining a random room full of strangers ruin their experience?**

Yes.