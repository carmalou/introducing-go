package main

import "fmt"

func main() {
	fmt.Print("Enter a distance in feet: ")

	var input float64
	fmt.Scanf("%f", &input)

	meters := input * 0.348

	fmt.Println(input, "feet is", meters, "in meters.")
}
