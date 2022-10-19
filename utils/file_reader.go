package utils

import (
	"bufio"
	"log"
	"os"
)

// Read_Input reads in a text file and populates a struct
// with the numbers in the file.
func ReadInput() ([]string, error) {
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	// Initialize slice to store input
	var inputList []string

	// Create scanner with the file read in
    scanner := bufio.NewScanner(file)

	// Store each line (number) in the input file in the slice
    for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}
		inputList = append(inputList, line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return inputList, nil
}