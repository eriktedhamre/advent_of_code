package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	row uint64
	col uint64
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
	fmt.Print(partTwo(file))
}

func partTwo(file *os.File) uint64 {
	var sum uint64 = 0
	var oneMillion uint64 = 1000000
	var image [][]rune
	var line string
	var emptyCol uint64
	var emptyRows []uint64
	var emptyCols []uint64
	var prefixEmptyRows []uint64
	var prefixEmptyCols []uint64
	var nodes []Node

	// Keep one row slice where you denote if a row is empty
	// Keep one col slice where you denote if a column is empty
	// Calculate a prefix sum slice for both of them
	// Update the indicies of the saved nodes
	// according to the values in the prefix sum slices
	// 82
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		row := []rune(line)
		image = append(image, row)
		emptyRows = append(emptyRows, boolToInt(!strings.ContainsAny(string(line), "#")))
		// if !strings.ContainsAny(string(line),"#") {
		// 	// empty_row := make([]rune, len(image[0]))
		// 	// for i, _ := range empty_row {
		// 	// 	empty_row[i] = '.'
		// 	// }
		// 	// image = append(image, empty_row)
		// }
	}

	for col := 0; col < len(image[0]); col++ {
		emptyCol = 1
	COL_LOOP:
		for row := 0; row < len(image); row++ {
			if image[row][col] == '#' {
				emptyCol = 0
				break COL_LOOP
			}
		}
		emptyCols = append(emptyCols, emptyCol)
	}

	prefixEmptyRows = prefixSum(emptyRows)
	prefixEmptyCols = prefixSum(emptyCols)

	for row := 0; row < len(image); row++ {
		for col := 0; col < len(image[0]); col++ {
			if image[row][col] == '#' {
				nodes = append(nodes, Node{uint64(row) + prefixEmptyRows[row]*oneMillion,
					uint64(col) + prefixEmptyCols[col]*oneMillion})
			}
		}
	}

	for i := 0; i < len(nodes); i++ {

		for j := i + 1; j < len(nodes); j++ {

			if nodes[i].row > nodes[j].row {
				sum += nodes[i].row - nodes[j].row
			} else {
				sum += nodes[j].row - nodes[i].row
			}

			if nodes[i].col > nodes[j].col {
				sum += nodes[i].col - nodes[j].col
			} else {
				sum += nodes[j].col - nodes[i].col
			}
		}

	}

	return sum

}

func partOne(file *os.File) uint64 {
	var sum uint64 = 0
	var image [][]rune
	var line string
	var emptyCol uint64
	var emptyRows []uint64
	var emptyCols []uint64
	var prefixEmptyRows []uint64
	var prefixEmptyCols []uint64
	var nodes []Node

	// Keep one row slice where you denote if a row is empty
	// Keep one col slice where you denote if a column is empty
	// Calculate a prefix sum slice for both of them
	// Update the indicies of the saved nodes
	// according to the values in the prefix sum slices

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		row := []rune(line)
		image = append(image, row)
		emptyRows = append(emptyRows, boolToInt(!strings.ContainsAny(string(line), "#")))
		// if !strings.ContainsAny(string(line),"#") {
		// 	// empty_row := make([]rune, len(image[0]))
		// 	// for i, _ := range empty_row {
		// 	// 	empty_row[i] = '.'
		// 	// }
		// 	// image = append(image, empty_row)
		// }
	}

	for col := 0; col < len(image[0]); col++ {
		emptyCol = 1
	COL_LOOP:
		for row := 0; row < len(image); row++ {
			if image[row][col] == '#' {
				emptyCol = 0
				break COL_LOOP
			}
		}
		emptyCols = append(emptyCols, emptyCol)
	}

	prefixEmptyRows = prefixSum(emptyRows)
	prefixEmptyCols = prefixSum(emptyCols)

	for row := 0; row < len(image); row++ {
		for col := 0; col < len(image[0]); col++ {
			if image[row][col] == '#' {
				nodes = append(nodes, Node{uint64(row) + prefixEmptyRows[row],
					uint64(col) + prefixEmptyCols[col]})
			}
		}
	}

	for i := 0; i < len(nodes); i++ {

		for j := i + 1; j < len(nodes); j++ {

			if nodes[i].row > nodes[j].row {
				sum += nodes[i].row - nodes[j].row
			} else {
				sum += nodes[j].row - nodes[i].row
			}

			if nodes[i].col > nodes[j].col {
				sum += nodes[i].col - nodes[j].col
			} else {
				sum += nodes[j].col - nodes[i].col
			}
		}

	}

	return sum

}

func prefixSum(nums []uint64) []uint64 {
	prefixSums := make([]uint64, len(nums))
	if len(nums) == 0 {
		return prefixSums
	}

	prefixSums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSums[i] = prefixSums[i-1] + nums[i]
	}
	return prefixSums
}

func boolToInt(b bool) uint64 {
	if b {
		return 1
	}
	return 0
}
