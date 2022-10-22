package common

import (
	"bytes"
	"text/template"
)

type site uint

const (
	Citilink     = iota // == 0
	dns          = iota // == 1
	yandexmarket = iota
	KazanExpress = iota
	ozon         = iota
)

// getPriceFromCitilink
func GetPriceFrom(site site, articul string) int {
	switch site {
	case Citilink:
		return getPriceFromCitilink(articul)
	case dns:
		return getPriceFromDns(articul)
	case kazanexpress:
		return kazanexpress.ParseWithSelenium(articul)
	}
	return 0
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
