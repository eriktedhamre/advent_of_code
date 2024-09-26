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
	fmt.Print(partTwo(file))
}

func partTwo(file *os.File) uint64 {
	var row string
	var time uint64
	var distance uint64
	var builder strings.Builder
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	row = scanner.Text()
	row = strings.TrimPrefix(row, "Time:")
	for _, v := range strings.Fields(row) {
		builder.WriteString(v)
	}
	tmp, _ := strconv.Atoi(builder.String())
	time = uint64(tmp)
	builder.Reset()
	scanner.Scan()
	row = scanner.Text()
	row = strings.TrimPrefix(row, "Distance:")
	for _, v := range strings.Fields(row) {
		builder.WriteString(v)
	}
	tmp, _ = strconv.Atoi(builder.String())
	distance = uint64(tmp)

	return waysToWin(time, distance)
}

func partOne(file *os.File) uint64 {
	var cumSum uint64 = 1
	var row string
	var times []uint64
	var distances []uint64
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	row = scanner.Text()
	row = strings.TrimPrefix(row, "Time:")
	for _, v := range strings.Fields(row) {
		tmp, err := strconv.Atoi(v)
		if err != nil {
			fmt.Errorf("Atoi conversion failed for %s", v)
			panic(nil)
		}
		times = append(times, uint64(tmp))
	}
	scanner.Scan()
	row = scanner.Text()
	row = strings.TrimPrefix(row, "Distance:")
	for _, v := range strings.Fields(row) {
		tmp, err := strconv.Atoi(v)
		if err != nil {
			fmt.Errorf("Atoi conversion failed for %s", v)
			panic(nil)
		}
		distances = append(distances, uint64(tmp))
	}

	for i, time := range times {
		cumSum *= waysToWin(time, distances[i])
	}

	return cumSum
}

func waysToWin(time uint64, distance uint64) uint64 {
	var i uint64
	var wins uint64 = 0
	for i = 0; i <= time; i++ {
		if (i * (time - i)) > distance {
			wins++
		}
	}
	return wins
}
