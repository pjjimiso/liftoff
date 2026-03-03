This is a simple program written in Go that queries [Launch Library 2](https://ll.thespacedevs.com/2.3.0/launches/upcoming) for the next launch from Vandenberg SFB and displays a countdown timer. This was originally designed to fit nicely into my tmux status bar so that I could see when the next launch was from my terminal. I gradually added more features such as being able to see additional launch details and being able to show additional launch countdown timers beyond just the next one.

#### Arguments
`-location`: Location ID to query (default: 11, Vandenberg SFB)
`-limit`: number of upcoming launches to show
`-full`: show full launch details

#### Usage:
Basic usage shows a countdown to the next launch
```
./liftoff-linux-amd64
2026/03/03 15:00:17 reading from cache file: /home/pjjimiso/.cache/liftoff_cache.json
🚀 22h 59m
```

Redirect to stderr to suppress additional logging output
```
./liftoff-linux-amd64 2>/dev/null
🚀 22h 59m
```

Show countdown for the next 3 launches
```
./liftoff-linux-amd64 -limit 3 2>/dev/null
```

Show full launch details
```
./liftoff-linux-amd64 -full
2026/03/03 15:21:03 reading from cache file: /home/pjjimiso/.cache/liftoff_cache.json
Name: Falcon 9 Block 5 | Starlink Group 17-18
Status: Go for Launch
Launch Time (local): 2026-03-04 14:00:00 MST
Service Provider: SpaceX
Rocket: Falcon 9 Block 5
Mission: Starlink Group 17-18 - A batch of 25 satellites for the Starlink mega-constellation - SpaceX's project for space-based Internet communication system.
Launch Pad: Space Launch Complex 4E - Vandenberg SFB, CA, USA
Launch details last updated 2026-02-25 18:03:29 +0000 UTC
🚀 22h 38m
```
