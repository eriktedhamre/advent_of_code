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
			var val_i uint64
			var val_j uint64
			switch hands[i].hand[k] {
			case 'A':
				val_i = 14
			case 'K':
				val_i = 13
			case 'Q':
				val_i = 12
			case 'J':
				val_i = 11
			case 'T':
				val_i = 10
			case '9':
				val_i = 9
			case '8':
				val_i = 8
			case '7':
				val_i = 7
			case '6':
				val_i = 6
			case '5':
				val_i = 5
			case '4':
				val_i = 4
			case '3':
				val_i = 3
			case '2':
				val_i = 2
			}
			switch hands[j].hand[k] {
			case 'A':
				val_j = 14
			case 'K':
				val_j = 13
			case 'Q':
				val_j = 12
			case 'J':
				val_j = 11
			case 'T':
				val_j = 10
			case '9':
				val_j = 9
			case '8':
				val_j = 8
			case '7':
				val_j = 7
			case '6':
				val_j = 6
			case '5':
				val_j = 5
			case '4':
				val_j = 4
			case '3':
				val_j = 3
			case '2':
				val_j = 2
			}
			return val_i > val_j
		}
		return false
	})

	for i, v := range hands {
		cumSum += uint64((len(hands) - i)) * v.bid
	}

	return cumSum
}

func calculateHandTier(hand *Hand, cards []int) int8 {
	fiveOfAKind := false
	fourOfAKind := false
	threeOfAKind := false
	onePair := 0
	highCard := 0
	for _, char := range hand.hand {
		switch char {
		case 'A':
			cards[12] += 1
		case 'K':
			cards[11] += 1
		case 'Q':
			cards[10] += 1
		case 'J':
			cards[9] += 1
		case 'T':
			cards[8] += 1
		case '9':
			cards[7] += 1
		case '8':
			cards[6] += 1
		case '7':
			cards[5] += 1
		case '6':
			cards[4] += 1
		case '5':
			cards[3] += 1
		case '4':
			cards[2] += 1
		case '3':
			cards[1] += 1
		case '2':
			cards[0] += 1
		}
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
