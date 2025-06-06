package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Print(partTwo(file))
}

func partTwo(file *os.File) uint64 {
	var sum uint64 = 0
	var line string
	var conditions []int = make([]int, 0)
	var springs []rune = make([]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		clear(cache)
		line = scanner.Text()
		fields := strings.Fields(line)
		numbersAsStrings := strings.Split(fields[1], ",")
		for i := range numbersAsStrings {
			number, _ := strconv.Atoi(numbersAsStrings[i])
			conditions = append(conditions, number)
		}
		springs = []rune(fields[0])
		springsTimes5 := make([]rune, 0, len(springs)*5+4)
		conditionsTimes5 := make([]int, 0, len(conditions)*5)

		springsTimes5 = append(springsTimes5, springs...)
		conditionsTimes5 = append(conditionsTimes5, conditions...)

		for i := 0; i < 4; i++ {
			conditionsTimes5 = append(conditionsTimes5, conditions...)
			springsTimes5 = append(springsTimes5, '?')
			springsTimes5 = append(springsTimes5, springs...)
		}

		sum += solverWithCache(springsTimes5, conditionsTimes5, 0, 0)
		conditions = conditions[:0]
	}

	return sum

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

func solverWithCache(springs []rune, conditions []int, fulfilled uint64, blocks int) uint64 {
	if len(springs) == 0 {
		if len(conditions) == 0 {
			return 1
		} else if blocks == conditions[0] {
			return boolToInt(len(conditions[1:]) == 0)
		}
		return boolToInt(len(conditions) == 0)
	}

	var key string = fmt.Sprintf("%d%d%d", springs, rune(fulfilled), rune(blocks))
	if value, ok := cache[key]; ok {
		return value
	}

	var res uint64 = 0
	switch {
	case springs[0] == '#':
		if len(conditions) == 0 || blocks == conditions[0] {
			res = 0
		} else if blocks < conditions[0] {
			res = solverWithCache(springs[1:], conditions, fulfilled, blocks+1)
		}

	case springs[0] == '.':
		if blocks != 0 {
			if len(conditions) > 0 && blocks == conditions[0] {
				res = solverWithCache(springs[1:], conditions[1:], fulfilled+1, 0)
			} else {
				res = 0
			}
		} else {
			res = solverWithCache(springs[1:], conditions, fulfilled, 0)
		}
	case springs[0] == '?':
		springs[0] = '#'
		copySlice := make([]rune, len(springs))
		copy(copySlice, springs)
		res += solverWithCache(copySlice, conditions, fulfilled, blocks)
		springs[0] = '.'
		copy(copySlice, springs)
		res += solverWithCache(copySlice, conditions, fulfilled, blocks)
	}
	cache[key] = res

	return res
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

func boolToInt(b bool) uint64 {
	if b {
		return 1
	}
	return 0
}
