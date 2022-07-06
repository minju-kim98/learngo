package main

import (
	"fmt"

	"github.com/minju-kim98/learngo/PJ1/accounts"
)

func main() {
	account := accounts.NewAccount("Minju")
	fmt.Println(account)
}
