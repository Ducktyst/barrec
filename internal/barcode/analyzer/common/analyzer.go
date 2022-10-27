package common

import (
	"bytes"
	"errors"
	"net/url"
	"text/template"

	"github.com/ducktyst/bar_recomend/internal/barcode/analyzer/kazanexpress"
)

type site uint

const (
	Citilink     = iota // == 0
	DNS          = iota // == 1
	YandexMarket = iota
	KazanExpress = iota
	Ozon         = iota
)

type Recommendation struct {
	Name  string
	Price int
	Url   string
}

// getPriceFromCitilink
func GetPriceFrom(site site, articul string) (Recommendation, error) {
	switch site {
	case KazanExpress:
		url := GenerateSearchUrl(kazanexpress.SEARCH_URL, articul)
		url, price, err := kazanexpress.ParseWithSelenium(url)

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
