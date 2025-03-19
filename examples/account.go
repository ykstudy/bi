package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func Account() {

	res, err := client.NewGetAccountService().
		OmitZeroBalances(true).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	account, _ := json.Marshal(res)
	fmt.Println(string(account))
}
