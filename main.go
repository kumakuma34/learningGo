package main

import (
	"fmt"
	"github.com/kumakuma34/learningGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	fmt.Println(err2)

	/*account := account.NewAccount("nico")
	account.Deposit(10)
	err := account.Withdrawl(20)
	if err != nil {
		fmt.Println(err)
		//log.Fatalln(err)
	}
	fmt.Println(account)*/
}
