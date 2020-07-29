## A Profile Video Upload Helper

__You can upload original video for Profile Video.__

***Profile Video Requirement:***

1. __Video resolution <= 800*800 and SQUARE (width and height must be equal.)__
2. __Video filesize <= 2MB__
3. __Video time <= 10 seconds__

**Bot Command:**

1. `/id`

2. Reply animation(gifs) message with `/pv {optional}`

**Command Example**

`/id`: Get current group/channel ID

`/pv`: Change your profile video

`/pv 1`: Your profile video cover is 01s:00ms

`/pv 1.11`: Your profile video cover is 01s:11ms

`/pv -10010000000000`: Change group/channel profile video

`/pv -10010000000000 1.1`: Change group/channel profile video with specific cover time

### Requirement
1. [TDLib](https://github.com/tdlib/td#building)
2. go get github.com/c0re100/go-tdlib

### Building

```
git clone https://github.com/c0re100/TelegramProfileVideoBot
cd TelegramProfileVideoBot
go build
```

### Prebuilt

[Release](https://github.com/c0re100/TelegramProfileVideoBot/releases)