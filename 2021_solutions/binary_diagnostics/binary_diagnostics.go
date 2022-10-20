package main

import (
	fileReader "advent-of-code/utils"
	"fmt"
	"strconv"
)

// mostAndLeastCommonBit returns the most and least common bit in index i
// of all lines in the input file
func mostAndLeastCommonBit(inputList []string, i int) (uint8, uint8) {
	zeroCount := 0
	oneCount := 0

	for _, line := range inputList {
		if line[i] == '0' {
			zeroCount++
		} else {
			oneCount++
		}
	}
	if zeroCount > oneCount {
		return '0', '1'
	}
	return '1', '0'
}

// convertBinaryStringToDecimal converts a binary string to a decimal
// number to be used in multiplication
func convertBinaryStringToDecimal(binaryString string) (int64, error) {
	decimalNumber, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		return 0, err
	}
	return decimalNumber, nil
}

// getGammaAndEpsilonRate gets the gamma and epsilon rate
// as decimal numbers from the input list
func getGammaAndEpsilonRate() (int64, int64, error) {
	inputList, err := fileReader.ReadInput()
	if err != nil {
		fmt.Println("error")
	}

	gammaString := ""
	epsilonString := ""

	for i := 0; i < len(inputList[0]); i++ {
		mostCommon, leastCommon := mostAndLeastCommonBit(inputList, i)
		gammaString += string(mostCommon)
		epsilonString += string(leastCommon)
	}

	// Convert binary string to decimal number to be used to answer puzzle
	gamma, err := convertBinaryStringToDecimal(gammaString)
	if err != nil {
		return 0, 0, err
	}
	epsilon, err := convertBinaryStringToDecimal(epsilonString)
	if err != nil {
		return 0, 0, err
	}
	//fmt.Println(gamma, " ", epsilon)
	return gamma, epsilon, nil
}

// get_OxygenGeneratorAndCO2ScrubberRating gets the oxygen generator and co2 scrubber ratings
// from the input list
func get_OxygenGeneratorAndCO2ScrubberRating(inputList []string, i int, useMostCommmon bool) string {
	// Base case for recursion
	if len(inputList) == 1 {
		return inputList[0]
	}

	mostCommon, leastCommon := mostAndLeastCommonBit(inputList, i)
	comparator := leastCommon

	// Check if calculating oxygen generator rating (using most common bit if a tie)
	if useMostCommmon {
		comparator = mostCommon
	}

	var tempInputList []string

	for _, line := range inputList {
		if line[i] == comparator {
			tempInputList = append(tempInputList, line)
		}
	}
	return get_OxygenGeneratorAndCO2ScrubberRating(tempInputList, i+1, useMostCommmon)
}

// PuzzleOne returns the result of puzzle one of Binary Diagnostic (day 3 of AoC)
// It returns the product of the gamma and epsilon rates
func PuzzleOne() (int64, error) {
	gamma, epsilon, err := getGammaAndEpsilonRate()
	if err != nil {
		return 0, err
	}
	return gamma * epsilon, nil
}

// PuzzleTwo returns the result of puzzle two of Binary Diagnostic (day 3 of AoC)
// It returns the product of the oxygenGeneratorRate and CO2ScrubberRate rates
func PuzzleTwo() (int64, error) {
	inputList, err := fileReader.ReadInput()
	if err != nil {
		fmt.Println("error")
	}
	oxygenGeneratorRateBinary := get_OxygenGeneratorAndCO2ScrubberRating(inputList, 0, true)
	CO2ScrubberRateBinary := get_OxygenGeneratorAndCO2ScrubberRating(inputList, 0, false)

	// Convert binary strings to decimal numbers for multiplication
	oxygenGeneratorRate, _ := convertBinaryStringToDecimal(oxygenGeneratorRateBinary)
	CO2ScrubberRate, _ := convertBinaryStringToDecimal(CO2ScrubberRateBinary)

	return oxygenGeneratorRate * CO2ScrubberRate, nil
}

// func main() {
// 	sol1, err := PuzzleOne()
// 	if err != nil {
// 		fmt.Println("There was an error getting the solution to puzzle 1")
// 	}
// 	fmt.Println(sol1)
// 	sol2, err := PuzzleTwo()
// 	if err != nil {
// 		fmt.Println("There was an error getting the solution to puzzle 2")
// 	}
// 	fmt.Println(sol2)
// }
