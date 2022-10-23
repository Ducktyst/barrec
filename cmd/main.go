package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ducktyst/bar_recomend/internal/analyzer/common"
	"github.com/ducktyst/bar_recomend/internal/barcode"

	"github.com/sirupsen/logrus"
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

// var path = "/Users/aleksej/Projects/bar_recommend/static/barcode_good_quality.jpg"
var path = `/Users/aleksej/Projects/bar_recommend/static/2022-10-22 21.50.32.jpg`

func main() {
	for i := range []int{1, 2, 3, 4, 5} {
		fmt.Print("\r", 5-i)
		time.Sleep(time.Second)
	}
	fmt.Print("\r")
	fin, err := os.Open(path)
	if err != nil {
		logrus.Errorf("error open file", err)
		panic(err)
	}
	img_barcode, err := barcode.ScanBarCode(fin)
	if err != nil {
		logrus.Errorf("ScanBarCode err: %v", err)
		panic(err)
	}
	fmt.Println(img_barcode)
	articul, err := barcode.GetProductArticul(img_barcode)
	if err != nil {
		logrus.Errorf("GetProductInfo err: %v", err)
		return
	}

	price, err := common.GetPriceFrom(common.KazanExpress, articul)
	if err != nil {
		logrus.Errorf("GetPriceFrom err: %v", err)
		return
	}

	fmt.Println("минимальная цена kazanexpress = ", price/100)
}

func initCommands() {
	// restart container by cron, по причине его нестабильности
	// очередь запросов на время перезапуска докера
	// оркестровка между запущенными интансами докера + селениума
}
