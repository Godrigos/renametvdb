package main

import (
	"path/filepath"
	"regexp"
	"strings"
)

func getSerieName(name string) (string, string) {
	// remove file extension from the file name
	name = strings.TrimSuffix(filepath.Base(name), filepath.Ext(name))

	// remove undesires characters from file name
	rName := regexp.MustCompile(`\[.*?\]|\(.*?\)|_|\W|\d|\b`)
	rSpaces := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	serieName := rName.ReplaceAllString(name, " ")
	serieName = strings.TrimSpace(serieName)
	serieName = rSpaces.ReplaceAllString(serieName, " ")

	rNum := regexp.MustCompile(`\[.*?\]|\(.*?\)|\D`)
	epNum := rNum.ReplaceAllString(name, " ")
	epNum = strings.TrimSpace(epNum)
	epNum = rSpaces.ReplaceAllString(epNum, " ")

	return serieName, epNum
}
