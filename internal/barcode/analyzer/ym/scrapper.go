package ym

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
)

const (
	YM_HOST          = "https://market.yandex.ru/"
	YM_SEARCH_URL    = "https://market.yandex.ru/search?cvredirect=2&text={{.search_text}}"
	YM_LIST_URL      = "https://market.yandex.ru/catalog--produkty/54434/list?srnum=1002&was_redir=1&rt=9&rs=eJwzEg1grGLh-HGSdRYj14X9F7ZebADiNgBj9QrL&text={{.search_text}}&hid=91307&how=aprice&allowCollapsing=1&local-offers-first=0"
	YM_UNIVERSAL_URL = "https://market.yandex.ru/search?cvredirect=0&text=&allowCollapsing=1&local-offers-first=0"
	SEARCH_URL       = YM_UNIVERSAL_URL
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
	productCard, err := wd.FindElement(selenium.ByCSSSelector, `article[data-autotest-id="product-snippet"]`)
	if err != nil {
		return "", 0, fmt.Errorf(`ym: get product card error: %w`, err)
	}

	/*
		третий div в article с атрибутом data-autotest-id="product-snippet
		в этом диве див в котором див в котором див в котором <a href...>

	*/
	fmt.Println(productCard.Text)
	detailUrlElem, err := productCard.FindElement(selenium.ByCSSSelector, `div:nth-child(3) + div + div + div + a`)
	if err != nil {
		return "", 0, fmt.Errorf(`ym: get product card url error: %w`, err)
	}
	detailUrl, err := detailUrlElem.CSSProperty("href")
	if err != nil {
		return "", 0, fmt.Errorf(`ym: get product card element url error: %w`, err)
	}
	// конец парсинг ссылки на карточку товара

	// начало парсинг цены
	/*
		в этом  <a href...> div в котором h3 в котором нужен второй span
	*/
	lowestPriceWe, err := detailUrlElem.FindElement(selenium.ByCSSSelector, `div + h3 > span:nth-child(2)`)
	if err != nil {
		return "", 0, fmt.Errorf(`ym: get lowestPriceWe err %w`, err)
	}
	// в span с ценой обнапжуены только целые числа без копеек
	priceVal, err := lowestPriceWe.Text()
	if err != nil {
		return "", 0, fmt.Errorf(`ym: get price value error %w`, err)
	}
	price, err := strconv.ParseInt(priceVal, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf(`ym: recognize price value %w`, err)
	}
	resPrice := int(price * 100)

	// конец парсинг цены
	return YM_HOST + detailUrl, resPrice, nil
}
