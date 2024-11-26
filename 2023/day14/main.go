package main

import (
	"bufio"
	"fmt"
	"os"
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
	for i := 0; i < len(grid[0]); i++ {
		moveRocksInColumn(grid, i)
	}
	return calculateTotalLoad(grid)
}

func moveRocksInColumn(grid [][]rune, col int) {
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
