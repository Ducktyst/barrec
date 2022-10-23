package kazanexpress

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	// https://github.com/SeleniumHQ/docker-selenium
)

var KAZAN_SEARCH_URL = "https://kazanexpress.ru/search?query={{.search_text}}&sorting=price&ordering=ascending"

// var SEARCH_URL = "https://pkg.go.dev/github.com/mediabuyerbot/go-webdriver"
var SEARCH_URL = KAZAN_SEARCH_URL

var OZON_SEARCH_URL = "https://www.ozon.ru/search/?from_global=true&text={{.search_text}}"

var port = 4445             // TODO: to env
var browserName = "firefox" // or "chrome"

// Возвращает цену с копейками
func ParseWithSelenium(url string) (int, error) {
	caps := selenium.Capabilities{
		"browserName": browserName,
	}

	//Call browser urlPrefix: test reference: defaulturlprefix ='http://127.0.0.1:4444/wd/hub'
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		return 0, fmt.Errorf(`get text from [data-test-id="text__price"] err %w`, err)
	}
	defer wd.Quit()

	wd.SetPageLoadTimeout(20 * time.Second)

	time.Sleep(time.Second * 10)

	lowestPriceWe, err := wd.FindElement(selenium.ByCSSSelector, `[data-test-id="text__price"]`)
	if err != nil {
		return 0, fmt.Errorf(`get text from [data-test-id="text__price"] err %w`, err)
	}
	price, err := lowestPriceWe.Text()
	if err != nil {
		return 0, fmt.Errorf(`get text from [data-test-id="text__price"] err %w`, err)
	}

	return strconv.Atoi(price)
}

// примеры безинтерфейсного браузера https://github.com/chromedp/examples
