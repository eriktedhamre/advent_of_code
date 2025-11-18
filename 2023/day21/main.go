package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	row int
	col int
}

var moves []pos = []pos{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

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

func partOne(file *os.File) uint64 {

	var line string
	var grid [][]rune = make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		grid = append(grid, []rune(line))
	}

	var startPos pos
	for i, row := range grid {
		for j, c := range row {
			if c == 'S' {
				startPos = pos{i, j}
			}
		}
	}

	//zeroes := make([]bool, len(grid[0]))

	return uint64(reachableInSteps(grid, startPos, 64))

}

func reachableInSteps(grid [][]rune, start pos, steps int) int {

	current := map[pos]struct{}{start: {}}

	for step := 0; step < steps; step++ {

		next := make(map[pos]struct{})

		for p := range current {

			for _, d := range moves {
				nr, nc := p.row+d.row, p.col+d.col

				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
					continue
				}

				if grid[nr][nc] == '#' {
					continue
				}

				next[pos{nr, nc}] = struct{}{}
			}
		}

		current = next
	}

	return len(current)
}

func clear(visited [][]bool, zeros []bool) {
	for i := range visited {
		copy(visited[i], zeros)
	}
}
