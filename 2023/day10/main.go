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

var north Pos = Pos{-1, 0}
var south Pos = Pos{1, 0}
var east Pos = Pos{0, 1}
var west Pos = Pos{0, -1}

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
	var tiles [][]rune
	var visited [][]bool
	var line string
	var S_row int
	var S_col int
	var rows int
	var cols int
	var goal bool
	var moveA Pos
	var moveB Pos

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
		for j, _ := range row {
			if row[j] == 'S' {
				S_row = i
				S_col = j
				break FIND_S
			}
		}
	}

	visited[S_row][S_col] = true

	loop := make([]Pos, 0)
	loop = append(loop, Pos{S_row, S_col})

	switch {
	// South -> North
	case inBound(S_row+1, S_col, rows, cols) && ((tiles[S_row+1][S_col] == '|') ||
		(tiles[S_row+1][S_col] == 'L') ||
		(tiles[S_row+1][S_col] == 'J')):
		loop = append(loop, Pos{S_row + 1, S_col})
		visited[S_row+1][S_col] = true
	// North -> South
	case inBound(S_row-1, S_col, rows, cols) && ((tiles[S_row-1][S_col] == '|') ||
		(tiles[S_row-1][S_col] == '7') ||
		(tiles[S_row-1][S_col] == 'F')):
		loop = append(loop, Pos{S_row - 1, S_col})
		visited[S_row-1][S_col] = true
	// East -> West
	case inBound(S_row, S_col+1, rows, cols) && ((tiles[S_row][S_col+1] == '-') ||
		(tiles[S_row][S_col+1] == 'J') ||
		(tiles[S_row][S_col+1] == '7')):
		loop = append(loop, Pos{S_row, S_col + 1})
		visited[S_row][S_col+1] = true
	// West -> East
	case inBound(S_row, S_col-1, rows, cols) && ((tiles[S_row][S_col-1] == '-') ||
		(tiles[S_row][S_col-1] == 'L') ||
		(tiles[S_row][S_col-1] == 'F')):
		loop = append(loop, Pos{S_row, S_col - 1})
		visited[S_row][S_col-1] = true
	}

	// If both nodes are visited we are back where we started :) I'm a winner
DONE:
	for {
		tail := loop[len(loop)-1]
		switch {
		case tiles[tail.row][tail.col] == '|':
			moveA = north
			moveB = south
		case tiles[tail.row][tail.col] == '-':
			moveA = east
			moveB = west
		case tiles[tail.row][tail.col] == 'L':
			moveA = north
			moveB = east
		case tiles[tail.row][tail.col] == 'J':
			moveA = north
			moveB = west
		case tiles[tail.row][tail.col] == '7':
			moveA = south
			moveB = west
		case tiles[tail.row][tail.col] == 'F':
			moveA = south
			moveB = east
		default:
			fmt.Println("This is a mistake")
		}
		if loop, goal = nextStepReachesGoal(tail, moveA, moveB, loop, visited); goal {
			break DONE
		}

	}

	return int64((len(loop) + 1) / 2)
}

func nextStepReachesGoal(pos, moveA, moveB Pos, loop []Pos, visited [][]bool) ([]Pos, bool) {
	Result := false
	if visited[pos.row+moveA.row][pos.col+moveA.col] &&
		visited[pos.row+moveB.row][pos.col+moveB.col] {
		Result = true
	} else if visited[pos.row+moveA.row][pos.col+moveA.col] {
		loop = append(loop, Pos{pos.row + moveB.row, pos.col + moveB.col})
		visited[pos.row+moveB.row][pos.col+moveB.col] = true
	} else {
		loop = append(loop, Pos{pos.row + moveA.row, pos.col + moveA.col})
		visited[pos.row+moveA.row][pos.col+moveA.col] = true
	}
	return loop, Result
}

func inBound(row, col, rows, cols int) bool {
	return row > -1 && row < rows && col > -1 && col < cols
}
