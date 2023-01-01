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

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
	//totalLength, upperName := lenAndUpper("hyunsoo")
	//fmt.Println(totalLength, upperName)
	//fmt.Println(multiply(2, 2))
	//repeatMe("sfd", "dsfd", "dsf", "sdf", "sfd")
}
