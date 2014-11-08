package scraper

import (
    "flag"
)

var (
    url = flag.String("url", "http://www.google.com", "url to scrape")
)


func init() {
    flag.Parse()
}


func getHTML(url string) {

}
