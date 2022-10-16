package main

import (
	"github.com/ducktyst/bar_recomend/internal/analyzer/kazanexpress"
	// "github.com/ducktyst/bar_recomend/internal/analyzer/ozon"
	// "github.com/ducktyst/bar_recomend/internal/analyzer/ym"
)

// import (
// 	"github.com/bieber/barcode"
// 	"github.com/sirupsen/logrus"
// )

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
	// fin, err := os.Open(path)
	// if err != nil {
	// 	logrus.Errorf("error open file", err)
	// 	panic(err)
	// }
	// barcode, err := barcode_scanner.ScanBarCode(fin)
	// if err != nil {
	// 	panic(err)
	// }
	// product, err := barcodefinder.GetProductInfo(barcode)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(product.BarCode, product.Name, product.Prices)
	// analyzer.CollyExample()
	// analyzer.ScapPort()
	// common.ExampleScrape()
	// ym.Scrap()
	// ozon.Scrap()
	kazanexpress.Scrap()
}
