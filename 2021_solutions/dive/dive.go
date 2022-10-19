package main

import (
	fileReader "advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

// get_Horizontal_Position_And_Depth_One calculates the final horizontal
// position and depth of the submarine using the input list
func get_Horizontal_Position_And_Depth(considerAim bool) (int, int, error) {
	inputList, err := fileReader.ReadInput()
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
		sections := strings.Split(input, " ")
		direction = sections[i]
		valueAsString := sections[i+1]
		value, _ = strconv.Atoi(valueAsString)

		// Cases for each puzzle
		if !considerAim {
			if direction == "forward" {
				horizontalPosition = horizontalPosition + value
			} else if direction == "down" {
				depth = depth + value
			} else {
				depth = depth - value
			}
		} else {
			if direction == "down" {
				aim = aim + value
			} else if direction == "up" {
				aim = aim - value
			} else {
				horizontalPosition = horizontalPosition + value
				depth = depth + (aim * value)
			}
		}

		// Reset sections for next line of input
		sections = nil
	}
	return horizontalPosition, depth, nil
}

// Puzzle_One returns the answer to puzzle one of dive (day 2 of AoC),
// it multiplies the horizontal position and depth to return the product
func Puzzle_One() (int, error) {
	horizontalPosition, depth, err := get_Horizontal_Position_And_Depth(false)
	if err != nil {
		return 0, err
	}
	return horizontalPosition * depth, nil
}

// Puzzle_Two returns the answer to puzzle two of dive (day 2 of AoC),
// it multiplies the horizontal position and depth to return the product
func Puzzle_Two() (int, error) {
	horizontalPosition, depth, err := get_Horizontal_Position_And_Depth(true)
	if err != nil {
		return 0, err
	}
	return horizontalPosition * depth, nil
}

// func main() {
// 	result1, err := Puzzle_One()
// 	if err != nil {
// 		fmt.Println("There was an error getting the solution to puzzle 1")
// 	}
// 	fmt.Println(result1)

// 	result2, err := Puzzle_Two()
// 	if err != nil {
// 		fmt.Println("There was an error getting the solution to puzzle 2")
// 	}
// 	fmt.Println(result2)
// }
