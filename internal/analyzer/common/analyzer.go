package common

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/ducktyst/bar_recomend/internal/analyzer/kazanexpress"
)

type site uint

const (
	Citilink     = iota // == 0
	DNS          = iota // == 1
	YandexMarket = iota
	KazanExpress = iota
	Ozon         = iota
)

// getPriceFromCitilink
func GetPriceFrom(site site, articul string) (int, error) {
	switch site {
	case KazanExpress:
		url := GenerateSearchUrl(kazanexpress.SEARCH_URL, articul)
		return kazanexpress.ParseWithSelenium(url)
	}
	return 0, errors.New("unknown error")
}

// https://www.scrapingbee.com/blog/web-scraping-go/
// GenerateSearchUrl возвращает ссылку с поисковым запросов
//
func GenerateSearchUrl(url, search_text string) string {
	t := template.New("search url")
	t.Parse(url)

	buffer := bytes.NewBufferString("")
	t.Execute(buffer, map[string]string{"search_text": search_text}) // TODO: search_text to const
	return string(buffer.Bytes())
}
