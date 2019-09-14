// Package main provides workaround application.
package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/nightstory/go-mod-godoc"
	"time"
)

const (
	// Target directory in which will all godoc
	// pages will be saved.
	// (only few html files related to your project,
	// and 2-3 .css & .js files of a godoc site)
	// It's a directory inside of a docker container.
	docsDir                = "/app/docs"

	// Need to replace some absolute links 
	// in all output files, see the replacer.json config
	replacerFilesPattern   = "*.html"

        // Path to a replacer config.
        // See README.md and replacer.go
        // It's a file inside of a docker container.
	replacerConfigFilePath = "/app/replacer.json"

	// Don't matter the port, 
	// but it's necessary to know it for downloading.
	godocPort              = 6060

	// it is necessary to wait some time
	// while `godoc -http=:PORT` is starting
	waitTime = 2 * time.Second
)

func main() {
	gomodgodoc.Start(docsDir, godocPort, waitTime, replacerConfigFilePath, replacerFilesPattern)
}
