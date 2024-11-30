package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type DirMod struct {
	rowwMod, colMod int
}

var dirMods = []DirMod{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Visited struct {
	North, South, East, West bool
}

type Light struct {
	row, col int
	dir      Direction
}

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
	var energized [][]Visited = make([][]Visited, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		grid = append(grid, []rune(line))
		energized = append(energized, make([]Visited, len(line)))
	}

	energizeGrid(grid, energized)

	return calcEnergized(energized)
}

func energizeGrid(grid [][]rune, energized [][]Visited) {

	var lights []Light = make([]Light, 0)
	var light Light
	var mod DirMod
	// unsure if -1 is necessary
	lights = append(lights, Light{row: 0, col: -1, dir: East})
	var newRow int
	var newCol int

	for len(lights) > 0 {
		light = lights[len(lights)-1]
		lights = lights[:len(lights)-1]
		mod = dirMods[light.dir]
		newRow = light.row + mod.rowwMod
		newCol = light.col + mod.colMod
		// Vi vill bara lägga till en nod i kön om vi inte har varit där
		//
		//
		//
		if !inBound(newRow, newCol, len(grid), len(grid[0])) {
			continue
		} else {
			light.row = newRow
			light.col = newCol
			switch grid[newRow][newCol] {
			case '.':
			case '/':
				switch light.dir {
				case North:
					light.dir = East
				case East:
					light.dir = North
				case South:
					light.dir = West
				case West:
					light.dir = South
				}
			case '\\':
				switch light.dir {
				case North:
					light.dir = West
				case East:
					light.dir = South
				case South:
					light.dir = East
				case West:
					light.dir = North
				}
			case '|':
				switch light.dir {
				case East:
					fallthrough
				case West:
					light.dir = North
					lightSplit := Light{row: newRow, col: newCol, dir: South}
					if !visited(energized, newRow, newCol, South) {
						lights = append(lights, lightSplit)
					}
					energized[newRow][newCol].South = true
				case North:
				case South:
				}
			case '-':
				switch light.dir {
				case North:
					fallthrough
				case South:
					light.dir = East
					lightSplit := Light{row: newRow, col: newCol, dir: West}
					if !visited(energized, newRow, newCol, West) {
						lights = append(lights, lightSplit)
					}
					energized[newRow][newCol].West = true
				case East:
				case West:
				}

			}
			if !visited(energized, newRow, newCol, light.dir) {
				lights = append(lights, light)
				switch light.dir {
				case North:
					energized[newRow][newCol].North = true
				case South:
					energized[newRow][newCol].South = true
				case East:
					energized[newRow][newCol].East = true
				case West:
					energized[newRow][newCol].West = true
				}
			}
		}

	}
}

func visited(energized [][]Visited, row, col int, dir Direction) bool {
	var res bool = false
	switch dir {
	case North:
		res = energized[row][col].North
	case South:
		res = energized[row][col].South
	case East:
		res = energized[row][col].East
	case West:
		res = energized[row][col].West
	}
	return res
}

func calcEnergized(energized [][]Visited) uint64 {
	var sum uint64 = 0
	for _, row := range energized {
		for _, c := range row {
			if c.East || c.North || c.South || c.West {
				sum++
			}
		}
	}
	return sum
}

func inBound(row, col, rows, cols int) bool {
	return row > -1 && row < rows && col > -1 && col < cols
}
