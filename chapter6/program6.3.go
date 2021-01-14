package main

import "fmt"

func main() {
	nextOdd := generateOddNumbers()

	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}

func generateOddNumbers() func() int {
	var i int = 1

	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}
