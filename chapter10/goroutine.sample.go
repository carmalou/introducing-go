package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, ":", n)
	}
}
