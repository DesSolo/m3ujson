## m3ujson
Convert m3u file to json. For use in programs like [node-ffmpeg-mpegts-proxy](https://github.com/Jalle19/node-ffmpeg-mpegts-proxy)

# Features
- support local m3u files
- support remote (http) m3u files
- fast

### Binary releases
latest [release](https://github.com/DesSolo/m3ujson/releases/latest)

### Usage
```shell
$ ./m3ujson https://example.com/playlist.m3u8 > m3u.json
$ cat m3u.json
[
    {
        "id": 1,
        "name": "Channel Name",
        "source": "http://example.com/0001/index.m3u8"
    }
]
```

### Customize output
Just use [jq](https://jqlang.org/manual/)

```shell
$ ./m3ujson https://example.com/playlist.m3u8 | jq 'map({url: "/\(.id)", name, source})' > m3u.json
$ cat m3u.json
[
    {
        "url": "/1",
        "name": "Channel Name",
        "source": "http://example.com/0001/index.m3u8"
    }
]
```