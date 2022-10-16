package kazanexpress

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ducktyst/bar_recomend/internal/analyzer/common"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var SEARCH_URL = "https://kazanexpress.ru/search?query={{.search_text}}&sorting=price&ordering=ascending"

func Scrap() {
	// var header string

	c := colly.NewCollector(
		// colly.AllowedDomains(""),
		colly.MaxDepth(1),
		colly.IgnoreRobotsTxt(),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	c.SetRequestTimeout(20 * time.Second)
	// c.WithTransport() // https://github.com/gocolly/colly/issues/145

	// Подготовка колбэков
	c.OnHTML("body", func(h *colly.HTMLElement) {
		fmt.Println("onHTML body")
		fmt.Println(h.Text)
	})
	c.OnHTML("#category-products", func(h *colly.HTMLElement) {
		fmt.Println("onHTML body")
		fmt.Println(h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Request to url", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response")
		fmt.Println("headers", *r.Headers)

		fmt.Println(r.StatusCode)
		for k, v := range *r.Headers {
			fmt.Println(k, v)
		}
		fmt.Println(string(r.Body))
	})

	// Подготовка запроса

	hdrs := http.Header{}
	hdrs.Set("Cookie", "tmr_lvid=e7239defe3d200918956fa91a21a04ca; tmr_lvidTS=1665854186336; _gid=GA1.2.975612261.1665854186; tmr_detect=1%7C1665854186584; _ym_uid=166585418766650638; _ym_d=1665854187; _ga=GA1.1.297887222.1665854186; mindboxDeviceUUID=75421d5a-b900-48c9-93e4-09651a8e2f63; directCrm-session=%7B%22deviceGuid%22%3A%2275421d5a-b900-48c9-93e4-09651a8e2f63%22%7D; SL_G_WPT_TO=ru; _ym_isad=1; _tt_enable_cookie=1; _ttp=76876ebe-9ee6-44b8-9b63-21083d2c398c; SL_GWPT_Show_Hide_tmp=1; SL_wptGlobTipTmp=1; _ga_BMF2JXTDDP=GS1.1.1665854186.1.1.1665854204.42.0.0; tmr_reqNum=104")

	requestDatqBuf := strings.NewReader("request data")

	url := common.GenerateSearchUrl(SEARCH_URL, "полотенце")
	// url = testdata.KazanExpressPath()

	// err := c.Visit(url)
	err := c.Request(http.MethodGet, url, requestDatqBuf, colly.NewContext(), hdrs)
	if err != nil {
		log.Fatalf("c.Visit err %v", err)
	}

	/*
		Пожалуйста, включите JavaScript, чтобы воспользоваться сайтом!
		https://github.com/gocolly/colly/issues/4
	*/
}
