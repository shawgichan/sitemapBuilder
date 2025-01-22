package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
	1. Get the webpage
	2. Parse all the links on the page
	3. Build proper urls with our links
	4. Filter out any links with a different domain
	5. Find all pages (BFS)
	6. Print out XML
*/

func main() {
	urlFlag := flag.String("url", "", "url of the website to scrape")
	flag.Parse()

	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}
