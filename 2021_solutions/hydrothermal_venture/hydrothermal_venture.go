package main

import (
	fileReader "advent-of-code/utils"
	"fmt"
	"strings"
	"strconv"
)


// getCoordinates returns four points to be used for coordinates
func getCoordinates(line string) (int, int, int, int, error) {
	coordinates := strings.Split(line, " -> ")
	a := strings.Split(coordinates[0], ",")
	b := strings.Split(coordinates[1], ",")

	x1, err := strconv.Atoi(a[0])
	if err != nil {
		fmt.Println("There was an error converting the coordinate string to an int")
		return 0, 0, 0, 0, err
	}
	y1, err := strconv.Atoi(a[1])
	if err != nil {
		fmt.Println("There was an error converting the coordinate string to an int")
		return 0, 0, 0, 0, err
	}
	x2, err := strconv.Atoi(b[0])
	if err != nil {
		fmt.Println("There was an error converting the coordinate string to an int")
		return 0, 0, 0, 0, err
	}
	y2, err := strconv.Atoi(b[1])
	if err != nil {
		fmt.Println("There was an error converting the coordinate string to an int")
		return 0, 0, 0, 0, err
	}

	return x1, y1, x2, y2, nil
}

// getCoordinatesBetweenPoints returns all the coordinates between two given points on a line
// i.e. if two coordinates are given with the same x value, the function returns all the 
// y coordinates between these points on the line (and vice versa if two y coordinates given).
func getCoordinatesBetweenPoints(p1 int, p2 int) []int {
	var pointsBetween []int

	if p1 > p2 {
		for ; p2 <= p1; p2++ {
			pointsBetween = append(pointsBetween, p2)
		}
		return pointsBetween
	} else {
		for ; p1 <= p2; p1++ {
			pointsBetween = append(pointsBetween, p1)
		}
		return pointsBetween
	}
}

// isHorizontalLine returns true if two x coordinates given are the same, 
// and thus make a horizontal line.
func isHorizontalLine(x1 int, x2 int) bool {
	if x1 != x2 {
		return false
	}
	return true
}

// isVerticalLine returns true if two y coordinates given are the same, 
// and thus make a vertical line
func isVerticalLine(y1 int, y2 int) bool {
	if y1 != y2 {
		return false
	}
	return true
}

// PuzzleOne returns the result of puzzle one of HydrothermalVenture (day 5 of AoC).
// It returns the number of points from the input filewhere at least two vertical/horizontal
// lines overlap.
func PuzzleOne() (int, error) {
	inputList, _ := fileReader.ReadInput()

	var x1, y1, x2, y2 int
	var err error
	var grid [1000][10000]uint8

	for _, line := range inputList {
		x1, y1, x2, y2, err = getCoordinates(line)
		if err != nil {
			fmt.Println("There was an error converting the coordinate string to an int")
			return 0, err
		}

		// For every point covered by a horizontal/vertical line, 
		// add one to this point in the grid
		if isHorizontalLine(x1, x2) {
			coordinatesBetweenPoints := getCoordinatesBetweenPoints(y1, y2)
			for _, point := range coordinatesBetweenPoints {
				grid[x1][point]++
			}
		} else if isVerticalLine(y1, y2) {
			coordinatesBetweenPoints := getCoordinatesBetweenPoints(x1, x2)
			for _, point := range coordinatesBetweenPoints {
				grid[point][y1]++
			}
		}

	}

	// Calculates the number of points where at least 2 lines are overlapping it
	totalPoints := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] >= 2 {
				totalPoints++
			}
		}
	}
		return totalPoints, nil
}

func main() {
	sol1, err := PuzzleOne()
	if err != nil {
		fmt.Println("There was an error getting the solution to puzzle 1")
	}
	fmt.Println(sol1)
}