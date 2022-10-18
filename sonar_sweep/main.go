package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// read_Input reads in a text file and populates a struct
// with the numbers in the file.
func readInput() ([]int, error) {
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	// Initialize slice to store input
	var inputList []int

	// Create scanner with the file read in
    scanner := bufio.NewScanner(file)

	// Store each line (number) in the input file in the slice
    for scanner.Scan() {
		numAsString := scanner.Text()
		num, err := strconv.Atoi(numAsString)
		if err != nil {
			return nil, err
		}
		inputList = append(inputList, num)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	fmt.Println(len(inputList))
	return inputList, nil
}

// Puzzle_One returns the result of puzzle one of Sonar Sweep (day 1 of AoC)
// It returns the number of entries in the input list that are greater than the 
// previous entry.
func Puzzle_One() (int, error) {
	inputList, err := readInput()
	if err != nil {
		return 0, err
	}

	result := 0
	temp := inputList[0]

	for _, num := range inputList {
		if num > temp {
			result++
		}
		temp = num
	}
	return result, nil
}

func main() {
	x, _ := Puzzle_One()
	fmt.Println(x)
}