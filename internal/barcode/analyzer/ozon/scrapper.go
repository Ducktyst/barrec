package ozon

import (
	"fmt"
	"log"
	"time"

	"github.com/ducktyst/bar_recomend/internal/analyzer/common"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var SEARCH_URL2 = "https://www.ozon.ru/category/polotentsa-15084/?deny_category_prediction=true&from_global=true&sorting=ozon_card_price&text={{.search_text}}"
var SEARCH_URL = "https://www.ozon.ru/search/?from_global=true&text={{.search_text}}"

func Scrap() {
	// var header string

	c := colly.NewCollector(
	// colly.AllowedDomains("www.ozon.ru"),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.IgnoreRobotsTxt = true
	c.SetRequestTimeout(20 * time.Second)

	c.OnHTML("body", func(h *colly.HTMLElement) {
		fmt.Println("onHTML body")
		fmt.Println(h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Request to url", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response")
		fmt.Println("len headers", *r.Headers)

		fmt.Println(r.StatusCode)
		for k, v := range *r.Headers {
			fmt.Println(k, v)
		}
		fmt.Println(string(r.Body))
	})

	url := common.GenerateSearchUrl(SEARCH_URL, "полотенце")
	err := c.Visit(url)
	if err != nil {
		log.Fatalf("c.Visit err %v", err)
	}
	/*
		Request to url https://www.ozon.ru/search/?from_global=true&text=полотенце
		2022/10/15 20:49:53 c.Visit err Forbidden
	*/
}
