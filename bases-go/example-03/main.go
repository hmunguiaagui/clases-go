package main

import "fmt"

// Functions excersice
func main() {

	// Excersice 1
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

	// Excersice 2
	average := e2average(1, 2, 3, 4, 5)
	fmt.Printf("-----\nThe average of 1, 2, 3, 4, 5 is %.2f\n", average)

	// Excersice 3
	salary := e3salary("A", 120)
	fmt.Printf("-----\nThe salary of employee type A with 120 minutes worked is %.2f\n", salary)
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

// Function that returns the averge of N numbers
func e2average(numbers ...float64) (average float64) {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	average = sum / float64(len(numbers))
	return average
}

// Function that returns the salary by employee type and minutes worked
func e3salary(employeeType string, minutesWorked int) (salary float64) {
	hoursWorked := float64(minutesWorked / 60.0)
	switch employeeType {
	case "A":
		salary = 3000.0 * hoursWorked
		salary += salary * 0.5
	case "B":
		salary = 1500.0 * hoursWorked
		salary += salary * 0.2
	case "C":
		salary = 1000.0 * hoursWorked
	default:
		salary = 0.0
	}
	return salary
}
