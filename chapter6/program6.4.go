package main

import "fmt"

func main() {
	fmt.Print("Tell me a number, and I'll show you its preceding fibonacci sequence:")

	var input int
	fmt.Scanf("%d", &input)

	tmp := make([]int, 2)
	tmp[0] = 0
	tmp[1] = 1

	arr, err := fib(input, tmp)

	if err {
		fmt.Println("Error! Number does not exist in sequence.")
	} else {
		fmt.Println(arr)
	}

}

func fib(num int, slice []int) ([]int, bool) {
	// compare number to last index in slice
	lastIndex := slice[len(slice)-1]

	// if last number in slice is greater than num: ERROR
	if lastIndex > num {
		return slice, true
	}

	// if last number in slice matches passed in num: SUCCESS
	if lastIndex == num {
		return slice, false
	}

	// else append new number (last two numbers added together) and return fib call
	secondToLast := slice[len(slice)-2]
	finalIndex := lastIndex + secondToLast
	slice = append(slice, finalIndex)

	return fib(num, slice)
}
