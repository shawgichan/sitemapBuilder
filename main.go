package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "", "url of the website to scrape")
	flag.Parse()

	fmt.Println(*urlFlag)

	/*
		1. Get the webpage
		2. Parse all the links on the page
		3. Build proper urls with our links
		4. Filter out any links with a different domain
		5. Find all pages (BFS)
		6. Print out XML
	*/
}
