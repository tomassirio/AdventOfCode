package day5

import (
	"strconv"
	"strings"
)

type Day struct {}

func NewDay() Day{
	return Day{}
}

type movement struct {
	src, dst coordinate
}

type coordinate struct {
	x, y int
}

type mapMatrix struct {
	values [][]int
	size int
}

func (d Day) Solution1(input []string) int {
	movements := parseEntryToMovement(input)
	matrix := createMatrix(1000)

	matrix = markMatrix(movements, matrix)

	return countResult(matrix)
}

func (d Day) Solution2(input []string) int {
	return 0
}

func parseEntryToMovement(input []string) []movement{
	var movements []movement
	for _, entry := range input {
		splitEntries := strings.Split(entry, " -> ")
		src := getCoordinate(splitEntries[0])
		dst := getCoordinate(splitEntries[1])
		mov := movement{src, dst}
		movements = append(movements, mov)
	}
	return movements
}

func getCoordinate(input string) coordinate{
	inputSplit := strings.Split(input, ",")
	x, _ := strconv.Atoi(inputSplit[0])
	y, _ := strconv.Atoi(inputSplit[1])
	return coordinate{x, y}
}

func createMatrix(size int) mapMatrix {
	return fillMatrix(mapMatrix{make([][]int, size), size})
}

func fillMatrix(matrix mapMatrix) mapMatrix {
	for i := range matrix.values {
		matrix.values[i] = make([]int, matrix.size)
		for j := range matrix.values[i] {
			matrix.values[i][j] = -1
		}
	}
	return matrix
}

func markMatrix(movements []movement, matrix mapMatrix) mapMatrix{
	for _, mov := range movements {
		if mov.src.x < mov.dst.x {
			matrix = markMatrixX(mov, matrix, true)
		} else {
			matrix = markMatrixX(mov, matrix, false)
		}
	}
	return matrix
}

func markMatrixX(mov movement, matrix mapMatrix, positive bool) mapMatrix{
	if positive {
		for i := mov.src.x; i < mov.dst.x ; i++ {
			matrix = markMatrixY(i, mov, matrix, true)
		}
	} else {
		for i := mov.dst.x; i < mov.src.x ; i++ {
			matrix = markMatrixY(i, mov, matrix, false)
		}
	}
	return matrix
}

func markMatrixY(i int, mov movement, matrix mapMatrix, positive bool) mapMatrix {
	if positive {
		for j := mov.src.y; j < mov.dst.y ; j++ {
			matrix = markCell(i, j, matrix)
		}
	} else {
		for j := mov.dst.y; j < mov.src.y ; j++ {
			matrix = markCell(i, j, matrix)
		}
	}
	return matrix
}

func markCell(i, j int, matrix mapMatrix) mapMatrix {
	if matrix.values[i][j] == -1 {
		matrix.values[i][j] = 1
	} else {
		matrix.values[i][j] = matrix.values[i][j] + 1
	}
	return matrix
}

func countResult(matrix mapMatrix) int {
	res := 0
	for _, line := range matrix.values {
		for _, value := range line {
			if value > 1 {
				res += value
			}
		}
	}
	return res
}