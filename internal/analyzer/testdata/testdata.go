package testdata

import (
	"path/filepath"
	"runtime"
)

func KazanExpressPath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir, _ := filepath.Split(filename)
	fpath := filepath.Join(dir, "kazanexpress", "search_polotentse", "полотенце - Результаты поиска.html")
	// go_zbar/internal/analyzer/testdata/kazanexpress/search_polotentse/полотенце - Результаты поиска.html

	// wd, _ := os.Getwd()

	// fmt.Println(filename)
	// fmt.Println(wd)
	return fpath
}
