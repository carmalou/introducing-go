package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	i := 0

	for i < len(x) {
		if i == (len(x) - 1) {
			break
		}

		index1 := x[i]
		index2 := x[i+1]

		if index1 > index2 {

			x[i] = index2
			x[i+1] = index1
			i = 0
		} else {
			i++
		}
	}

	fmt.Println("Lowest number is:", x[0])
}
