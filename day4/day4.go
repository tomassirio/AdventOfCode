package day4

import (
	"strconv"
	"strings"
)

type Day struct {}

func NewDay() Day{
	return Day{}
}

type board [][]int

func (d Day) Solution1(input []string) int {

	bingoValues := getSelectedNumbers(input)
	boards := getBoards(getBoardsInput(input))

	var winnerBoard board
	var boardSum int
	var winnerNumber int

	for _, number := range bingoValues {
		boards, winnerBoard = markAndCheckAllBoards(number, boards)
		if winnerBoard != nil {
			boardSum = calculateMatrixSum(winnerBoard)
			winnerNumber = number
			break
		}
	}

	return calculateFinalResult(boardSum, winnerNumber)
}

func (d Day) Solution2(input []string) int {
	bingoValues := getSelectedNumbers(input)
	boards := getBoards(getBoardsInput(input))

	//If it has a Bingo
	//		----> Subtract that board
	//If boards.len == 1
	//		Return boards[0]
	var loserBoard board
	var loserNumber int

	for _, number := range bingoValues {
		var newLoser board
		boards, newLoser = markAndCheckAllBoards(number, boards)
		boards = removeAllWinnerBoards(boards)
		if len(boards) == 0 {
			loserBoard = newLoser
			loserNumber = number
			break
		}
	}
	boardSum := calculateMatrixSum(loserBoard)
	return calculateFinalResult(boardSum, loserNumber)
}

func getBoardsInput(input []string) []string {
	return deleteEmptyLines(input[2:])
}

func deleteEmptyLines (s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func parseStringToIntSlice(line string) []int {
	line = removeFirstEmptySpace(removeDoubleSpaces(line))
	valuesSeparated := strings.Split(line, " ")
	var solution []int
	for _, value := range valuesSeparated {
		intValue, _ := strconv.Atoi(value)
		solution = append(solution, intValue)
	}
	return solution
}

func removeDoubleSpaces(line string) string{
	return strings.ReplaceAll(line, "  ", " ")
}

func removeFirstEmptySpace(line string) string{
	result := line
	if line[0:1] == " " {
		result = result[1:]
	}
	return result
}

func getSelectedNumbers(input []string) []int{
	bingoValues := strings.Split(input[0], ",")
	var intArray []int

	for _, value := range bingoValues {
		intValue, _ := strconv.Atoi(value)
		intArray = append(intArray, intValue)
	}
	return intArray
}

func getBoards(boardsInput []string) []board {
	var boards []board
	lines := 0
	var tempBoard board
	for _, line := range boardsInput {
		var tempLine []int
		tempLine = parseStringToIntSlice(line)
		tempBoard = append(tempBoard, tempLine)
		lines++
		if lines > 4 {
			lines = 0
			boards = append(boards, tempBoard)
			tempBoard = nil
		}
	}
	return boards
}

func checkHorizontally(board board) bool{
	result := false
	for _, line := range board {
		bingoLine := true
		for _, value := range line {
			if value != -1 {
				bingoLine = false
			}
		}
		if bingoLine { result = true}
	}
	return result
}

func checkVertically(board board) bool{
	return checkHorizontally(transpose(board))
}

func transpose(board board) board {
	xl := len(board[0])
	yl := len(board)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = board[j][i]
		}
	}

	return result
}

func markAndCheckBoard(number int, board board) (board, bool){
	for i, line := range board {
		for j, position := range line {
			if position == number {
				board[i][j] = -1
			}
		}
	}

	return board, isBingo(board)
}

func markAndCheckAllBoards(number int, boards []board) ([]board, board) {
	var winnerBoard board = nil
	for i, boardValue := range boards {
		winner := false
		boards[i], winner = markAndCheckBoard(number, boardValue)
		if winner {
			winnerBoard = boards[i]
		}
	}
	return boards, winnerBoard
}

func calculateMatrixSum(board board) int {
	result := 0
	for _, line := range board {
		for _, value := range line {
			if value != -1 {
				result += value
			}
		}
	}
	return result
}

func calculateFinalResult(matrixSum, winnerNumber int) int {
	return matrixSum * winnerNumber
}

func removeBoard(index int, boards []board) []board {
	return append(boards[:index], boards[index+1:]...)
}

func removeAllWinnerBoards(boards []board) []board {
	tempBoards := make([]board, len(boards))
	copy(tempBoards, boards)
	erasedNumbers := 0
	for i, boardValue := range tempBoards {

		if isBingo(boardValue) {
			tempBoards = removeBoard(i-erasedNumbers, tempBoards)
			erasedNumbers ++
		}
	}
	return tempBoards
}

func isBingo(board board) bool {
	h := checkHorizontally(board)
	v := checkVertically(board)

	return h || v
}