package main

import (
	"fmt"
	"os"
)

func main() {
	readNonExistingFile("")
	fmt.Println("Execution finished")
}

// Function that reads a non-existing file and returns execute a panic
func readNonExistingFile(path string) *os.File {
	// Defer function to close file and recover panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic:", r)
		}
	}()

	file, err := os.Open("non-existing-file.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	defer file.Close()

	return file
}
