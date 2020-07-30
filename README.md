## A Profile Video Upload Helper

__You can upload original video for Profile Video.__

***Profile Video Requirement:***

1. __Video resolution <= 800*800 and SQUARE (width and height must be equal.)__
2. __Video filesize <= 2MB__
3. __Video time <= 10 seconds__

**How to use**

To change your profile video, please upload/forward the animation(gif) to Saved Messages first and reply this with `/pv`.

**Bot Command:**

1. `/id`

2. Reply animation(gif) message with `/pv {optional}`

**Command Example**

`/id`: Get current group/channel ID

`/pv`: Change your/group profile video

`/pv 1`: Your/Group profile video cover is 01s:00ms

`/pv 1.11`: Your/Group profile video cover is 01s:11ms

`/pv -10010000000000`: Change group/channel profile video

`/pv -10010000000000 1.1`: Change group/channel profile video with specific cover time

### Requirement
1. [TDLib](https://github.com/tdlib/td#building)
2. go get github.com/c0re100/go-tdlib
3. go get golang.org/x/crypto/ssh/terminal

### Building

```
git clone https://github.com/c0re100/TelegramProfileVideoBot
cd TelegramProfileVideoBot
go build
```

### Prebuilt

[Release](https://github.com/c0re100/TelegramProfileVideoBot/releases)