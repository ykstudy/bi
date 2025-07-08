package main

import (
	"context"
	"encoding/json"
	"example/data"
	"example/utils"
	"fmt"
)

func Kline() {
	symbol := "KERNELUSDT"
	for _, interval := range utils.IntervalList {
		klines, err := client.NewKlinesService().
			Symbol(symbol).
			Interval(string(interval)).
			Do(context.Background())

		var d []string
		for _, kline := range klines {
			b, _ := json.Marshal(kline)
			d = append(d, string(b))
		}

		// 写入到文件
		filePath := fmt.Sprintf("data/klines/%s/%s_%s.txt", symbol, utils.IntervalPathMap[interval], interval)
		err = data.WriteToFile(filePath, d)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
