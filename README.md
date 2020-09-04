## m3ujson
Convert m3u file to json. For use in programs like [node-ffmpeg-mpegts-proxy](https://github.com/Jalle19/node-ffmpeg-mpegts-proxy)

### Support iptv providers
 - Edem.tv

### Binary releases
 - linux amd64
 - mips
 > latest [release](https://github.com/DesSolo/m3ujson/releases/latest)

### Usage
```shell
$ ./m3ujson-1.0.0-linux-amd64 https://example.com/playlist.m3u8 > m3u.json
$ cat m3u.json
[
    {
        "name": "Channel Name",
        "provider": "Provider Name",
        "url": "/0001",
        "source": "http://example.com/0001/index.m3u8"
    }
]
```