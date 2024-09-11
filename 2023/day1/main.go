package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

func partTwo(file *os.File) uint64 {
	var cum_sum uint64 = 0
	scanner := bufio.NewScanner(file)

	var first rune
	var last rune
	var combined string
	for scanner.Scan() {
		first = 'i'
		last = 'i'
		line := scanner.Text()
		i := 0
		for i < len(line) {
			switch {
			case isDigit(line[i]):
				first, last = modifyFirstAndLast(first, last, rune(line[i]))
				i += 1
			case (len(line)-i >= 3) && line[i:i+3] == "one":
				first, last = modifyFirstAndLast(first, last, '1')
				i += 2
			case (len(line)-i >= 3) && line[i:i+3] == "two":
				first, last = modifyFirstAndLast(first, last, '2')
				i += 2
			case (len(line)-i >= 5) && line[i:i+5] == "three":
				first, last = modifyFirstAndLast(first, last, '3')
				i += 4
			case (len(line)-i >= 4) && line[i:i+4] == "four":
				first, last = modifyFirstAndLast(first, last, '4')
				i += 4
			case (len(line)-i >= 4) && line[i:i+4] == "five":
				first, last = modifyFirstAndLast(first, last, '5')
				i += 3
			case (len(line)-i >= 3) && line[i:i+3] == "six":
				first, last = modifyFirstAndLast(first, last, '6')
				i += 3
			case (len(line)-i >= 5) && line[i:i+5] == "seven":
				first, last = modifyFirstAndLast(first, last, '7')
				i += 4
			case (len(line)-i >= 5) && line[i:i+5] == "eight":
				first, last = modifyFirstAndLast(first, last, '8')
				i += 4
			case (len(line)-i >= 4) && line[i:i+4] == "nine":
				first, last = modifyFirstAndLast(first, last, '9')
				i += 3
			default:
				i++
			}
		}
		combined = string(first) + string(last)

		number, _ := strconv.Atoi(combined)
		cum_sum = cum_sum + uint64(number)
	}

	return cum_sum
}

func partOne() {
	var cum_sum uint64 = 0
	scanner := bufio.NewScanner(os.Stdin)

	var first rune
	var last rune
	var combined string
	for scanner.Scan() {
		first = 'i'
		last = 'i'
		for _, c := range scanner.Text() {
			if unicode.IsDigit(c) {
				if first == 'i' {
					first = c
					last = c
				} else {
					last = c
				}
			}
		}

		combined = string(first) + string(last)

		number, _ := strconv.Atoi(combined)
		cum_sum = cum_sum + uint64(number)
	}

	fmt.Print(cum_sum)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func modifyFirstAndLast(first rune, last rune, value rune) (rune, rune) {
	if first == 'i' {
		first = value
	}
	last = value
	return first, last
}
