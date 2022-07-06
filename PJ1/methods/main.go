package main

import (
	"fmt"
	"log"

	"github.com/minju-kim98/learngo/PJ1/methods/accounts"
)

func main() {
	account := accounts.NewAccount("Minju")
	account.Deposit(100)
	err := account.Withdraw(10)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account)
}
