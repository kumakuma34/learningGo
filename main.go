package main

import (
	"fmt"
	"github.com/kumakuma34/learningGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("")
	}
	fmt.Println(word)
	/*account := account.NewAccount("nico")
	account.Deposit(10)
	err := account.Withdrawl(20)
	if err != nil {
		fmt.Println(err)
		//log.Fatalln(err)
	}
	fmt.Println(account)*/
}
