package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Edges struct {
	left  string
	right string
}
type Cycle struct {
	node           string
	commands_index uint64
	steps          uint64
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
	var cycles [][]Cycle
	var activeNodes []int
	var steps uint64 = 0
	var lcm uint64

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

	cycles = make([][]Cycle, len(nodes))

	// A slice of visited Z nodes and their distance from the start of the command string
	// If we visit the same node with the same command index we are in a cycle
	// stop iterating
	// If we have a cycle for each start node calculate
	// the lowest common multiple for all combinations end nodes
	// apparently we needed the steps from the start too since that was the period

	for i := 0; i < len(nodes); i++ {
		activeNodes = append(activeNodes, i)
	}

DONE:
	for {
		for commands_index, move := range commands {
			var nodesToRemove []int
			steps++
			for j, index := range activeNodes {
				if move == 'R' {
					nodes[index] = edges[nodes[index]].right
				} else {
					nodes[index] = edges[nodes[index]].left
				}
				if nodes[index][2] == 'Z' {
					if !isNewCycle(cycles[index], nodes[index], uint64(commands_index)) {
						nodesToRemove = append(nodesToRemove, j)
					}
					cycles[index] = append(cycles[index],
						Cycle{node: nodes[index],
							commands_index: uint64(commands_index),
							steps:          steps})
				}
			}
			activeNodes = removeIndices(activeNodes, nodesToRemove)
			if len(activeNodes) == 0 {
				break DONE
			}
		}
	}

	// If we hade more than one Z in the cycle this wouldn't work
	lcm = cycles[0][0].steps
	for i := 1; i < len(cycles); i++ {
		lcm = naiveLCM(lcm, cycles[i][0].steps)
	}

	return lcm
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

func isNewCycle(cycles []Cycle, name string, commands_index uint64) bool {
	result := true
	for _, cycle := range cycles {
		if cycle.node == name && cycle.commands_index == commands_index {
			result = false
		}
	}
	return result
}

func removeIndices(slice []int, indices []int) []int {
	sort.Ints(indices) // Sort indices in ascending order
	for i, index := range indices {
		slice = append(slice[:index-i], slice[index-i+1:]...)
	}
	return slice
}

func naiveGCD(a uint64, b uint64) uint64 {
	if b == 0 {
		return a
	} else {
		return naiveGCD(b, a%b)
	}
}

func naiveLCM(a uint64, b uint64) uint64 {
	return (a * b) / naiveGCD(a, b)
}
