package common

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

const (
	DNS_SEARCH_URL      = ""
	CITILINK_SEARCH_URL = ""
	YM_SEARCH_URL       = "https://market.yandex.ru/search?cvredirect=2&text={{.search_text}}"
	YM_LIST_URL         = "https://market.yandex.ru/catalog--produkty/54434/list?srnum=1002&was_redir=1&rt=9&rs=eJwzEg1grGLh-HGSdRYj14X9F7ZebADiNgBj9QrL&text={{.search_text}}&hid=91307&how=aprice&allowCollapsing=1&local-offers-first=0"
)

/*
css selectors:
https://www.w3schools.com/cssref/css_selectors.asp

<div style="box-sizing: border-box; padding-top: 0px; padding-bottom: 4060px; margin-top: 0px;" data-test-id="virtuoso-item-list">
блок всех товаров
</div>

<span data-autotest-value="129" data-autotest-currency="₽" data-auto="mainPrice"><span>129</span>&nbsp;<span class="-B-PA">₽</span></span>
*/
// https://www.scrapingbee.com/blog/web-scraping-go/
func getPriceFromCitilink(articul string) int {
	return 0
}

func getPriceFromDns(articul string) int {
	return 0
}

func getPriceYandexMarket(articul string) int {
	c := colly.NewCollector()
	c.OnHTML("div[data-test-id]", func(h *colly.HTMLElement) {})
	return 0
}

func CollyExample() {
	c := colly.NewCollector()
	var ch = make(chan int, 0)
	go func(ch chan int) {
		time.Sleep(time.Second * 2)
		ch <- 1
	}(ch)

	go func() {
		// Find and visit all links
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			e.Request.Visit(e.Attr("href"))
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		c.Visit("http://go-colly.org/")

	}()

	<-ch
	return
}

func ScapPort() {
	// Temp Variables
	var tcountry, tport string

	// Colly collector
	c := colly.NewCollector()

	//Ignore the robot.txt
	c.IgnoreRobotsTxt = true
	// Time-out after 20 seconds.
	c.SetRequestTimeout(20 * time.Second)
	// use random agents during requests
	extensions.RandomUserAgent(c)

	// set limits to colly opoeration
	c.Limit(&colly.LimitRule{
		//  // Filter domains affected by this rule
		DomainGlob: "searates.com/*",
		//  // Set a delay between requests to these domains
		Delay: 1 * time.Second,
		//  // Add an additional random delay
		RandomDelay: 3 * time.Second,
	})

	// Find and visit all country links
	c.OnHTML("#clist", func(e *colly.HTMLElement) {
		// fmt.Println("Country List: ", h.ChildAttrs("a", "href"))
		e.ForEach("li.col-xs-6.col-md-3", func(_ int, el *colly.HTMLElement) {
			tcountry = el.ChildText("a")
			link := el.ChildAttr("a", "href")
			fmt.Println("Country: ", tcountry, link)
			e.Request.Visit(link)
		})

	})

	// Find and visit all ports links
	c.OnHTML("#plist", func(h *colly.HTMLElement) {
		// fmt.Println("Port List: ", h.ChildAttrs("a", "href"))
		h.ForEach("li.col-xs-6.col-md-3", func(_ int, el *colly.HTMLElement) {
			tport = el.ChildText("a")
			link := el.ChildAttr("a", "href")
			fmt.Println("Port: ", tport, link)
			h.Request.Visit(link)
		})
	})

	// Find and visit all ports info page
	c.OnHTML("div.row", func(e *colly.HTMLElement) {
		portAuth := e.ChildText("table#port_det tbody:nth-child(1) tr:nth-child(2) td:nth-child(2)")
		fmt.Println("Port Authority: ", portAuth)
	})

	c.Visit("https://www.searates.com/maritime/")
}

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}
