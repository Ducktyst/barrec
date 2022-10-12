package analyzer

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
