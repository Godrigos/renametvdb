package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var countNotFound int
	var count int
	path := "./"
	list, err := ListFiles(path)
	if err != nil {
		log.Println(err)
	}

	for i := range list {
		search, ep := getSerieName(list[i])
		eps, code := getEpisodeInfo(search, ep)

		if code == 401 {
			fmt.Println("Not authorized!")
		} else if code == 404 {
			fmt.Printf("No data on episode \"%s\" of series \"%s\"!\n",
				ep, search)
			countNotFound++
		} else if code == 405 {
			fmt.Printf("Series \"%s\" not found!\n", search)
		} else {
			rName := regexp.MustCompile(`[!?:@|</*\"\\>]`)
			rSpaces := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
			newName := rName.ReplaceAllString(search, " ")
			newName = strings.TrimSpace(newName)
			newName = rSpaces.ReplaceAllString(newName, " ")

			newName = fmt.Sprintf("%s S%02dE%02d", search, eps.Data[0].AiredSeason,
				eps.Data[0].AiredEpisodeNumber)
			ext := filepath.Ext(list[i])
			err := os.Rename(path+list[i], path+newName+ext)
			if err != nil {
				panic(err)
			}
			fmt.Println(list[i], "->", newName+ext)
			count++
		}
	}
	fmt.Printf("\n%d file(s) renamed!\n", count)
}
