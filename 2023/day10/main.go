package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	row int
	col int
}

var moves []Pos = []Pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

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
	fmt.Print(partOne(file))
}

func partOne(file *os.File) int64 {
	var sum int64
	var tiles [][]rune
	var visited [][]bool
	var line string
	var S_row int
	var S_col int
	var rows int
	var cols int
	var loop_length int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		row := []rune(line)
		tiles = append(tiles, row)
	}

	visited = make([][]bool, len(tiles))
	for i := range visited {
		visited[i] = make([]bool, len(tiles[0]))
	}

	rows = len(tiles)
	cols = len(tiles[0])

FIND_S:
	for i, row := range tiles {
		for j, col := range row {
			if row[col] == 'S' {
				S_row = i
				S_col = j
				break FIND_S
			}
		}
	}

	visited[S_row][S_col] = true

	stack := make([]Pos, 0)

	switch {
	// South -> North
	case inBound(S_row+1, S_col, rows, cols) && ((tiles[S_row+1][S_col] == '|') ||
		(tiles[S_row+1][S_col] == 'L') ||
		(tiles[S_row+1][S_col] == 'J')):
		stack = append(stack, Pos{S_row + 1, S_col})
		visited[S_row+1][S_col] = true
	// North -> South
	case inBound(S_row-1, S_col, rows, cols) && ((tiles[S_row-1][S_col] == '|') ||
		(tiles[S_row-1][S_col] == '7') ||
		(tiles[S_row-1][S_col] == 'F')):
		stack = append(stack, Pos{S_row - 1, S_col})
		visited[S_row-1][S_col] = true
	// East -> West
	case inBound(S_row, S_col+1, rows, cols) && ((tiles[S_row][S_col+1] == '-') ||
		(tiles[S_row][S_col+1] == 'J') ||
		(tiles[S_row][S_col+1] == '7')):
		stack = append(stack, Pos{S_row, S_col + 1})
		visited[S_row][S_col+1] = true
	// West -> East
	case inBound(S_row, S_col-1, rows, cols) && ((tiles[S_row][S_col-1] == '-') ||
		(tiles[S_row][S_col-1] == 'L') ||
		(tiles[S_row][S_col-1] == 'F')):
		stack = append(stack, Pos{S_row, S_col - 1})
		visited[S_row][S_col-1] = true
	}
	loop_length++

	// If both nodes are visited we are back where we started :) I'm a winner
	for len(stack) > 0 {
		switch {
		case stack[0] == '':

		}

	}

	return sum
}

func inBound(row, col, rows, cols int) bool {
	return row > -1 && row < rows && col > -1 && col < cols
}
