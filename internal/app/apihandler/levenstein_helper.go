package apihandler

import leven "github.com/ducktyst/bar_recomend/internal/barcode/levenstein"

type product struct {
	ID      int
	Articul string
}
type productID int
type levensteinDist int

func allDistancesMap(products []product) map[productID]map[productID]levensteinDist {

	var productLevisteinMap = make(map[productID]map[productID]levensteinDist, len(products))

	for i := range products {
		currentProduct := products[i]
		for j := range products {
			compareProduct := products[j]
			if currentProduct.ID == compareProduct.ID {
				productLevisteinMap[productID(currentProduct.ID)][productID(compareProduct.ID)] = 0
				continue
			}
			dist := leven.Distance(currentProduct.Articul, compareProduct.Articul)
			productLevisteinMap[productID(currentProduct.ID)][productID(compareProduct.ID)] = levensteinDist(dist)
		}
	}

	return productLevisteinMap
}
