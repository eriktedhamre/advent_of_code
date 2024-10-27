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

func partOne(file *os.File) uint64 {
	var sum uint64 = 0
	var line string
	var condition []int = make([]int, 0)
	var springs []rune = make([]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		fields := strings.Fields(line)
		numbersAsStrings := strings.Split(fields[1], ",")
		for i := range numbersAsStrings {
			number, _ := strconv.Atoi(numbersAsStrings[i])
			condition = append(condition, number)
		}
		springs = []rune(fields[0])

		sum += solver(springs, condition)
		condition = condition[:0]
	}

	return sum

}

func solver(springs []rune, condition []int) uint64 {
	if !contains(springs, '?') {
		return matches(springs, condition)
	}
	index := findFirst(springs, '?')
	springs[index] = '#'
	copySlice := make([]rune, len(springs))
	copy(copySlice, springs)
	addHashTag := solver(copySlice, condition)
	springs[index] = '.'
	copy(copySlice, springs)
	addDot := solver(copySlice, condition)
	return addHashTag + addDot
}

func matches(springs []rune, conditions []int) uint64 {
	var result uint64 = 1
	var segments []string = strings.FieldsFunc(string(springs), springFields)

	if len(segments) == len(conditions) {
		for i := range segments {
			if len(segments[i]) != conditions[i] {
				result = 0
			}
		}
	} else {
		result = 0
	}

	return result

}

func contains(slice []rune, elem rune) bool {
	var result bool = false
DONE:
	for _, v := range slice {
		if v == elem {
			result = true
			break DONE
		}
	}
	return result
}

func findFirst(slice []rune, elem rune) int64 {
	var result int64 = -1
DONE:
	for i := range slice {
		if slice[i] == elem {
			result = int64(i)
			break DONE
		}
	}
	return result
}

func springFields(c rune) bool {
	return c == '.'
}
