package barcode

// func GetRecommendationsImg(path string) ([]common.Recommendation, error) {
// 	fin, err := os.Open(path)
// 	if err != nil {
// 		return nil, fmt.Errorf("open file error: %w", err)
// 	}
// 	img_barcode, err := ScanBarCodeFile(fin)
// 	if err != nil {
// 		return nil, fmt.Errorf("barcode.ScanBarCodeFile error: %w", err)
// 	}

// 	articul, err := GetProductArticul(img_barcode)
// 	if err != nil {
// 		return nil, fmt.Errorf(" barcode.GetProductArticul error: %w", err)
// 	}

// 	kazanexpressRecommendation, err := common.GetPriceFrom(common.KazanExpress, articul)
// 	if err != nil {
// 		return nil, fmt.Errorf("common.GetPriceFrom err: %w", err)
// 	}
// 	res := make([]common.Recommendation, 0, 5)
// 	res[0] = kazanexpressRecommendation

// 	return res, nil
// }
