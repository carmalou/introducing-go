package main

import "fmt"

func main() {

	var smallestInt int = variadicFunc(3, 4, 5, 2, 1)

	fmt.Println("The smallest number is:", smallestInt)
}

func variadicFunc(args ...int) (smallest int) {
	smallest = args[0]

	for i := 0; i < len(args); i++ {
		if smallest > args[i] {
			smallest = args[i]
		}
	}

	return
}
