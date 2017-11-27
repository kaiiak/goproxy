# Goproxy [![Build Status](https://travis-ci.org/kaiiak/goproxy.svg?branch=master)](https://travis-ci.org/kaiiak/goproxy) [![GoDoc](https://godoc.org/github.com/kaiiak/goproxy?status.svg)](https://godoc.org/github.com/kaiiak/goproxy) [![GoReport](https://goreportcard.com/badge/github.com/kaiiak/goproxy)](https://goreportcard.com/report/github.com/kaiiak/goproxy)

A socks5 proxy build with golang.

## Install
```sh
go get -u github.com/kaiiak/goproxy
```

## Basic Usage


### Server

```sh
goproxy -s -config YOUR_CONFIG_PATH
```

### Client

```sh
goproxy -c -host REMOTE_ADDRESS -port REMOTE_PORT -username YOUR_USER_NAME -password YOUR_PASSWORD
```

