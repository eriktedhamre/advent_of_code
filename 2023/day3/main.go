package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Pair struct {
	row, col int
}

var moves []Pair = []Pair{{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	{1, -1}, {1, 1}, {-1, 1}, {-1, -1}}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Print(partTwo(file))
}

func partOne(file *os.File) uint64 {
	var cumSum uint64 = 0
	var tmp int
	var row string
	var gridRow []rune
	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row = scanner.Text()
		gridRow = gridRow[:0]
		gridRow = append(gridRow, []rune(row)...)
		rowCopy := make([]rune, len(gridRow))
		copy(rowCopy, gridRow)
		grid = append(grid, rowCopy)
	}

	var builder strings.Builder
	for i, row := range grid {
		builder.Reset()
		for j, r := range row {
			if unicode.IsDigit(r) {
				builder.WriteString(string(r))
				if j == len(grid[0])-1 {
					if isValid(builder.String(), grid, i, j) {
						tmp, _ = strconv.Atoi(builder.String())
						cumSum += uint64(tmp)
					}
				}
			} else if len(builder.String()) == 0 {
				continue
			} else {
				if isValid(builder.String(), grid, i, j-1) {
					tmp, _ = strconv.Atoi(builder.String())
					cumSum += uint64(tmp)
				}
				builder.Reset()
			}
		}
	}

	return cumSum
}

func partTwo(file *os.File) uint64 {

	var cumSum uint64 = 0
	var tmp int
	var row string
	var gridRow []rune
	var grid [][]rune
	gearToNumbers := make(map[Pair][]uint64)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row = scanner.Text()
		gridRow = gridRow[:0]
		gridRow = append(gridRow, []rune(row)...)
		rowCopy := make([]rune, len(gridRow))
		copy(rowCopy, gridRow)
		grid = append(grid, rowCopy)
	}

	var builder strings.Builder
	for i, row := range grid {
		builder.Reset()
		for j, r := range row {
			if unicode.IsDigit(r) {
				builder.WriteString(string(r))
				if j == len(grid[0])-1 {
					pairs := isValidPartTwo(builder.String(), grid, i, j)
					for _, key := range pairs {
						_, ok := gearToNumbers[key]
						if !ok {
							tmp, _ = strconv.Atoi(builder.String())
							gearToNumbers[key] = []uint64{uint64(tmp)}
						} else {
							tmp, _ = strconv.Atoi(builder.String())
							gearToNumbers[key] = append(gearToNumbers[key], uint64(tmp))
						}
					}
				}
			} else if len(builder.String()) == 0 {
				continue
			} else {
				pairs := isValidPartTwo(builder.String(), grid, i, j-1)
				for _, key := range pairs {
					_, ok := gearToNumbers[key]
					if !ok {
						tmp, _ = strconv.Atoi(builder.String())
						gearToNumbers[key] = []uint64{uint64(tmp)}
					} else {
						tmp, _ = strconv.Atoi(builder.String())
						gearToNumbers[key] = append(gearToNumbers[key], uint64(tmp))
					}
				}
				builder.Reset()
			}
		}
	}

	for _, value := range gearToNumbers {
		if len(value) == 2 {
			cumSum += value[0] * value[1]
		}
	}

	return cumSum
}

func isValid(number string, grid [][]rune, row int, col int) bool {
	valid := false
	start := col - len(number) + 1
	dot := '.'
	var r rune
done:
	for i := 0; i < len(number); i++ {
		for _, p := range moves {
			if (row+p.row < len(grid)) &&
				(row+p.row >= 0) &&
				(start+i+p.col < len(grid[0])) &&
				(start+i+p.col >= 0) {
				r = grid[row+p.row][start+i+p.col]

				if r != dot && !unicode.IsDigit(r) {
					valid = true
					break done
				}
			}
		}
	}
	return valid
}

func isValidPartTwo(number string, grid [][]rune, row int, col int) []Pair {
	var result []Pair
	start := col - len(number) + 1
	var r rune
	for i := 0; i < len(number); i++ {
		for _, p := range moves {
			if (row+p.row < len(grid)) &&
				(row+p.row >= 0) &&
				(start+i+p.col < len(grid[0])) &&
				(start+i+p.col >= 0) {
				r = grid[row+p.row][start+i+p.col]

				if r == '*' {
					pair := Pair{row + p.row, start + i + p.col}
					if !contains(result, pair) {
						result = append(result, pair)
					}
				}
			}
		}
	}
	return result
}

func contains[T comparable](slice []T, item T) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
