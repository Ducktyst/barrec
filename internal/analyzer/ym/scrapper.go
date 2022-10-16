package ym

import (
	"fmt"
	"time"

	"github.com/ducktyst/bar_recomend/internal/analyzer/common"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

const (
	YM_SEARCH_URL = "https://market.yandex.ru/search?cvredirect=2&text={{.search_text}}"
	YM_LIST_URL   = "https://market.yandex.ru/catalog--produkty/54434/list?srnum=1002&was_redir=1&rt=9&rs=eJwzEg1grGLh-HGSdRYj14X9F7ZebADiNgBj9QrL&text={{.search_text}}&hid=91307&how=aprice&allowCollapsing=1&local-offers-first=0"
)

type price int // полная цена с копейками
func (p price) String() string {
	return fmt.Sprintf("%d.%d", p%100, p/100)
}

type product struct {
	barcode     string
	FullName    string
	sourcePrice string
	Price       price
}

func Scrap() {
	return
	// var header string

	c := colly.NewCollector(
		colly.AllowedDomains("market.yandex.ru"),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.IgnoreRobotsTxt = true
	c.SetRequestTimeout(20 * time.Second)

	c.OnHTML("body", func(h *colly.HTMLElement) {
		fmt.Println("onHTML body")
		fmt.Println(h.Text)
	})

	c.OnHTML("main", func(h *colly.HTMLElement) {
		//data-auto="mainPrice"
		fmt.Println("onHTML main")
		fmt.Println(h.Text)
		fmt.Println(h.DOM.Nodes)
		// 	fmt.Println(h.DOM.Children().Get(0))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Request to url", r.URL.String())
	})

	url := common.GenerateSearchUrl(YM_LIST_URL, "полотенце")
	c.Visit(url)
	/*
		Скажите, что вы не робот. Recaptcha
	*/
}

// func ScrapYM() {
// 	// Request the HTML page.
// 	url := common.GenerateSearchUrl(YM_LIST_URL, "полотенце") // TODO: search_text as func param

// 	res, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close()
// 	if res.StatusCode != 200 {
// 		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
// 	}

// 	// Load the HTML document
// 	doc, err := goquery.NewDocumentFromReader(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(doc)
// 	doc.Find("body").Each(func(i int, s *goquery.Selection) {
// 		fmt.Println(s.Text())
// 	})

// 	// Find the review items
// 	doc.Find("#main").Each(func(i int, s *goquery.Selection) {
// 		fmt.Println("mains")
// 		// For each item found, get the title
// 		title := s.Find("div").Text()
// 		fmt.Printf("div %d: %s\n", i, title)
// 	})
// }
