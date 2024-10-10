package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand string
	bid  uint64
	tier int8
}

var cardMap = map[byte]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

var cardMapTwo = map[byte]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
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
	var cumSum uint64 = 0
	var input []string
	var row string
	var hands []Hand
	cards := make([]int, 13)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row = scanner.Text()

		input = strings.Fields(row)
		int_bid, err := strconv.Atoi(input[1])
		if err != nil {
			fmt.Errorf("Failed strconv for %s", row)
		}
		hands = append(hands, Hand{hand: input[0], bid: uint64(int_bid)})
	}

	for i, _ := range hands {
		hands[i].tier = calculateHandTierTwo(&hands[i], cards)
	}

	sort.Slice(hands, func(i int, j int) bool {
		if hands[i].tier != hands[j].tier {
			return hands[i].tier > hands[j].tier
		}
		for k := 0; k < 5; k++ {
			if hands[i].hand[k] == hands[j].hand[k] {
				continue
			}
			return cardMapTwo[hands[i].hand[k]] > cardMapTwo[hands[j].hand[k]]
		}
		return false
	})

	for i, v := range hands {
		cumSum += uint64((len(hands) - i)) * v.bid
	}

	return cumSum
}

func partOne(file *os.File) uint64 {
	var cumSum uint64 = 0
	var input []string
	var row string
	var hands []Hand
	cards := make([]int, 13)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row = scanner.Text()

		input = strings.Fields(row)
		int_bid, err := strconv.Atoi(input[1])
		if err != nil {
			fmt.Errorf("Failed strconv for %s", row)
		}
		hands = append(hands, Hand{hand: input[0], bid: uint64(int_bid)})
	}

	for i, _ := range hands {
		hands[i].tier = calculateHandTier(&hands[i], cards)
	}

	sort.Slice(hands, func(i int, j int) bool {
		if hands[i].tier != hands[j].tier {
			return hands[i].tier > hands[j].tier
		}
		for k := 0; k < 5; k++ {
			if hands[i].hand[k] == hands[j].hand[k] {
				continue
			}
			return cardMap[hands[i].hand[k]] > cardMap[hands[j].hand[k]]
		}
		return false
	})

	for i, v := range hands {
		cumSum += uint64((len(hands) - i)) * v.bid
	}

	return cumSum
}

func calculateHandTierTwo(hand *Hand, cards []int) int8 {
	fiveOfAKind := false
	fourOfAKind := false
	threeOfAKind := false
	onePair := 0
	highCard := 0
	var jokers int
	for _, char := range hand.hand {
		cards[cardMap[byte(char)]] += 1
	}

	for _, v := range cards {
		switch v {
		case 0:
			continue
		case 1:
			highCard += 1
		case 2:
			onePair += 1
		case 3:
			threeOfAKind = true
		case 4:
			fourOfAKind = true
		case 5:
			fiveOfAKind = true
		}
	}

	jokers = cards[cardMap[byte('J')]]

	for i, _ := range cards {
		cards[i] = 0
	}

	switch {
	case fiveOfAKind ||
		(fourOfAKind && jokers == 1) ||
		(threeOfAKind && jokers == 2) ||
		(onePair == 1 && jokers == 3) ||
		(jokers == 4):
		return 7
	case fourOfAKind ||
		(threeOfAKind && (jokers == 1)) ||
		(jokers == 3) ||
		(onePair == 2 && jokers == 2):
		return 6
	case (onePair == 1 && threeOfAKind) ||
		(onePair == 2 && jokers == 1):
		return 5
	case threeOfAKind ||
		(onePair == 1 && jokers == 1) ||
		(jokers == 2):
		return 4
	case (onePair == 2) ||
		(onePair == 1 && jokers == 1):
		return 3
	case (onePair == 1) ||
		jokers == 1:
		return 2
	default:
		return 1
	}

}

func calculateHandTier(hand *Hand, cards []int) int8 {
	fiveOfAKind := false
	fourOfAKind := false
	threeOfAKind := false
	onePair := 0
	highCard := 0
	for _, char := range hand.hand {
		cards[cardMap[byte(char)]] += 1
	}

	for _, v := range cards {
		switch v {
		case 0:
			continue
		case 1:
			highCard += 1
		case 2:
			onePair += 1
		case 3:
			threeOfAKind = true
		case 4:
			fourOfAKind = true
		case 5:
			fiveOfAKind = true
		}
	}

	for i, _ := range cards {
		cards[i] = 0
	}

	switch {
	case fiveOfAKind:
		return 7
	case fourOfAKind:
		return 6
	case onePair == 1 && threeOfAKind:
		return 5
	case threeOfAKind:
		return 4
	case onePair == 2:
		return 3
	case onePair == 1:
		return 2
	default:
		return 1
	}

}
