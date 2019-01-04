# comix
[![Build Status](https://travis-ci.com/tizz98/comix.svg?branch=master)](https://travis-ci.com/tizz98/comix)

Comix is a raspberry pi based, daily comic calendar. Configurable with many different comic sources.

**Still a work in progress**

## Sources
1. [XKCD](https://xkcd.com/)

## Building

### Go Modules
Make sure you have this environment variable exported. `export GO111MODULE=on`

### Commands

#### Building
1. Locally `make build`, outputs `./comix` binary
1. For raspberry pi `make build-pi`, outputs `./comix-pi` binary

#### Running
note: all commands can use the `--verbose` flag for debug output.

- Download comics and display them: `./comix downloader --source xkcd`
- Version `./comix verion`

### Testing

1. Run redis `docker run -d redis -p 6379`
1. Run tests `REDIS_ADDRESS=127.0.0.1:6379 make test`

### Example image output
![](./example.png)
