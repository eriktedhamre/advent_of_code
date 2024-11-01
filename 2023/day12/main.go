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
	fmt.Print(partOne(file))
}

// Rad * Kolumn
// Om en Cell har ett godkänt state testa båda nästa state
// N length traversed in the string
// K number of conditions reached
// I did brute force for part 1 and dynamic programming for part 2. The state space is (which character you're on, which block you're on, how long your current block is)
// Jonathan Paulson
// Got stuck on trying to represent the state in only two dimensions uniquely.
// index in springs, index in condition, number of consecutive blocks.
// Also happend to see someone do it with a Map which seems easier so gonna try that first
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

	var key string = fmt.Sprintf("%d%d%d%d", springs[0], len(springs), rune(fulfilled), rune(blocks))
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache map[string]uint64 = make(map[string]uint64)

var testCaseOneSprings []rune = []rune{'?', '#', '#', '#', '?', '?', '?', '?', '?', '?', '?', '?'}
var testCaseOneConditions []int = []int{3, 2, 1}

var testCaseTwoSprings []rune = []rune{'.', '?', '?', '.', '.', '?', '?', '.', '.', '.', '?', '#', '#', '.'}
var testCaseTwoConditions []int = []int{1, 1, 3}

var testCaseThreeSprings []rune = []rune{'?', '#', '#', '#', '?', '?', '?', '?', '?', '?', '?', '?'}
var testCaseThreeConditions []int = []int{3, 2, 1}

func main() {

	testSprings := testCaseThreeSprings
	testConditions := testCaseThreeConditions

	springsTimes5 := make([]rune, 0, len(testSprings)*5+4)
	conditionsTimes5 := make([]int, 0, len(testConditions)*5)

	springsTimes5 = append(springsTimes5, testSprings...)
	conditionsTimes5 = append(conditionsTimes5, testConditions...)

	for i := 0; i < 4; i++ {
		conditionsTimes5 = append(conditionsTimes5, testConditions...)
		springsTimes5 = append(springsTimes5, '?')
		springsTimes5 = append(springsTimes5, testSprings...)
	}

	fmt.Println(solverWithCache(springsTimes5, conditionsTimes5, 0, 0))
}

// Rad * Kolumn
// Om en Cell har ett godkänt state testa båda nästa state
// N length traversed in the string
// K number of conditions reached
// I did brute force for part 1 and dynamic programming for part 2. The state space is (which character you're on, which block you're on, how long your current block is)
// Jonathan Paulson
// Got stuck on trying to represent the state in only two dimensions uniquely.
// index in springs, index in condition, number of consecutive blocks.
// Also happend to see someone do it with a Map which seems easier so gonna try that first
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

func solverWithCacheTwo(springs []rune, conditions []int, fulfilled uint64, blocks int) uint64 {
	if len(springs) == 0 {
		if len(conditions) == 0 {
			return 1
		} else if blocks == conditions[0] {
			return boolToInt(len(conditions[1:]) == 0)
		}
		return boolToInt(len(conditions) == 0)
	}

	var key string
	key = fmt.Sprintf("%d%d%d%d", springs[0], len(springs), rune(fulfilled), rune(blocks))

	var res uint64 = 0
	switch {
	case springs[0] == '#':
		if value, ok := cache[key]; ok {
			res = value
		} else if len(conditions) == 0 || blocks == conditions[0] {
			res = 0
		} else if blocks < conditions[0] {
			res = solverWithCache(springs[1:], conditions, fulfilled, blocks+1)
		}

	case springs[0] == '.':
		if value, ok := cache[key]; ok {
			res = value
		} else if blocks != 0 {
			if len(conditions) > 0 && blocks == conditions[0] {
				res = solverWithCache(springs[1:], conditions[1:], fulfilled+1, 0)
			} else {
				res = 0
			}
		} else {
			res = solverWithCache(springs[1:], conditions, fulfilled, 0)
		}
	case springs[0] == '?':
		copySlice := make([]rune, len(springs))
		springs[0] = '#'
		keyHash := fmt.Sprintf("%d%d%d%d", springs[0], len(springs), rune(fulfilled), rune(blocks))
		if value, ok := cache[keyHash]; ok {
			res += value
		} else {
			copy(copySlice, springs)
			res += solverWithCache(copySlice, conditions, fulfilled, blocks)
		}
		springs[0] = '.'
		keyDot := fmt.Sprintf("%d%d%d%d", springs[0], len(springs), rune(fulfilled), rune(blocks))
		if value, ok := cache[keyDot]; ok {
			res += value
		} else {
			copy(copySlice, springs)
			res += solverWithCache(copySlice, conditions, fulfilled, blocks)
		}
		springs[0] = '?'
	}
	cache[key] = res

	return res
}
