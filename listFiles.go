package main

import (
	"io/ioutil"
	"strings"
)

// ListFiles lists files inside the provided dir path.
func ListFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	var videos []string

	for file := range files {
		filetype, _ := GetFileContentType(dir + files[file].Name())
		if strings.Contains(filetype, "video") {
			videos = append(videos, files[file].Name())
		}
	}

	return videos, err
}
