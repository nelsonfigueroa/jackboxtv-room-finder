# JackboxTV Room Finder

Python script that finds available rooms on [jackbkox.tv](https://jackbox.tv/). You can join random ongoing games with strangers.

This version is a Python port of the original golang script

## Included Files

- `roomfinder.py` - the actual script.
- `roomcodes.py` - contains an array with all possible room code permutations. Used in `roomfinder.py`.

## Running

You'll need Python3 installed. 

```
python3 ./jackbox-tv-room-finder
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
