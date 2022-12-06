package main

import (
	"flag"
	"fmt"

	"github.com/AlekSi/pointer"
)

var barcodeFlag = flag.String("b", "", "штрихкод по которому необходимо наполнить базу")
var countFlag = flag.Int("n", 10, "количество элементов для заполнения. \n значение >10 мошет привести к ошибкам из-за пагинации на сайтах")

func main() {
	flag.Parse()

	fmt.Println("barcodeFlag has value ", *barcodeFlag)
	if pointer.GetString(barcodeFlag) == "" {
		panic("flag -b must be filled")
	}
	fmt.Println("countFlag has value ", *countFlag)
}

func parseYM(barcode string, count int) {

}
