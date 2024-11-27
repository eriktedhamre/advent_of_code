package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	label string
	lens  int
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
	var line string
	var initSeq []string = make([]string, 0)
	var res uint64 = 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line = scanner.Text()
	initSeq = strings.Split(line, ",")
	for _, v := range initSeq {
		res += hashAlgorithm(v)
	}

	return res
}

func partTwo(file *os.File) uint64 {
	var line string
	var initSeq []string = make([]string, 0)
	var label string
	var op rune
	var number int
	var boxIndex uint64

	var boxes [][]Lens = make([][]Lens, 256)
	for i := 0; i < len(boxes); i++ {
		boxes[i] = make([]Lens, 0)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line = scanner.Text()
	initSeq = strings.Split(line, ",")
	for _, v := range initSeq {
		label, op, number = splitCommand(v)
		boxIndex = hashAlgorithm(label)

		if op == '=' {
			boxes[boxIndex] = equalOp(label, boxes[boxIndex], number)
		} else {
			boxes[boxIndex] = dashOp(label, boxes[boxIndex])
		}
	}

	return calculateFocusingPower(boxes)
}

func equalOp(label string, box []Lens, number int) []Lens {
	var exists bool = false
DONE:
	for i := 0; i < len(box); i++ {
		if box[i].label == label {
			box[i].lens = number
			exists = true
			break DONE
		}
	}
	if !exists {
		box = append(box, Lens{label: label, lens: number})
	}
	return box
}

func dashOp(label string, box []Lens) []Lens {
	var removeIndex int = -1

DONE:
	for i := 0; i < len(box); i++ {
		if box[i].label == label {
			removeIndex = i
			break DONE
		}
	}

	if removeIndex != -1 {
		return append(box[:removeIndex], box[removeIndex+1:]...)
	} else {
		return box
	}
}

func calculateFocusingPower(boxes [][]Lens) uint64 {
	var sum uint64 = 0
	for i, box := range boxes {
		for j, lens := range box {
			sum += uint64((i + 1)) * uint64(j+1) * uint64(lens.lens)
		}
	}
	return sum
}

func splitCommand(s string) (string, rune, int) {
	var tmp []string
	if strings.Contains(s, "=") {
		tmp = strings.Split(s, "=")
		number, err := strconv.Atoi(tmp[1])
		if err != nil {
			fmt.Errorf("Unsuccessfull Atoi for %s", tmp[1])
		}
		return tmp[0], '=', number
	} else {
		return strings.TrimSuffix(s, "-"), '-', -1
	}
}

func hashAlgorithm(s string) uint64 {
	var sum uint64 = 0
	for _, v := range s {
		sum += uint64(v)
		sum *= 17
		sum = sum % 256
	}
	return sum
}
