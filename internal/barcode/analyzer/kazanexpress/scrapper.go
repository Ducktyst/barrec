package kazanexpress

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/tebeka/selenium"
	// https://github.com/SeleniumHQ/docker-selenium
)

var (
	KAZAN_EXPRESS_HOST = "https://kazanexpress.ru"
	KAZAN_SEARCH_URL   = "https://kazanexpress.ru/search?query={{.search_text}}&sorting=price&ordering=ascending"
)
var SEARCH_URL = KAZAN_SEARCH_URL

var OZON_SEARCH_URL = "https://www.ozon.ru/search/?from_global=true&text={{.search_text}}"

var host = "http://localhost"
var port = 4445             // TODO: to env
var browserName = "firefox" // or "chrome"

// Возвращает цену с копейками
func ParseWithSelenium(url string) (string, int, error) {
	logrus.Info("ParseWithSelenium")
	caps := selenium.Capabilities{
		"browserName": browserName,
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("%s:%d/wd/hub", host, port))
	if err != nil {
		return "", 0, fmt.Errorf(`selenium.NewRemote err %w`, err)
	}
	defer wd.Quit()

	wd.SetPageLoadTimeout(20 * time.Second)

	// начало парсинг ссылки
	productCard, err := wd.FindElement(selenium.ByCSSSelector, `[data-test-id="item__product-card"]`)
	if err != nil {
		return "", 0, fmt.Errorf(`get product card from [data-test-id="item__product-card"] error: %w`, err)
	}
	detailUrl, err := productCard.GetAttribute("href")
	if err != nil {
		return "", 0, fmt.Errorf(`get product card url error: %w`, err)
	}
	// конец парсинг сслыки

	// начало парсинг цены
	lowestPriceWe, err := wd.FindElement(selenium.ByCSSSelector, `[data-test-id="text__price"]`)
	if err != nil {
		return "", 0, fmt.Errorf(`get text from [data-test-id="text__price"] err %w`, err)
	}
	priceVal, err := lowestPriceWe.Text()
	if err != nil {
		return "", 0, fmt.Errorf(`get price value error %w`, err)
	}
	priceSrc := strings.Split(priceVal, " ")
	if len(priceSrc) < 1 {
		return "", 0, fmt.Errorf(`recognize price value %w`, err)
	}
	price, err := strconv.Atoi(priceSrc[0])
	if err != nil {
		return "", 0, fmt.Errorf(`recognize price value %w`, err)
	}
	// конец парсинг цены

	return detailUrl, price, nil
}

// примеры безинтерфейсного браузера https://github.com/chromedp/examples
