package ym

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
)

const (
	YM_HOST       = "https://market.yandex.ru/"
	YM_SEARCH_URL = "https://market.yandex.ru/search?cvredirect=2&text={{.search_text}}"
	YM_LIST_URL   = "https://market.yandex.ru/catalog--produkty/54434/list?srnum=1002&was_redir=1&rt=9&rs=eJwzEg1grGLh-HGSdRYj14X9F7ZebADiNgBj9QrL&text={{.search_text}}&hid=91307&how=aprice&allowCollapsing=1&local-offers-first=0"
	SEARCH_URL    = YM_LIST_URL
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
		return "", 0, fmt.Errorf(`ym: wd.Get err %w`, err)
	}
	logrus.Info("url ", url)
	// начало парсинг ссылки на карточку товара
	/*
		<div class="_37suf" data-tid="7e29d3e" data-tid-prop="7e29d3e">
			<div class="_3Fff3" data-tid="7e29d3e" data-tid-prop="7e29d3e"><h3 class="_2UHry _2vVOc" data-tid="1410220a" data-tid-prop="1410220a" data-zone-name="title">
			<a href="/product--polotentse-miks/101795672345?nid=18031541&amp;show-uid=16681107038922576327316002&amp;context=search&amp;text=%20%D0%9F%D0%9E%D0%9B%D0%9E%D0%A2%D0%95%D0%9D%D0%A6%D0%95%20%D0%9C%D0%90%D0%A5%D0%A0%D0%9E%D0%92%D0%9E%D0%95&amp;rs=eJwzEg1grGLh-HGSdRYj14X9F7ZebADiNgBj9QrL&amp;sku=101795672345&amp;cpc=ZlVRFCkVKySiReK0ltfKS1ql9j3RQ3YgqgJ21KH1_7IkhMRzqEsSXabhrHnRiZMWu74oCyhDRIwCYfCeX9M1VVj6EaZV9bSYIx_keVA0yPbGaMRonqMqi0f0G5b_6iEMnsrcEeHNNPhOidYKoYnQnp5lRbPHVK3Z1wW6JPQSqCSQDz_QifTxjFWxs37uAb28&amp;do-waremd5=Vhia5dEsmtp6iCPjstwgeA" target="_blank" class="_2f75n _24Q6d cia-cs" title="Полотенце микс" data-tid="931ebae7" data-tid-prop="931ebae7" data-baobab-name="title" data-node-cache-key="title-16681107038922576327316002" data-node-id="h6og1y"><span class="_29F8F" data-tid="2e5bde87">Полотенце</span><span class="" data-tid="2e5bde87"> микс</span></a></h3></div><div class="ipd7l" data-tid="7e29d3e" data-tid-prop="7e29d3e"><div class="fUyko _2LiqB" data-tid="258b22d7">Примечание : микс - Выбор конкретных цветов и моделей не предоставляется. На фотографиях могут быть представлены не все варианты.</div><div></div><div class="" data-zone-name="jump-table"></div></div></div>
	*/
	// или по Xpath
	productCard, err := wd.FindElement(selenium.ByCSSSelector, `[data-test-id="item__product-card"]`)
	if err != nil {
		return "", 0, fmt.Errorf(`ym: get product card from [data-test-id="item__product-card"] error: %w`, err)
	}
	detailUrl, err := productCard.GetAttribute("href")
	if err != nil {
		return "", 0, fmt.Errorf(`ym: get product card url error: %w`, err)
	}
	// конец парсинг ссылки на карточку товара

	// начало парсинг цены
	/*
		<span data-autotest-value="398"
				data-autotest-currency="₽"
				data-auto="mainPrice"><span>398</span>
					&nbsp;<span class="-B-PA">₽</span>
		</span>
	*/
	lowestPriceWe, err := wd.FindElement(selenium.ByCSSSelector, `[data-test-id="text__price"]`)
	if err != nil {
		return "", 0, fmt.Errorf(`ym: get text from [data-test-id="text__price"] err %w`, err)
	}
	priceVal, err := lowestPriceWe.Text()
	if err != nil {
		return "", 0, fmt.Errorf(`ym: get price value error %w`, err)
	}
	priceSrc := strings.Split(priceVal, " ")
	if len(priceSrc) < 1 || len(priceSrc[0]) == 0 {
		return "", 0, fmt.Errorf(`ym: ecognize price value %w`, err)
	}
	priceFloatSrc := strings.ReplaceAll(priceSrc[0], ",", ".")
	price, err := strconv.ParseFloat(priceFloatSrc, 64)
	if err != nil {
		return "", 0, fmt.Errorf(`ym: recognize price value %w`, err)
	}

	// конец парсинг цены
	return YM_HOST + detailUrl, int(price * 100), nil // проверить конвертацию, 100,90 => 10090 , 100,909 => 10090, а не 10091
}
