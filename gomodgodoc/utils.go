package gomodgodoc

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func readFile(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func downloadSite(url string, dir string) error {
	cmdWget := exec.Command("wget", "-r", "-np", "-N", "-E", "-p", "-k", url, "-e", "robots=off")
	cmdWget.Dir = dir

	if _, err := cmdWget.Output(); err != nil {
		return err
	}

	return nil
}

func substring(input string, start int, length int) string {
	if len(input) == 0 {
		return input
	}

	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func placeholdersFromEnv() map[string]string {
	result := map[string]string{}
	arr := os.Environ()

	for _, envVar := range arr {
		kv := strings.Split(envVar, "=")
		k := kv[0]

		if strings.HasPrefix(k, "PLACEHOLDER_") {
			result[fmt.Sprintf("${%s}", substring(k, 12, len(k) - 12))] = kv[1]
		}
	}
	return result
}

func filesList(path string, pattern string) ([]string, error) {
	fileList := make([]string, 0)
	e := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if matched, err := filepath.Match(pattern, f.Name()); matched && err == nil {
			fileList = append(fileList, path)
		}
		return err
	})

	return fileList, e
}