package main

import (
	"fmt"
	"github.com/kumakuma34/learningGo/account"
)

func main() {
	account := account.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdrawl(20)
	if err != nil {
		fmt.Println(err)
		//log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
