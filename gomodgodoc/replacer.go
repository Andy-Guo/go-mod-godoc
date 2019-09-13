package gomodgodoc

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type replaceSetting struct {
	IsRegExp bool   `json:"regex"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type contentReplacer struct {
	settings                 []replaceSetting
	replaceValuePlaceholders map[string]string
}

func newContentReplacer(settings []replaceSetting, keyPlaceholders map[string]string) *contentReplacer {
	return &contentReplacer{settings: settings, replaceValuePlaceholders: keyPlaceholders}
}

func (replacer contentReplacer) replace(dir string, pattern string) error {
	pathsList, err := filesList(dir, pattern)

	if err != nil {
		panic(err)
	}

	for _, path := range pathsList {
		bytes, err := ioutil.ReadFile(path)

		if err != nil {
			return err
		}

		oldContents := string(bytes)
		newContents := oldContents

		for _, setting := range replacer.settings {
			if setting.IsRegExp {
				re := regexp.MustCompile(setting.Key)
				newContents = re.ReplaceAllString(newContents, applyPlaceholders(replacer.replaceValuePlaceholders, setting.Value))
			} else {
				newContents = strings.ReplaceAll(newContents, setting.Key, applyPlaceholders(replacer.replaceValuePlaceholders, setting.Value))
			}
		}

		if err := ioutil.WriteFile(path, []byte(newContents), 0); err != nil {
			return err
		}
	}

	return nil
}

func applyPlaceholders(placeholders map[string]string, str string) string {
	newStr := str
	for k, v := range placeholders {
		newStr = strings.ReplaceAll(newStr, k, v)
	}
	return newStr
}
