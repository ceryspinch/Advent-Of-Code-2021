package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// read_Input reads in a text file and populates a struct
// with the numbers in the file.
func readInput() ([]string, error) {
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
		numAsString := scanner.Text()
		if err != nil {
			return nil, err
		}
		inputList = append(inputList, numAsString)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputList, nil
}

// get_Horizontal_Position_And_Depth_One calculates the final horizontal
// position and depth of the submarine using the input list
func get_Horizontal_Position_And_Depth_One() (int, int, error) {
	inputList, err := readInput()
	if err != nil {
		fmt.Println("error")
	}

	horizontalPosition := 0
	depth := 0
	direction := ""
	value := 0
	i := 0

	for _, input := range inputList {
		// Split input line into array with two items (direction and value)
		sections := strings.Fields(input)
		direction = sections[i]
		valueAsString := sections[i+1]
		value, _ = strconv.Atoi(valueAsString)

		if direction == "forward" {
			horizontalPosition = horizontalPosition + value
		} else if direction == "down" {
			depth = depth + value
		} else {
			depth = depth - value
		}

		// Reset sections for next line of input
		sections = nil
	}

	return horizontalPosition, depth, nil
}

// get_Horizontal_Position_And_Depth_Two calculates the final horizontal
// position and depth of the submarine using the input list
func get_Horizontal_Position_And_Depth_Two() (int, int, error) {
	inputList, err := readInput()
	if err != nil {
		fmt.Println("error")
	}

	horizontalPosition := 0
	depth := 0
	direction := ""
	value := 0
	i := 0
	aim := 0

	for _, input := range inputList {
		// Split input line into array with two items (direction and value)
		sections := strings.Fields(input)
		direction = sections[i]
		valueAsString := sections[i+1]
		value, _ = strconv.Atoi(valueAsString)

		if direction == "down" {
			aim = aim + value
		} else if direction == "up" {
			aim = aim - value
		} else {
			horizontalPosition = horizontalPosition + value
			depth = depth + (aim * value)
		}

		// Reset sections for next line of input
		sections = nil
	}

	return horizontalPosition, depth, nil
}

// Puzzle_One returns the answer to puzzle one of dive (day 2 of AoC),
// it multiplies the horizontal position and depth to return the product
func Puzzle_One() (int, error) {
	horizontalPosition, depth, err := get_Horizontal_Position_And_Depth_One()
	if err != nil {
		return 0, err
	}
	product := horizontalPosition * depth
	return product, nil
}

// Puzzle_Two returns the answer to puzzle one of dive (day 2 of AoC),
// it multiplies the horizontal position and depth to return the product
func Puzzle_Two() (int, error) {
	horizontalPosition, depth, err := get_Horizontal_Position_And_Depth_Two()
	if err != nil {
		return 0, err
	}
	product := horizontalPosition * depth
	return product, nil
}

func main() {
	result1, err := Puzzle_One()
	if err != nil {
		fmt.Println("There was an error getting the solution to puzzle 1")
	}
	fmt.Println(result1)

	result2, err := Puzzle_Two()
	if err != nil {
		fmt.Println("There was an error getting the solution to puzzle 2")
	}
	fmt.Println(result2)
}
