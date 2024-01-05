package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Function that reads a CSV file and returns a slice of tickets
func ReadCSV() []Ticket {
	path := "tickets.csv"
	// Open CSV file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// Read file
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = -1
	// Read all records
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	// Declare a ticket slice
	var tickets []Ticket
	// Loop through records
	for _, each := range rawCSVdata {
		id, _ := strconv.Atoi(each[0])
		price, _ := strconv.Atoi(each[5])
		// Append tickets to the slice
		tickets = append(tickets, Ticket{
			Id:                 id,
			Name:               each[1],
			Email:              each[2],
			DestinationCountry: each[3],
			FlightTime:         each[4],
			Price:              price,
		})
	}
	return tickets

}
