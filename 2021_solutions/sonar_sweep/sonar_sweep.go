package main

import (
	fileReader "advent-of-code/utils"
	//"fmt"
	"strconv"
)

func convertInputListToInt() ([]int, error){
	inputList, err := fileReader.ReadInput()
	if err != nil {
		return nil, err
	}

	var inputListInt []int
	for _, line := range inputList {
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		inputListInt = append(inputListInt, num)
	}

	return inputListInt, nil

}

// PuzzleOne returns the result of puzzle one of Sonar Sweep (day 1 of AoC)
// It returns the number of entries in the input list that are greater than the 
// previous entry.
func PuzzleOne() (int, error) {
	inputList, err := convertInputListToInt()
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

// PuzzleTwo returns the result of puzzle two of Sonar Sweep (day 1 of AoC)
// It returns the number of sums of a three measurement sliding window
// that are greater than the sum of the previous group of three numbers.
func PuzzleTwo() (int, error) {
	inputList, err := fileReader.ReadInput()
	if err != nil {
		return 0, err
	}

	result := 0 

	for i := 1; i + 2 < len(inputList); i++ {
		sumA := inputList[i-1] + inputList[i] + inputList[i+1]
		sumB := inputList[i] + inputList[i+1] + inputList[i+2]
		if sumB > sumA {
			result++
		}
	}
	return result, nil
}

// func main() {
// 	sol1, err := PuzzleOne()
// 	if err != nil {
// 		fmt.Println("There was an error getting the solution to puzzle 1")
// 	}
// 	sol2, _ := PuzzleTwo()
// 	if err != nil {
// 		fmt.Println("There was an error getting the solution to puzzle 2")
// 	}
// 	fmt.Println("Puzzle one solution: ", sol1, " puzzle two solution: ", sol2)
// }
