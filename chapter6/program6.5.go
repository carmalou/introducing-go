package main

import "fmt"

func main() {
	fmt.Println("Hello from line 6!")

	x, y := 6, 7

	swap(&x, &y)

	fmt.Println(x)
	fmt.Println(y)
}

func swap(xPtr *int, yPtr *int) {
	x := *xPtr
	*xPtr = *yPtr
	*yPtr = x
}
