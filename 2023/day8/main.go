package main

import (
	"bufio"
	"fmt"
	"os"
)

type Edges struct {
	left  string
	right string
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
	var row string
	var commands string
	var edges map[string]Edges = make(map[string]Edges)
	var nodes []string
	var endsWithZ int = 0
	var steps uint64 = 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	commands = scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		row = scanner.Text()

		// AAA = (BBB, BBB)
		if row[2] == 'A' {
			nodes = append(nodes, row[0:3])
		}
		edges[row[0:3]] = Edges{left: row[7:10], right: row[12:15]}
	}

DONE:
	for {
		for _, move := range commands {
			endsWithZ = 0
			for i, node := range nodes {
				if move == 'R' {
					nodes[i] = edges[node].right
				} else {
					nodes[i] = edges[node].left
				}
				if nodes[i][2] == 'Z' {
					endsWithZ++
				}
			}
			steps++
			if endsWithZ == len(nodes) {
				break DONE
			}
		}
	}

	return steps
}

func partOne(file *os.File) uint64 {
	var row string
	var commands string
	var edges map[string]Edges = make(map[string]Edges)
	var node string = "AAA"
	var steps uint64 = 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	commands = scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		row = scanner.Text()

		// AAA = (BBB, BBB)
		edges[row[0:3]] = Edges{left: row[7:10], right: row[12:15]}
	}

DONE:
	for {
		for _, v := range commands {
			if v == 'R' {
				node = edges[node].right
			} else {
				node = edges[node].left
			}
			steps++
			if node == "ZZZ" {
				break DONE
			}
		}
	}

	return steps
}
