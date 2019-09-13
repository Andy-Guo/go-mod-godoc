// Package main provides workaround application.
package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/nightstory/go-mod-godoc/gomodgodoc"
	"time"
)

const (
	docsDir                = "docs"
	replacerFilesPattern   = "*.html"
	replacerConfigFilePath = "../replacer.json"
	godocPort              = 6060

	// it is necessary to wait some time
	// while `godoc -http=:PORT` is starting
	waitTime = 2 * time.Second
)

func main() {
	gomodgodoc.Start(docsDir, godocPort, waitTime, replacerFilesPattern, replacerConfigFilePath)
}
