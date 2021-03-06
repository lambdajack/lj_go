package lj_http

import (
	"log"
)

//lint:ignore U1000 unused example;
func exDownload() {

	// Specify the destination file name ("we_love_robots.txt")
	err := Download("https://duckduckgo.com/robots.txt", "path/to/out/dir", "we_love_robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the destination file name from the URL ("robots.txt")
	err = Download("https://duckduckgo.com/robots.txt", "path/to/out/dir", "")
	if err != nil {
		log.Fatal(err)
	}
}
