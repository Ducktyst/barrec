package main

import (
	"fmt"
	"time"

	"github.com/ducktyst/bar_recomend/internal/api/handler"
)

var path = `/Users/aleksej/Projects/bar_recommend/static/2022-10-22 21.50.32.jpg`

func main() {
	countdown()

	recommends, err := handler.HandleRecommendationsByImg(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range recommends {
		fmt.Printf("\nминимальная цена = %#v ", r)
	}

}

func initCommands() {
	// restart container by cron, по причине его нестабильности
	// очередь запросов на время перезапуска докера
	// оркестровка между запущенными интансами докера + селениума
}

func countdown() {
	for i := range []int{1, 2, 3, 4, 5} {
		fmt.Print("\r", 5-i)
		time.Sleep(time.Second)
	}
	fmt.Print("\r")
}
