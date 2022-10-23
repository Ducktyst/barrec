package barcode

import (
	"errors"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"

	"github.com/ducktyst/bar_recomend/internal/analyzer/common"
)

type Price struct {
	priceSrc float64
	Price    int // цена с копейками 115р.65коп. == 11565
	ShopUrl  string
	ShopName string
}
type ProductInfo struct {
	Name    string
	BarCode string
	Prices  []Price
}

var (
	ProductNotFound  = errors.New("product with _selected_ barcode not found")
	barcode_list_url = "https://barcode-list.ru/barcode/RU/Поиск.htm?barcode={{.search_text}}"
	progaonline_url  = `curl 'https://progaonline.com/kod' \
  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
  -H 'Accept-Language: ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7' \
  -H 'Cache-Control: max-age=0' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Cookie: _pk_id.8.ce68=ebd2b5aae10f6673.1666460131.; _ga=GA1.2.2002941948.1666460131; _gid=GA1.2.1934568619.1666460131; _ym_uid=1666460131463303529; _ym_d=1666460131; _ym_isad=2; _pk_ref.8.ce68=%5B%22%22%2C%22%22%2C1666517581%2C%22https%3A%2F%2Fyandex.ru%2F%22%5D; _pk_ses.8.ce68=1; _ym_visorc=w; XSRF-TOKEN=eyJpdiI6IjVKWEtzRk9wM09wa1FaZWxrSnV4ZVE9PSIsInZhbHVlIjoiN1lqXC8zYzJzTkJCSjFqU2pkblBiM2RuVTdQVnVwcXBtRkRIUE1kSzNPd1NOTXNMTDVFOHhVNjY4NWlUZTBZRFYiLCJtYWMiOiI2NmRhODY2OGVhZGI3ZDE3YzgwMTEwZmZkNjNjNmRjMjA2YzFlMTQ0ODZmYTUxYjE4MGM5NTdmYWJjYWMwNTE5In0%3D; progaonlinecom_session=eyJpdiI6IjBGTjNUUE5ZN2picE85ckJaVkRrd0E9PSIsInZhbHVlIjoibWN2aGtqT0phQUdTcXhGaEdRczdkUE41Z2haUHd6VDlROXZHWnZ5UFZlb3JreHhYMVwvdkV5b2lEeW8xRndkXC9cLyIsIm1hYyI6IjRkYTg5MDUxNjdkNGNkYTUxZTRlNDdkNGVjNjMxZmUwNzAzOTcwMWRlNDRhZDlhNTI5ZDA1MGRiOTBmYjBlOGUifQ%3D%3D' \
  -H 'Origin: https://progaonline.com' \
  -H 'Referer: https://progaonline.com/kod' \
  -H 'Sec-Fetch-Dest: document' \
  -H 'Sec-Fetch-Mode: navigate' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'Sec-Fetch-User: ?1' \
  -H 'Upgrade-Insecure-Requests: 1' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36 OPR/91.0.4516.77' \
  -H 'sec-ch-ua: "Not-A.Brand";v="99", "Opera";v="91", "Chromium";v="105"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  --data-raw '_token=zywGhDyGdOpI5eOa1U3mFt2f8rsBO4gMcFXkeROU&text={{.search_text}}' \
  --compressed`
)

func GetProductArticul(barcode string) (string, error) {
	url := common.GenerateSearchUrl(barcode_list_url, barcode)

	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.IgnoreRobotsTxt(),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	c.SetRequestTimeout(20 * time.Second)

	var name string
	c.OnHTML("[class=randomBarcodes] > tbody > tr:nth-child(2) > td:nth-child(3)", func(h *colly.HTMLElement) {
		name = h.Text
	})
	c.Visit(url)

	if name == "" {
		return "", ProductNotFound
	}
	return name, nil
}

func ScrapHTML() {
	// Подготовка запроса

	hdrs := http.Header{}
	hdrs.Set("Cookie", "tmr_lvid=e7239defe3d200918956fa91a21a04ca; tmr_lvidTS=1665854186336; _gid=GA1.2.975612261.1665854186; tmr_detect=1%7C1665854186584; _ym_uid=166585418766650638; _ym_d=1665854187; _ga=GA1.1.297887222.1665854186; mindboxDeviceUUID=75421d5a-b900-48c9-93e4-09651a8e2f63; directCrm-session=%7B%22deviceGuid%22%3A%2275421d5a-b900-48c9-93e4-09651a8e2f63%22%7D; SL_G_WPT_TO=ru; _ym_isad=1; _tt_enable_cookie=1; _ttp=76876ebe-9ee6-44b8-9b63-21083d2c398c; SL_GWPT_Show_Hide_tmp=1; SL_wptGlobTipTmp=1; _ga_BMF2JXTDDP=GS1.1.1665854186.1.1.1665854204.42.0.0; tmr_reqNum=104")

	// requestDatqBuf := strings.NewReader("request data")

	// url := generateSearchUrl(SEARCH_URL, "полотенце")
	// url = testdata.KazanExpressPath()

	// err := c.Visit(url)
	// err := c.Request(http.MethodGet, url, requestDatqBuf, colly.NewContext(), hdrs)
	// if err != nil {
	// log.Fatalf("c.Visit err %v", err)
	// }

}
