package common

import (
	"fmt"
	"image/png"
	"io"
	"os"
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

	_ = os.Remove("img.jpg")
	f, err := os.Create("img.jpg")
	if err != nil {
		logrus.Errorf(`os.Create("img.jpg"): %v`, err)
		// panic(err)
	}
	defer f.Close()
	if err = png.Encode(f, src); err != nil {
		logrus.Errorf("failed to encode: %v", err)
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
