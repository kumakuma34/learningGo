package main

import (
	"fmt"
	"github.com/kumakuma34/learningGo/account"
)

func main() {
	account := account.NewAccount("nico")
	fmt.Println(account)
}
