package main

import (
	"fmt"

	"github.com/minju-kim98/learngo/PJ1/banking"
)

func main() {
	account := banking.Account{Owner: "minju", Balance: 10000}

	fmt.Println(account)
}
