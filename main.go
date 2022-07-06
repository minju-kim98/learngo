package main

import (
	"fmt"

	"github.com/minju-kim98/learngo/banking"
)

func main() {
	account := banking.Account{Owner: "Minju", Balance: 10000}
	fmt.Println(account)
}
