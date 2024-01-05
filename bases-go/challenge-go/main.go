package main

import (
	"clases-go/bases-go/challenge-go/desafio-go-bases/internal/tickets"
	"fmt"
)

func main() {
	// Get the total of tickets for a destination
	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the total tickets that travel to Brazil are:", total)
	}
	// Get the total of tickets that travel in periods
	timePeriods, totals, err := tickets.GetCountByPeriod()
	if err != nil {
		fmt.Println(err)
	} else {
		// Print the time periods and the totals
		for i, timePeriod := range timePeriods {
			fmt.Println("the total tickets that travel in", timePeriod, "are:", totals[i])
		}
	}
	// Get the average travel to a destination
	average, err := tickets.AverageDestination("China")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the average travel to China is:", average)
	}
}
