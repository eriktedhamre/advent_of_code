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
	var sum uint64 = 0
	var line string
	var grid []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		if len(line) == 0 {
			sum += processGrid(grid)
			grid = grid[:0]
		} else {
			grid = append(grid, line)
		}
	}

	sum += processGrid(grid)

	return sum

}

func processGrid(grid []string) uint64 {

	if horizontal := checkHorizontal(grid); horizontal != -1 {
		return uint64(horizontal) * 100
	}
	return uint64(checkHorizontal(matrixRotate(grid)))
}

func checkHorizontal(grid []string) int {

	var top int = 0
	var bottom int = len(grid) - 1

	for i := len(grid) % 2; i < bottom; i += 2 {
		if checkSpan(grid, bottom, i) {
			return (bottom+1-i)/2 + i
		}
	}

	for i := bottom - (len(grid) % 2); i > 0; i -= 2 {
		if checkSpan(grid, i, top) {
			return (i + 1) / 2
		}
	}

	return -1
}

func checkSpan(grid []string, bottom, top int) bool {

	if bottom-top == 1 {
		return grid[bottom] == grid[top]
	}

	for i := 0; i < (bottom-top+1)/2; i++ {
		if grid[bottom-i] != grid[top+i] {
			return false
		}
	}
	return true
}

func matrixRotate(grid []string) []string {
	tmp := make([][]rune, len(grid[0]))
	result := make([]string, 0)
	for row := range tmp {
		tmp[row] = make([]rune, len(grid))
	}

	for i, s := range grid {
		for j := range s {
			tmp[j][i] = rune(s[j])
		}
	}

	for i := range tmp {
		result = append(result, string(tmp[i]))
	}

	return result
}
