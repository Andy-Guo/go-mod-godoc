// Package gomodgodoc contains workaround for https://github.com/golang/go/issues/26827
package gomodgodoc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Start starting godoc, saving all your package docs with wget,
// and then serving a static server.
func Start(docsDir string, godocPort int, waitTime time.Duration, replacerConfigFilePath string, replacerFilesPattern string) {
	projectModule := os.Getenv("MODULE_NAME")
	servePortStr := os.Getenv("SERVE_PORT")
	servePort, err := strconv.Atoi(servePortStr)
	if err != nil {
		panic(fmt.Errorf("invalid SERVE_PORT value: '%s'", servePortStr))
	}

	fmt.Println("[starting]", projectModule)

	downloadedSiteDir := fmt.Sprintf("%s/localhost:%d", docsDir, godocPort)
	if _, err := os.Stat(downloadedSiteDir); !os.IsNotExist(err) {
		fmt.Println("[removing old site sources]")
		if err := os.RemoveAll(downloadedSiteDir); err != nil {
			panic(err)
		}
	}

	fmt.Println("[godocs] starting...")

	cmd := exec.Command("godoc", fmt.Sprintf("-http=:%d", godocPort))
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	go func() {
		fmt.Printf("[sleeping %ds until godoc is ready] ...", waitTime)
		time.Sleep(waitTime)
		fmt.Println("[wget site downloading] starting...")
		url := fmt.Sprintf("http://localhost:%d/pkg/%s", godocPort, projectModule)

		if err := downloadSite(url, docsDir); err != nil {
			panic(err)
		}

		fmt.Println("[wget site downloading] done")

		if err := runReplacer(docsDir, replacerConfigFilePath, replacerFilesPattern); err != nil {
			panic(err)
		}

		if err := cmd.Process.Kill(); err != nil {
			panic(err)
		}
	}()

	if err := cmd.Wait(); err != nil {
		fmt.Printf("err: %+v\n", err)
	}

	startStaticServer(godocPort, servePort, projectModule)
}

func runReplacer(docsDir string, replacerConfigFilePath string, replacerFilesPattern string) error {
	bytes, err := readFile(replacerConfigFilePath)
	if err != nil {
		return nil
	}

	var replaceSettings []replaceSetting
	if err := json.Unmarshal(bytes, &replaceSettings); err != nil {
		return nil
	}

	replacer := newContentReplacer(replaceSettings, placeholdersFromEnv())
	if err := replacer.replace(docsDir, replacerFilesPattern); err != nil {
		return err
	}

	return nil
}

func startStaticServer(godocPort int, servePort int, projectModule string) {
	baseDownloaded := fmt.Sprintf("docs/localhost:%d", godocPort)

	moduleDir := fmt.Sprintf("%s/pkg/%s", baseDownloaded, projectModule)
	serveDir := substring(moduleDir, 0, strings.LastIndex(moduleDir, "/"))

	if err := os.Mkdir(fmt.Sprintf("%s/lib", serveDir), 0755); err != nil {
		panic(err)
	}
	if err := os.Rename(fmt.Sprintf("%s/lib/godoc", baseDownloaded), fmt.Sprintf("%s/lib/godoc", serveDir)); err != nil {
		panic(err)
	}

	siteDocs := http.FileServer(http.Dir(serveDir))
	http.Handle("/", siteDocs)

	fmt.Printf("Serving at http://localhost:%d\n", servePort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", servePort), nil); err != nil {
		panic(err)
	}
}