package main

import (
	fileReader "advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BingoNumber struct {
	Number int
	Marked bool
}

type BingoBoard struct {
	Numbers [][]BingoNumber
	Won bool
}

func drawNumbers() ([]int, error) {
	inputList, err := fileReader.ReadInput()
	if err != nil {
		fmt.Println("error reading input.txt")
	} 

	numbersToCallString := strings.Split(inputList[0], ",")
	var numbersToCall []int

	// Convert list of numbers from string to int
	for _, num := range numbersToCallString {
		numAsInt, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		} 
		numbersToCall = append(numbersToCall, numAsInt)
	}

	return numbersToCall, nil
}

// getBoards populates the bingo boards from the input file
func getBoards() ([]BingoBoard, error) {
	inputList, err := fileReader.ReadInput()
	if err != nil {
		fmt.Println("error reading input.txt")
		return nil, err
	} 

	// Discard first line from input list as that has been handled already
	boardNumbersInput := inputList[1:]

	var bingoBoards []BingoBoard

	board := createNewBoard()
	space := regexp.MustCompile(`\s+`)
	currentBoardRow := 0

	for _, line := range boardNumbersInput {
		// If at end of a row, continue
		if line == "" {
			continue
		}
		// Remove any extra whitespace
		line = strings.TrimSpace(line)
		line = space.ReplaceAllString(line, " ")

		// Split line into individual numbers
		numbers := strings.Split(line, " ")

		for col, numAsString := range numbers {
			// Convert each number from string to int
			num, err := strconv.Atoi(numAsString)
			if err != nil {
				panic(err)
			}
			// Set the number in the current position
			board.Numbers[currentBoardRow][col].Number = num
		}

		// If at final row of the board, add the current board to the final list
		// and create a new board, otherwise move on to the next row
		if currentBoardRow == 4 {
			bingoBoards = append(bingoBoards, board)
			board = createNewBoard()
			currentBoardRow = 0
		} else {
			currentBoardRow++
		}
	}
	return bingoBoards, nil
}


// newBoard creates a new BingoBoard
func createNewBoard() BingoBoard {
	board := BingoBoard{}
	board.Numbers = make([][]BingoNumber, 5)
	for x := range board.Numbers {
		board.Numbers[x] = make([]BingoNumber, 5)
		for y := range board.Numbers[x] {
			board.Numbers[x][y] = BingoNumber{}
		}
	}

	return board
}

// markNumbers marks the numbers on a board if they have been drawn
func markNumbers(drawnNumbers []int, boards []BingoBoard) []BingoBoard {
	for _, number := range drawnNumbers {
		for _, board := range boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board.Numbers[i][j].Number == number {
						board.Numbers[i][j].Marked = true
					}
				}
			}
		}
	}
	return boards
}

// getWinningBoards returns any bingo boards that have won with the numbers drawn
func getWinningBoards(boards []BingoBoard) []BingoBoard {
	// Check if any row is all marked
	winningBoards := []BingoBoard{}
	for _, board := range boards {
		if board.Won {
			continue
		}
		for i := 0; i < 5; i++ {
			allMarkedInRow := true
			for j := 0; j < 5; j++ {
				if !board.Numbers[i][j].Marked {
					allMarkedInRow = false
					break
				}
			}
			if allMarkedInRow {
				winningBoards = append(winningBoards, board)
				break
			}
		}
	}

	// Check if any column is all marked
	for _, board := range boards {
		if board.Won {
			continue
		}
		for i := 0; i < 5; i++ {
			allMarkedInColumn := true
			for j := 0; j < 5; j++ {
				if !board.Numbers[j][i].Marked {
					allMarkedInColumn = false
					break
				}
			}
			if allMarkedInColumn {
				winningBoards = append(winningBoards, board)
				break
			}
		}
	}

	return winningBoards
}

// getScoreOfWinningBoard returns the sum of the numbers that have not been 
// marked on a winning board multiplied by the winning number
func getScoreOfWinningBoard(winningNumber int, winningBoard BingoBoard) int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !winningBoard.Numbers[i][j].Marked {
				score += winningBoard.Numbers[i][j].Number
			}
		}
	}
	return score * winningNumber
}

// PuzzleOne returns the answer of puzzle one of Giant Squid (day 4 of AoC).
func PuzzleOne() (int, error) {
	
	scoreOfWinningBoard := 0
	drawnNumbers, err := drawNumbers()
	if err != nil {
		fmt.Println("error retrieving the numbers drawn")
		return scoreOfWinningBoard, err
	}
	boards, err := getBoards()
	if err != nil {
		fmt.Println("error retrieving the numbers drawn")
		return scoreOfWinningBoard, err
	}

	for _, number := range drawnNumbers {
		for _, board := range boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board.Numbers[i][j].Number == number {
						board.Numbers[i][j].Marked = true
					}
				}
			}
		}

		winningBoards := getWinningBoards(boards)
		if len(winningBoards) == 0 {
			continue
		} else {
			scoreOfWinningBoard = getScoreOfWinningBoard(number, winningBoards[0])
			return scoreOfWinningBoard, nil
		}

	}
	return scoreOfWinningBoard, nil
}

func main() {
	 score, _ := PuzzleOne()
	 fmt.Println(score)
}