package main

import (
	"bufio"
	"fmt"
	"math"
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
	var cumSum uint64 = 0
	var wins uint64
	var winningNumbers []int
	var myNumbers []int
	var row string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		wins = 0
		row = scanner.Text()
		winningNumbers, myNumbers = stringToIntSlice(row)

		for _, number := range myNumbers {
			if contains(winningNumbers, number) {
				wins++
			}
		}
		switch wins {
		case 0:
			cumSum += 0
		case 1:
			cumSum += 1
		default:
			cumSum += uint64(math.Pow(float64(2), float64(wins-1)))
		}

	}
	return cumSum

}

func stringToIntSlice(row string) ([]int, []int) {
	var winRes []int
	var handRes []int
	var tmp int
	_, row, _ = strings.Cut(row, ":")
	winningHand, myHand, _ := strings.Cut(row, "|")
	winningStrings := strings.Fields(winningHand)
	myStrings := strings.Fields(myHand)
	for _, s := range winningStrings {
		tmp, _ = strconv.Atoi(s)
		winRes = append(winRes, tmp)
	}
	for _, s := range myStrings {
		tmp, _ = strconv.Atoi(s)
		handRes = append(handRes, tmp)
	}
	return winRes, handRes
}

func contains[T comparable](slice []T, item T) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
