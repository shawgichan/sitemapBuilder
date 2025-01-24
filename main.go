package main

import (
	"flag"
	"fmt"
	"io"
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

	pages := get(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
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
	return hrefs(resp.Body, base)
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}
