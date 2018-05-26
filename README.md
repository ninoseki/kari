# kari

Indicator extractor API based on [cacador](https://github.com/sroberts/cacador).

## How to use

In order to send a data you must perform an HTTP POST request to the following URL:

`https://desolate-springs-18074.herokuapp.com/extract`

The API call expects the following parameter:

- `data`: data which you want to extract indicators inside.

Example code:

```sh
url -F "data=1.1.1.1\ngoogle.com\nf6f8179ac71eaabff12b8c024342109b" desolate-springs-18074.herokuapp.com/extract
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
      "ngoogle.com"
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
