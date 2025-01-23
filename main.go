package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/shawgichan/link"
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

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	links, _ := link.Parse(resp.Body)
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)

		}
	}
	for _, href := range hrefs {
		fmt.Println(href)
	}

}
