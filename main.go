package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return
func nakedReturn(name string) (length int, uppercase string) {
	defer fmt.Println("I'm doen") //함수가 끝났을 때 무엇을 할지 선언을 해 줄 수 있는 함수
	length = 1                    //위에 선언문에서 이미 생성되어 있기 때문에 선언하면 안되고 값을 대입해야 함
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number) //range는 인덱스와 값을 리턴함
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

	return 1
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	hyunsoo := person{name: "hyunsoo", age: 26, favFood: favFood}
	fmt.Println(hyunsoo)

	nico := map[string]string{"name": "nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	}
	fmt.Println(nico)

	a := 2
	b := &a //a의 메모리 주소에 연결되어 있기 때문에 a가 변경되면 같이 변경됨
	a = 5
	*b = 20
	fmt.Println(a)
	//git test

	names := []string{"nico", "lynn", "dal"}
	//names[3] = "alala"
	//names[4] = "alala"
	//names = append(names, "flynn") //modify가 아니고 새로운 slice를 리턴함
	//slice는 Array와 같은 역할을 하는 애들이지만 length 제한이 없음
	fmt.Println(names)
}
