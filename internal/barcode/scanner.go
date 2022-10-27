package barcode

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/bieber/barcode"
	"github.com/sirupsen/logrus"
)

func ScanBarCodeFile(fin *os.File) (string, error) {
	defer fin.Close()
	src, err := jpeg.Decode(fin)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	img := barcode.NewImage(src)
	scanner := barcode.NewScanner().SetEnabledAll(true)

	symbols, _ := scanner.ScanImage(img)

	if len(symbols) != 1 {
		return "", fmt.Errorf("image must contain exactly one barcode")

	}

	return symbols[0].Data, nil
}
