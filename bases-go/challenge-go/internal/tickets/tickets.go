package tickets

import (
	"errors"
	"strconv"
)

type Ticket struct {
	Id                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              int
}

// Getting the slice of tickets from ReadCSV function
var tickets = ReadCSV()

// Function that returns the total of tickets for a destination
func GetTotalTickets(destination string) (int, error) {
	total := 0
	for _, ticket := range tickets {
		if ticket.DestinationCountry == destination {
			total++
		}
	}
	if total == 0 {
		return 0, errors.New("no tickets found for that destination")
	}
	return total, nil
}

// Function that calculates how much people travel in the early morning (between 0 and 6 am), morning (between 6 and 12 pm), afternoon (between 12 and 19 pm) and night (between 19 and 0 pm), including the hours and minutes. Returns s slice of string and slice of int and error if the time is invalid
func GetCountByPeriod() ([]string, []int, error) {
	// Use a slice to store the time periods and the total of tickets for each period
	timePeriods := []string{"early morning", "morning", "afternoon", "night"}
	// Declare a slice to store the total of tickets for each period
	var totals []int
	// Declare a variable to store the total
	var totalA, totalB, totalC, totalD int
	// Loop through the tickets slice
	for _, ticket := range tickets {
		// Split the flight time into hours and cast it to int
		flightHours, _ := strconv.Atoi(ticket.FlightTime[:2])
		// Switch case to check the time period
		switch {
		case flightHours >= 00 && flightHours < 06:
			// Increment the total
			totalA++
		case flightHours >= 06 && flightHours < 12:
			// Increment the total
			totalB++
		case flightHours >= 12 && flightHours < 19:
			// Increment the total
			totalC++
		case flightHours >= 19:
			// Increment the total
			totalD++
		default:
			return nil, nil, errors.New("invalid time")

		}
	}
	totals = append(totals, totalA, totalB, totalC, totalD)

	return timePeriods, totals, nil
}

// Function that calcules the average of people that travel to a destination. Returns an int and error if the destination is invalid
func AverageDestination(destination string) (float64, error) {
	// Declare a variable to store the total of tickets in csv file
	var total int = len(tickets)
	// Declare a variable to store the total of tickets for the destination
	var totalDestination int
	// Loop through the tickets slice
	for _, ticket := range tickets {
		// Check if the destination is the same as the destination parameter
		if ticket.DestinationCountry == destination {
			// Increment the total of tickets for the destination
			totalDestination++
		}
	}
	// Check if the total of tickets for the destination is equal than 0
	if totalDestination == 0 {
		return 0, errors.New("no tickets found for that destination")
	}
	// Calculate the average of tickets for the destination in float64
	average := (float64(totalDestination) / float64(total)) * 100.0
	// Return the average
	return average, nil

}
