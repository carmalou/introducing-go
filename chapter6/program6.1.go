package main

import "fmt"

func main() {
	fmt.Print("Tell me a number, and I'll tell you what half of it is, and if it is even: ")

	var input int
	fmt.Scanf("%d", &input)

	num, even := half(input)

	fmt.Println("Half:", num, "Even:", even)
}

func half(num int) (int, bool) {
	if num%2 == 0 {
		even := num / 2
		return even, true
	}

	odd := num / 2
	return odd, false
}
