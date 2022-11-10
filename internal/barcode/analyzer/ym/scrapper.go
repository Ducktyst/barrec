package ym

import (
	"fmt"

	"github.com/tebeka/selenium"
)

const (
	YM_HOST       = "https://market.yandex.ru/"
	YM_SEARCH_URL = "https://market.yandex.ru/search?cvredirect=2&text={{.search_text}}"
	YM_LIST_URL   = "https://market.yandex.ru/catalog--produkty/54434/list?srnum=1002&was_redir=1&rt=9&rs=eJwzEg1grGLh-HGSdRYj14X9F7ZebADiNgBj9QrL&text={{.search_text}}&hid=91307&how=aprice&allowCollapsing=1&local-offers-first=0"
)

type price int // полная цена с копейками
func (p price) String() string {
	return fmt.Sprintf("%d.%d", p%100, p/100)
}

type product struct {
	barcode     string
	FullName    string
	sourcePrice string
	Price       price
}

// Возвращает цену с копейками в последних двух символах
func ParseWithSelenium(wd selenium.WebDriver, url string) (string, int, error) {
	if err := wd.Get(url); err != nil {
		return "", 0, fmt.Errorf(`wd.Get err %w`, err)
	}
	// начало парсинг ссылки на карточку товара
	productCard, err := wd.FindElement(selenium.ByCSSSelector, `[data-test-id="item__product-card"]`)
	if err != nil {
		return "", 0, fmt.Errorf(`get product card from [data-test-id="item__product-card"] error: %w`, err)
	}
	detailUrl, err := productCard.GetAttribute("href")
	if err != nil {
		return "", 0, fmt.Errorf(`get product card url error: %w`, err)
	}
	// конец парсинг ссылки на карточку товара

	// начало парсинг цены
	// lowestPriceWe, err := wd.FindElement(selenium.ByCSSSelector, `[data-test-id="text__price"]`)
	// if err != nil {
	// 	return "", 0, fmt.Errorf(`get text from [data-test-id="text__price"] err %w`, err)
	// }
	// priceVal, err := lowestPriceWe.Text()
	// if err != nil {
	// 	return "", 0, fmt.Errorf(`get price value error %w`, err)
	// }
	// priceSrc := strings.Split(priceVal, " ")
	// if len(priceSrc) < 1 || len(priceSrc[0]) == 0 {
	// 	return "", 0, fmt.Errorf(`recognize price value %w`, err)
	// }
	// priceFloatSrc := strings.ReplaceAll(priceSrc[0], ",", ".")
	// price, err := strconv.ParseFloat(priceFloatSrc, 64)
	// if err != nil {
	// 	return "", 0, fmt.Errorf(`recognize price value %w`, err)
	// }
	var price = 0
	// конец парсинг цены
	return YM_HOST + detailUrl, int(price * 100), nil // проверить конвертацию, 100,90 => 10090 , 100,909 => 10090, а не 10091
}
