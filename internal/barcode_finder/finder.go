package barcodefinder

import "errors"

type Price struct {
	priceSrc float64
	Price    int // цена с копейками 115р.65коп. == 11565
	ShopUrl  string
	ShopName string
}
type ProductInfo struct {
	Name    string
	BarCode string
	Prices  []Price
}

var (
	ProductNotFound = errors.New("product with _selected_ barcode not found")
)

func GetProductInfo(barcode string) (ProductInfo, error) {
	return ProductInfo{}, ProductNotFound
}
