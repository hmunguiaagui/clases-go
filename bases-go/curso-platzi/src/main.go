package main

import "fmt"

func main() {
	// Declare a constant
	const pi float32 = 3.14
	fmt.Println("pi", pi)
	// Declare a variable
	var base float64 = 10.1 // Explicit declaration
	fmt.Println("base", base)
	height := 8.2 // Implicit declaration
	fmt.Println("height", height)
	var area = base * height // Implicit declaration, default type is float64
	fmt.Printf("area %f \n-----------\n", area)

	// Arithmetic operators
	numA := 10
	numB := 5
	fmt.Println("numA =", numA)
	fmt.Println("numB =", numB)
	fmt.Println("numA + numB =", numA+numB)
	fmt.Println("numA - numB =", numA-numB)
	fmt.Println("numA * numB =", numA*numB)
	fmt.Println("numA / numB =", numA/numB)
	fmt.Println("numA % numB =", numA%numB)
}
