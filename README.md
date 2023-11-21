# JackboxTV Room Finder

Go script that finds available rooms on [jackbox.tv](https://jackbox.tv/). You can join random ongoing games with strangers.

The code takes a bruteforce approach to finding empty rooms. The code hits the jackbox.tv API with different room codes and prints out open rooms.

**2023-01-25 Update**: It looks like jackbox.tv has set up rate limiting. If your IP address is rate limited, you will need a new IP address before trying again.

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

Your terminal will fill up with open rooms along with the room code and name of the game:

```
Finding open rooms...

Room Code: OFVZ, Game: rapbattle
Room Code: IOBE, Game: drawful2international
Room Code: RBZN, Game: monstermingle
Room Code: WTXS, Game: blanky-blank
```

## Joining a Room

To join a room, copy and paste the room code of your choosing into the "ROOM CODE" field on [jackbox.tv](https://jackbox.tv/). Write down any name you want in the "NAME" field, then click "PLAY"

## FAQ 

**Won't joining a random room full of strangers ruin their experience?**

Yes.