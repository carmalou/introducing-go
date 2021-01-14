package main

import "fmt"

func main() {
	fmt.Print("Enter a temperature in fahrenheit: ")

	var input float64
	fmt.Scanf("%f", &input)

	var celsius float64 = ((input - 32.0) * (5.0 / 9.0))

	fmt.Println(input, "in fahrenheit is", celsius, "in celsius")
}
