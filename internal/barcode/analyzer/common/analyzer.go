package common

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"text/template"
	"time"

	"github.com/ducktyst/bar_recomend/internal/barcode/analyzer/kazanexpress"
	"github.com/tebeka/selenium"
)

type site uint

const (
	Citilink     = iota // == 0
	DNS          = iota // == 1
	YandexMarket = iota
	KazanExpress = iota
	Ozon         = iota
)

var host = "http://localhost"
var port = 4445             // TODO: to env
var browserName = "firefox" // or "chrome"

type Recommendation struct {
	Name  string
	Price int
	Url   string
}

// getPriceFromCitilink
func GetPriceFrom(site site, articul string) (Recommendation, error) {
	caps := selenium.Capabilities{
		"browserName": browserName,
	}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("%s:%d/wd/hub", host, port)) // move to global context?
	if err != nil {
		return Recommendation{}, err
	}
	defer wd.Quit()

	switch site {
	case KazanExpress:
		wd.SetImplicitWaitTimeout(30 * time.Second)
		url := GenerateSearchUrl(kazanexpress.SEARCH_URL, articul)
		url, price, err := kazanexpress.ParseWithSelenium(wd, url)

		return Recommendation{
			Name:  articul,
			Price: price,
			Url:   url}, err
	}
	return Recommendation{}, errors.New("unknown error")
}

// https://www.scrapingbee.com/blog/web-scraping-go/
// GenerateSearchUrl возвращает ссылку с поисковым запросов
//
func GenerateSearchUrl(search_url, search_text string) string {
	t := template.New("search url")
	t.Parse(search_url)

	buffer := bytes.NewBufferString("")
	search_text_encoded := url.QueryEscape(search_text)
	t.Execute(buffer, map[string]string{"search_text": search_text_encoded}) // TODO: search_text placeholder to const

	return string(buffer.Bytes())
}
