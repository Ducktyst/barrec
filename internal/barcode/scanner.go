package barcode

import (
	"fmt"
	"image/png"
	"io"
	"time"

	"github.com/bieber/barcode"
	"github.com/sirupsen/logrus"
)

// func ScanBarCodeFile(fin *os.File) (string, error) {
func ScanBarCodeFile(fin io.ReadCloser) (string, error) {
	defer fin.Close()
	src, err := png.Decode(fin)
	if err != nil {
		logrus.Errorf("png.Decode(fin) err: %v", err)
		return "", fmt.Errorf("cant decode image")
	}

	img := barcode.NewImage(src)
	scanner := barcode.NewScanner().SetEnabledAll(true)

	logrus.Info(time.Now().Format(time.RFC3339), " scanner.ScanImage start ", img)
	barcodes, _ := scanner.ScanImage(img)
	if len(barcodes) != 1 {
		err = fmt.Errorf("image must contain exactly one barcode, got %v", len(barcodes))
		logrus.Error(err)
		return "", err
	}

	return barcodes[0].Data, nil
}
