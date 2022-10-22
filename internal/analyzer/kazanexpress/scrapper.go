package kazanexpress

import (
	"fmt"
	"strings"
	"time"

	"github.com/ducktyst/bar_recomend/internal/analyzer/common"
	"github.com/tebeka/selenium"
	// https://github.com/SeleniumHQ/docker-selenium
)

var KAZAN_SEARCH_URL = "https://kazanexpress.ru/search?query={{.search_text}}&sorting=price&ordering=ascending"

// var SEARCH_URL = "https://pkg.go.dev/github.com/mediabuyerbot/go-webdriver"
var SEARCH_URL = KAZAN_SEARCH_URL

var OZON_SEARCH_URL = "https://www.ozon.ru/search/?from_global=true&text={{.search_text}}"

var port = 4445             // TODO: to env
var browserName = "firefox" // "chrome"

func ParseWithSelenium() {
	//2. Call browser instance
	//Set browser compatibility. We set the browser name to chrome
	caps := selenium.Capabilities{
		"browserName": browserName,
	}

	fmt.Println("New remote")
	//Call browser urlPrefix: test reference: defaulturlprefix ='http://127.0.0.1:4444/wd/hub'
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	// if err := wd.SetImplicitWaitTimeout(5 * time.Second); err != nil {
	// 	panic(fmt.Errorf("Could not set implicit wait timeout: %w", err))
	// }
	wd.SetPageLoadTimeout(20 * time.Second)

	//Delay exiting chrome
	defer wd.Quit()

	fmt.Println("GenerateSearchUrl")
	// 3. radio, checkbox and select box operation (functions to be improved, https://github.com/tebeka/selenium/issues/141)
	url := common.GenerateSearchUrl(SEARCH_URL, "полотенце")
	// fmt.Println("wd.Get(", url, ")")
	if err := wd.Get(url); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// process webpage

	// fmt.Println(wd.CurrentURL())

	// fmt.Println(wd.FindElement)
	// fmt.Println(wd.ActiveElement())
	// fmt.Println(wd.Status())
	time.Sleep(time.Second * 10)
	wes, err := wd.FindElements(selenium.ByCSSSelector, `[data-test-id="text__price"]`)
	// wes, err := wd.FindElements(selenium.ByID, `category-products`)
	if err != nil {
		panic(err)
	}
	fmt.Println("len(wes)", len(wes))
	for _, we := range wes[:10] {
		fmt.Println("Element")
		text, err := we.Text()
		if err != nil {
			fmt.Println(fmt.Errorf(`get text from [data-test-id="text__price"] err %w`, err))
			continue
		}
		fmt.Println(strings.Split(text, " ")[0], "= цена")

	}

	lowestPriceWe, err := wd.FindElement(selenium.ByCSSSelector, `[data-test-id="text__price"]`)
	// wes, err := wd.FindElements(selenium.ByID, `category-products`)
	if err != nil {
		panic(err)
	}
	price, err := lowestPriceWe.Text()
	if err != nil {
		fmt.Println(fmt.Errorf(`get text from [data-test-id="text__price"] err %w`, err))
	}
	fmt.Println(strings.Split(price, " ")[0], "= минимальная цена")
	// we.Click()

	fmt.Println("success")
}

// примеры безинтерфейсного браузера https://github.com/chromedp/examples
