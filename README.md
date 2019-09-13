[![Go Report Card](https://goreportcard.com/badge/github.com/nightstory/go-mod-godoc)](https://goreportcard.com/report/github.com/nightstory/go-mod-godoc) [![GoDoc](https://godoc.org/github.com/nightstory/go-mod-godoc?status.svg)](https://godoc.org/github.com/nightstory/go-mod-godoc) [![GitHub license](https://img.shields.io/github/license/nightstory/go-mod-godoc)](https://github.com/nightstory/go-mod-godoc/blob/master/LICENSE)
# Go Modules godoc workaround
More about this issue: https://github.com/golang/go/issues/26827

## TL;DR
Not so easy to deal with godoc with go modules and if your project is outside of `GOPATH`, and also I had some problems with go1.13's `go doc`.
I wrote this workaround to simply host local version of godoc in docker.

## Prepare
1. Fill in `.env`, see more in the file.
2. Fill in `replacer.json`, it's very simple to do: 
  - File contains a JSON array with replacement rules;
  - All `*.html` files contents will be processed with these rules;
  - `key` is what to replace;
  - `value` is a replacement result. 
    - If `regex` is set to true, `key` will be interpreted as a regular expression and the first match will be replaced with the `value`.

## Start
Simply run `docker-compose up -d` and open `http://localhost:$SERVER_PORT` in the browser.