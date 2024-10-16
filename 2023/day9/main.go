package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fmt.Print(partOne(file))
}

func partOne(file *os.File) int64 {
	var sum int64
	var histories [][]int64
	var row string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row = scanner.Text()

		rowCopy := strings.Fields(row)
		history, _ := convertToInt64Slice(rowCopy)
		histories = append(histories, history)
		sum += solveHistory(histories)
		histories = histories[:0]
	}
	return sum
}

func solveHistory(histories [][]int64) int64 {
	// Negative numbers :)
	var hCount uint64 = 0
	var diff int64
	var allZeros bool
DONE:
	for {
		allZeros = true
		if len(histories[hCount]) < 2 {
			break DONE
		}
		histories = append(histories, make([]int64, 0, len(histories[hCount])-1))
		for i := 0; i < len(histories[hCount])-1; i++ {
			diff = histories[hCount][i+1] - histories[hCount][i]

			if diff != 0 {
				allZeros = false
			}
			histories[hCount+1] = append(histories[hCount+1], diff)
		}
		hCount++
		if allZeros {
			break DONE
		}
	}

	histories[len(histories)-1] = append(histories[len(histories)-1], 0)
	for i := len(histories) - 2; i > -1; i-- {
		histories[i] = append(histories[i], histories[i+1][len(histories[i+1])-1]+histories[i][len(histories[i])-1])
	}
	return histories[0][len(histories[0])-1]
}

func convertToInt64Slice(strSlice []string) ([]int64, error) {

	intSlice := make([]int64, len(strSlice))

	// Iterate over the string slice and convert each element
	for i, str := range strSlice {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error converting %s at index %d: %w", str, i, err)
		}
		intSlice[i] = num
	}

	return intSlice, nil
}
