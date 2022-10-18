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

// get_Gamma_And_Epsilon_Rate gets the gamma and epsilon rate 
// as decimal numbers from the input list
func get_Gamma_And_Epsilon_Rate() (int64, int64, error) {
	inputList, err := readInput()
	if err != nil {
		fmt.Println("error")
	}

	// Create array to store the number of 0's present in the list at each index
	var zeroCount [12]int

	gammaString := ""
	epsilonString := ""

	for _, line := range inputList {
		for i, bit := range line {
			if bit == '0' {
				zeroCount[i] += 1
			}
		}
	}

	for _, count := range zeroCount {
		if count > len(inputList)/2 {
			gammaString += "0"
			epsilonString += "1"
		} else {
			gammaString += "1"
			epsilonString += "0"
		}
	}

	// Convert binary numbers to decimals
	gamma, err := strconv.ParseInt(gammaString, 2, 64) 
	if err != nil {
		return 0, 0, err
	}
	epsilon, err := strconv.ParseInt(epsilonString, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	return gamma, epsilon, nil
}


func Puzzle_One() (int64, error){
	gamma, epsilon, err := get_Gamma_And_Epsilon_Rate()
	if err != nil {
		return 0, err
	} 
	return gamma * epsilon, nil
}


func main() {
	result1, err := Puzzle_One()
	if err != nil {
		fmt.Println("There was an error getting the solution to puzzle 1")
	}
	fmt.Println(result1)
}
