# kari

[![Build Status](https://travis-ci.org/ninoseki/kari.svg?branch=master)](https://travis-ci.org/ninoseki/kari)
[![Coverage Status](https://coveralls.io/repos/github/ninoseki/kari/badge.svg?branch=master)](https://coveralls.io/github/ninoseki/kari?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/3e27e0e02e734b0f9fa29903d88b9555)](https://www.cgoodacy.com/app/ninoseki/kari)

Indicator extractor API based on [cacador](https://github.com/sroberts/cacador).

## How to use

In order to send data, you must perform an HTTP POST request to the following URL:

`https://kari-extractor.herokuapp.com/extract`

The API call expects the following parameter:

- `data`: data which you want to extract indicators inside.

Example code:

```sh
curl -F "data=1.1.1.1 google.com f6f8179ac71eaabff12b8c024342109b" kari-extractor.herokuapp.com/extract
```

Example response:

```json
{
  "hashes": {
    "md5s": [
      "f6f8179ac71eaabff12b8c024342109b"
    ],
    "sha1s": null,
    "sha256s": null,
    "sha512s": null,
    "ssdeeps": null
  },
  "networks": {
    "domains": [
      "google.com"
    ],
    "emails": null,
    "ipv4s": [
      "1.1.1.1"
    ],
    "ipv6s": null,
    "urls": null
  },
  "files": {
    "docs": null,
    "exes": null,
    "flashes": null,
    "imgs": null,
    "macs": null,
    "webs": null,
    "zips": null
  },
  "utilities": {
    "cves": null
  },
  "time": "2018-05-26 02:48:46.953084088 +0000 UTC m=+45.807881187"
}
```

## How to run the app locally

```sh
$ heroku local web
[OKAY] Loaded ENV .env File as KEY=VALUE Format
4:50:27 PM web.1 |  [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
4:50:27 PM web.1 |   - using env:       export GIN_MODE=release
4:50:27 PM web.1 |   - using code:      gin.SetMode(gin.ReleaseMode)
4:50:27 PM web.1 |  [GIN-debug] Loaded HTML Templates (2):
4:50:27 PM web.1 |      -
4:50:27 PM web.1 |      - index.html
4:50:27 PM web.1 |  [GIN-debug] GET    /                         --> main.setupRouter.func1 (3 handlers)
4:50:27 PM web.1 |  [GIN-debug] POST   /extract                  --> main.setupRouter.func2 (3 handlers)
4:50:27 PM web.1 |  [GIN-debug] Listening and serving HTTP on :5000
```
