# JackboxTV Room Finder

Go script that finds available rooms on [jackbkox.tv](https://jackbox.tv/). You can join random ongoing games with strangers.

## Included Files

- `main.go` - the actual script
- `room_codes.txt` - textfile with all possible room code permutations. Used by `main.go`.

## Running

You'll need Go installed. Then run:

```
go run main.go
```

Your terminal will fill up with open rooms along with the room code and game being played:

```
Finding open rooms...

Room Code: FRYI, Game: drawful2
Room Code: CVZK, Game: fibbage2
Room Code: FRYX, Game: fibbage
Room Code: OFYE, Game: drawful
Room Code: FSBZ, Game: quiplash3
Room Code: RBYK, Game: drawful2
Room Code: CWDI, Game: awshirt
Room Code: WTXA, Game: drawful
```

## Joining a Room

To join a room, copy and paste the room code of your choosing into the "ROOM CODE" field on [jackbox.tv](https://jackbox.tv/). Write down any name you want in the "NAME" field, then click "PLAY"

## FAQ 

**Won't joining a random room full of strangers ruin their experience?**

Yes.