package common

import (
	"bytes"
	"text/template"
)

type site uint

const ( // iota is reset to 0
	citilink = iota // c0 == 0
	dns      = iota // c1 == 1
	c2       = iota // c2 == 2
)

// getPriceFromCitilink
func GetPriceFrom(site site, articul string) int {
	switch site {
	case citilink:
		return getPriceFromCitilink(articul)
	case dns:
		return getPriceFromDns(articul)
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
