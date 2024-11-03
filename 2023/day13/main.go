package main

import (
	"bufio"
	"fmt"
	"os"
)

var cache map[string]uint64 = make(map[string]uint64)

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

// func partTwo(file *os.File) uint64 {
// 	var sum uint64 = 0
// 	var line string

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		clear(cache)
// 		line = scanner.Text()
		
// 	}
// 	return sum

// }

func partOne(file *os.File) uint64 {
	//var sum uint64 = 0
	var line string
	var grid [] string
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		grid = append(grid, line)
	}

	// for _, row := range grid {
	// 	for _, col := range row {
	// 		fmt.Printf("%c", col)
	// 	}
	// 	fmt.Print("\n")
	// }

	// starting from the middle
	// for each row see if adjacent rows are equal
	// if so continue one step in each direction
	//
	// starting from the middle
	// for each column see if each adjavent column are equal
	// if so continue one step in each direction

	return uint64(checkHorizontal(grid)) * 100

}

func checkHorizontal(grid []string) int {
	
	var bottom int = len(grid) - 1
	var top int = 0

	// span_size/2 + i - 1??????
	
	for i := len(grid) % 2; i < bottom; i += 2 {
		if bottom == i {
			
		}
		if checkSpan(grid, bottom, i) {
			fmt.Printf("%d, %d, %d \n", bottom, i, (bottom +1 - i)/2 + i)
			return (bottom + 1 - i)/2 + i // big sus
		}
	}

	for i := bottom - len(grid) % 2; i > 0; i-=2 {
		if checkSpan(grid, i, top) {
			return i/2 + i // big sus :2 electric bogaloo
		}
	}
	return -1
}

func checkSpan(grid []string, bottom, top int) bool {

	if bottom - top == 1 {
		return grid[bottom] == grid[top]
	}

	for i := 0; i < (bottom - top)/2; i++ {
		if grid[bottom - i] != grid[top + i] {
			return false
		}
	}
	return true
}

func isEqual(sOne, sTwo []rune) bool {
	if len(sOne) != len(sTwo) {
		return false
	}

	for i := 0; i < len(sOne); i++ {
		if sOne[i] != sTwo[i] {
			return false
		}
	}
	return true
}

func inBound(row, col, rows, cols int) bool {
	return row > -1 && row < rows && col > -1 && col < cols
}

