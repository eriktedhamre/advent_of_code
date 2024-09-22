package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Hands struct {
	mine, wins []int
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

func partTwo(file *os.File) uint64 {
	var cumSum uint64 = 0
	var wins uint64
	var winningNumbers []int
	var myNumbers []int
	var row string
	scanner := bufio.NewScanner(file)
	var numberOfTickets []int
	var hands []Hands
	for scanner.Scan() {
		wins = 0
		row = scanner.Text()
		winningNumbers, myNumbers = stringToIntSlice(row)

		winCopy := make([]int, len(winningNumbers))
		myCopy := make([]int, len(myNumbers))
		copy(winCopy, winningNumbers)
		copy(myCopy, myNumbers)
		hands = append(hands, Hands{mine: myCopy, wins: winCopy})
	}

	numberOfTickets = make([]int, len(hands))
	for i := 0; i < len(numberOfTickets); i++ {
		numberOfTickets[i] = 1
	}

	for hand_index, h := range hands {
		wins = 0
		for _, v := range h.mine {
			if contains(h.wins, v) {
				wins++
			}
		}
		for i := 0; i < int(wins); i++ {
			if hand_index+1+i < len(numberOfTickets) {
				numberOfTickets[hand_index+1+i] += numberOfTickets[hand_index]
			}
		}
	}

	for _, v := range numberOfTickets {
		cumSum += uint64(v)
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
