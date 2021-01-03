# JackboxTV Room Finder

Go script that finds available rooms on [jackbkox.tv](https://jackbox.tv/). You can join random ongoing games with strangers.

## Included Files

- `main.go` - the actual script
- `room_codes.txt` - textfile with all possible room code permutations. Used by `main.go`.

## Running

You'll need Go installed. Then run:

```sh
go run main.go
```

Your terminal will fill up with open rooms along with additional information about the room. For example:

```json
2021/01/03 03:23:37 {"roomid":"BDPS","server":"ecast.jackboxgames.com","apptag":"fibbage","appid":"3Mcei9GjIFpBUGwhQRtHRyGQpQUYoJfy","numAudience":0,"audienceEnabled":false,"joinAs":"player","requiresPassword":false}
2021/01/03 03:23:38 {"roomid":"CVZT","server":"ecast.jackboxgames.com","apptag":"rapbattle","appid":"5983d1cf-bf20-5def-7224-3b6b07fa0a06","numAudience":0,"audienceEnabled":true,"joinAs":"audience","requiresPassword":false}
2021/01/03 03:23:39 {"roomid":"AOVS","server":"ecast.jackboxgames.com","apptag":"quiplash","appid":"bc140c96-9c1d-4640-aaf2-a4f57c0786d1","numAudience":0,"audienceEnabled":true,"joinAs":"player","requiresPassword":false}
2021/01/03 03:23:42 {"roomid":"CWBN","server":"ecast.jackboxgames.com","apptag":"bracketeering","appid":"ba051223-0c2e-4420-8f90-1356d94d6284","numAudience":0,"audienceEnabled":true,"joinAs":"audience","requiresPassword":false}
2021/01/03 03:23:49 {"roomid":"BDVB","server":"ecast.jackboxgames.com","apptag":"fakinit","appid":"322cd918bbf2ac6af0c05b6f3b8b3ce8","numAudience":0,"audienceEnabled":true,"joinAs":"audience","requiresPassword":false}
```

`roomid` is the code you'll need to join the room.

`apptag` is the name of the game being played in the room.

## Joining a Room

To join a room, copy and paste the `roomid` of your choosing into the "ROOM CODE" field on [jackbox.tv](https://jackbox.tv/). Write down any name you want in the "NAME" field, then click "PLAY"

## FAQ 

**Won't joining a random room full of strangers ruin their experience?**

Yes.