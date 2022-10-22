package main

// "github.com/ducktyst/bar_recomend/internal/analyzer/ozon"

import (
	"fmt"
	"time"
)

// func main() {
// 	scanner := barcode.NewScanner()

// 	img := barcode.NewImage("../../static/barcode_good_quality.jpg")

// 	symbols, err := scanner.ScanImage(img)
// 	if err != nil {
// 		logrus.Error()
// 	}

// 	logrus.Log(symbols)
// }

//https://stackoverflow.com/questions/70482959/colly-difference-between-request-visit-and-collector-visit
var path = "/Users/aleksej/Projects/bar_recommend/static/barcode_good_quality.jpg"

func main() {
	for i := range []int{1, 2, 3, 4, 5} {
		fmt.Print("\r", 5-i)
		time.Sleep(time.Second)
	}
	fmt.Print("\r")

	articul = scanner.ScanBarCode()
	common.GetMinPriceFrom(common.KazanExpress, articul)
}

func initCommands() {
	// restart container by cron, по причине его нестабильности
	// очередь запросов на время перезапуска докера
	// оркестровка между запущенными интансами докера + селениума
}
