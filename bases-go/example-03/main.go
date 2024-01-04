package main

import "fmt"

// Functions excersice
func main() {

	const (
		salary1 = 50000.0
		salary2 = 50001.0
		salary3 = 100000.0
		salary4 = 100001.0
	)

	e1 := e1Tax(50000)
	e2 := e1Tax(50001)
	e3 := e1Tax(100000)
	e4 := e1Tax(100001)
	fmt.Printf("The tax of employee 1 with salary %.2f is %.2f\n", salary1, e1)
	fmt.Printf("The tax of employee 2 with salary %.2f is %.2f\n", salary2, e2)
	fmt.Printf("The tax of employee 3 with salary %.2f is %.2f\n", salary3, e3)
	fmt.Printf("The tax of employee 4 with salary %.2f is %.2f\n", salary4, e4)
}

// Function that returns the tax of a salary employee
func e1Tax(salary float64) (tax float64) {
	switch {
	case salary > 50000 && salary <= 100000:
		tax = salary * 0.17
		return tax
	case salary > 100000:
		tax = salary * 0.27
		return tax
	default:
		return 0
	}
}
