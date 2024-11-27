package main

import (
	"bufio"
	"fmt"
	"os"
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

func hashAlgorithm(s string) uint64 {
	var sum uint64 = 0
	for _, v := range s {
		sum += uint64(v)
		sum *= 17
		sum = sum % 256
	}
	return sum
}
