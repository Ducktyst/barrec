package kazanexpress

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
	// https://github.com/SeleniumHQ/docker-selenium
)

var (
	KAZAN_EXPRESS_HOST = "https://kazanexpress.ru"
	KAZAN_SEARCH_URL   = "https://kazanexpress.ru/search?query={{.search_text}}&sorting=price&ordering=ascending"
)
var SEARCH_URL = KAZAN_SEARCH_URL

// Возвращает цену с копейками в последних двух символах
func ParseWithSelenium(wd selenium.WebDriver, url string) (string, int, error) {
	logrus.Infof("ParseWithSelenium %s", url)
	if err := wd.Get(url); err != nil {
		return "", 0, fmt.Errorf(`wd.Get err %w`, err)
	}
	// начало парсинг ссылки на карточку товара
	productCard, err := wd.FindElement(selenium.ByCSSSelector, `[data-test-id="item__product-card"]`)
	if err != nil {
		return "", 0, fmt.Errorf(`get product card from [data-test-id="item__product-card"] error: %w`, err)
	}
	detailUrlElem, err := productCard.FindElement(selenium.ByCSSSelector, `[class="subtitle-item"]`)
	if err != nil {
		return "", 0, fmt.Errorf(`get product card elem url error: %w`, err)
	}
	logrus.Debug(detailUrlElem.Text())
	detailUrl, err := detailUrlElem.GetAttribute("href")
	if err != nil {
		return "", 0, fmt.Errorf(`get product card url error: %w`, err)
	}
	// конец парсинг ссылки на карточку товара

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
	if len(priceSrc) < 1 || len(priceSrc[0]) == 0 {
		return "", 0, fmt.Errorf(`recognize price value %w`, err)
	}
	priceFloatSrc := strings.ReplaceAll(priceSrc[0], ",", ".")
	price, err := strconv.ParseFloat(priceFloatSrc, 64)
	if err != nil {
		return "", 0, fmt.Errorf(`recognize price value %w`, err)
	}
	// конец парсинг цены
	resUrl := KAZAN_EXPRESS_HOST + detailUrl
	resPrice := int(price * 100)
	return resUrl, resPrice, nil // проверить конвертацию, 100,90 => 10090 , 100,909 => 10090, а не 10091
}

// примеры безинтерфейсного браузера https://github.com/chromedp/examples
