package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	var line string
	var grid [][]rune = make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		grid = append(grid, []rune(line))
	}
	for i := 0; i < len(grid[0]); i++ {
		moveRocksInColumnNorth(grid, i)
	}
	return calculateTotalLoad(grid)
}

func partTwo(file *os.File) uint64 {
	var line string
	var grid [][]rune = make([][]rune, 0)
	cache := make(map[string]string)
	var key string
	var noKey bool = true
	cycles := 1000000000

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		grid = append(grid, []rune(line))
	}
	for i := 0; i < cycles; i++ {
		if noKey {
			key = createKey(grid)
		}
		if value, ok := cache[key]; ok {
			key = value
			noKey = false
		} else {
			updateGrid(grid, key)
			runOneCycle(grid)
			noKey = true
			cache[key] = createKey(grid)
		}

	}
	if noKey {
		key = cache[key]
	}
	updateGrid(grid, key)
	for _, v := range grid {
		fmt.Println(string(v))
	}
	return calculateTotalLoad(grid)
}

func runOneCycle(grid [][]rune) {
	for i := 0; i < len(grid[0]); i++ {
		moveRocksInColumnNorth(grid, i)
	}
	for i := 0; i < len(grid); i++ {
		moveRocksinRowWest(grid, i)
	}
	for i := 0; i < len(grid[0]); i++ {
		moveRocksInColumnSouth(grid, i)
	}
	for i := 0; i < len(grid); i++ {
		moveRocksinRowEast(grid, i)
	}
}

func createKey(grid [][]rune) string {
	var builder strings.Builder

	for _, row := range grid {
		builder.WriteString(string(row))
	}

	return builder.String()
}

func updateGrid(grid [][]rune, key string) {
	var row int
	var col int
	var cols int = len(grid[0])
	for i, r := range key {
		row = i / cols
		col = i % cols
		grid[row][col] = r
	}
}

func moveRocksInColumnNorth(grid [][]rune, col int) {
	var blockerRow int = 0
	var numberOfRocksSinceBlocker int = 0

	for i := 0; i < len(grid); i++ {
		switch grid[i][col] {
		case 'O':
			grid[blockerRow+numberOfRocksSinceBlocker][col] = 'O'
			if blockerRow+numberOfRocksSinceBlocker != i {
				grid[i][col] = '.'
			}
			numberOfRocksSinceBlocker++
		case '#':
			numberOfRocksSinceBlocker = 0
			blockerRow = i + 1
		default:
		}
	}
}

func moveRocksinRowWest(grid [][]rune, row int) {
	var blockerCol int = 0
	var numberOfRocksSinceBlocker int = 0

	for i := 0; i < len(grid[0]); i++ {
		switch grid[row][i] {
		case 'O':
			grid[row][blockerCol+numberOfRocksSinceBlocker] = 'O'
			if blockerCol+numberOfRocksSinceBlocker != i {
				grid[row][i] = '.'
			}
			numberOfRocksSinceBlocker++
		case '#':
			numberOfRocksSinceBlocker = 0
			blockerCol = i + 1
		default:
		}
	}
}

func moveRocksInColumnSouth(grid [][]rune, col int) {
	var blockerRow int = len(grid) - 1
	var numberOfRocksSinceBlocker int = 0

	for i := len(grid) - 1; i > -1; i-- {
		switch grid[i][col] {
		case 'O':
			grid[blockerRow-numberOfRocksSinceBlocker][col] = 'O'
			if blockerRow-numberOfRocksSinceBlocker != i {
				grid[i][col] = '.'
			}
			numberOfRocksSinceBlocker++
		case '#':
			numberOfRocksSinceBlocker = 0
			blockerRow = i - 1
		default:
		}
	}
}

func moveRocksinRowEast(grid [][]rune, row int) {
	var blockerCol int = len(grid[0]) - 1
	var numberOfRocksSinceBlocker int = 0

	for i := len(grid[0]) - 1; i > -1; i-- {
		switch grid[row][i] {
		case 'O':
			grid[row][blockerCol-numberOfRocksSinceBlocker] = 'O'
			if blockerCol-numberOfRocksSinceBlocker != i {
				grid[row][i] = '.'
			}
			numberOfRocksSinceBlocker++
		case '#':
			numberOfRocksSinceBlocker = 0
			blockerCol = i - 1
		default:
		}
	}
}

func calculateTotalLoad(grid [][]rune) uint64 {
	var sum uint64 = 0
	var rockCount uint64 = 0
	for i := 0; i < len(grid); i++ {
		rockCount = 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'O' {
				rockCount++
			}
		}
		sum += rockCount * uint64((len(grid) - i))
	}
	return sum
}
