## A Profile Video Upload Helper

__You can upload original video for Profile Video.__

***Profile Video Requirement:***

1. __Video resolution <= 1200*1200 and SQUARE (width and height must be equal.)__
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
1. [modded TDLib](https://github.com/c0re100/td)

### Building

```
git clone https://github.com/c0re100/TelegramProfileVideoBot
cd TelegramProfileVideoBot
go build
```

### Prebuilt

[Release](https://github.com/c0re100/TelegramProfileVideoBot/releases)