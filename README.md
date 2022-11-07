# JackboxTV Room Finder

Go script that finds available rooms on [jackbkox.tv](https://jackbox.tv/). You can join random ongoing games with strangers.

## Included Files

- `main.go` - the actual script.
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