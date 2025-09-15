package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type axis int
const (
	x axis = iota
	m
	a
	s
	none
)

type operator int
const (
	smaller operator = iota
	greater 
)

type constraint struct {
	ax axis
	op operator
	limit int
	dest string
}

type rating struct {
	xValue, mValue, aValues, sValue int
}

// If there's cycles again, I'm gonna be so mad :)
//
//

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
	var workflows = make(map[string][]constraint, 0)
	var line string
	var splits []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		splits = strings.Split(line, " ")
		mod, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Print(err)
			panic(err)
		}
	}
	return 0
}

func readWorkflow(input string, workflows map[string][]constraint)  {
	keySplits := strings.SplitAfterN(input, "{", 2)
	key, found := strings.CutSuffix(keySplits[0], "{")
	if !found {
		fmt.Errorf("workflow does not begin with XXX{")
	}
	noBrackets := strings.TrimSuffix(keySplits[1], "}")
	workflowSplits := strings.Split(noBrackets, ",")

	// x>2440:R or A
	mapValues := make([]constraint, 0)
	for _, v := range workflowSplits {
		// if A
		if !strings.ContainsRune(v, ':') {
			mapValues = append(mapValues, constraint{ax: none, dest: v})
		} else {
			var ax axis
			switch v[0] {
			case 'x':
				ax = x
			case 'm':
				ax = m
			case 'a':
				ax = a
			case 's':
				ax = s
			}
			var op operator
			switch v[1] {
			case '<':
				op = smaller
			case '>':
				op = greater
			}
			limit, dest, _ := strings.Cut(v[2:], ":")
			limit, err := strconv.Atoi(limit)

			mapValues = append(mapValues, constraint{ax: ax, op: op, limit: limit, })
		}
		// if x>2440:R
	}
	workflows[key] = mapValues
}