package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eriktedhamre/advent_of_code/utils"
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
	ax    axis
	op    operator
	limit uint64
	dest  string
}

type partRating struct {
	xValue, mValue, aValue, sValue uint64
}

type partLimits struct {
	dest string
	max  partRating
	min  partRating
}

// Start with partLimits 1...4000
// Whenever we make a choice create a branching partLimits
// If we reach the same node with the same part limit we have hit a loop and consider these "R"
// If we reach an A return the passing part limit
// How do we calculate overlap between passing part limits?
// Can we have overlap between passing part limits?

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
	var parts = make([]partRating, 0)
	var line string

	scanner := bufio.NewScanner(file)
	// Read Workflows
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			break
		}

		readWorkflow(line, workflows)
	}

	//fmt.Println(workflows)

	// If number of iterations for a part is larger than the number of parts abort

	for scanner.Scan() {
		line = scanner.Text()
		parts = append(parts, readPart(line))
	}

	//fmt.Println(parts)

	var sum uint64 = 0

	dest := "in"
	steps := 0
	for i := 0; i < len(parts); i++ {
		dest = "in"
		steps = 0
	PART:
		for {
			dest = findDest(parts[i], workflows[dest])
			steps++
			if dest == "A" {
				sum += parts[i].aValue + parts[i].xValue + parts[i].mValue + parts[i].sValue
			}
			if (dest == "R") || (steps > (len(workflows) + 1)) {
				break PART
			}
		}
	}
	return sum
}

func partTwo(file *os.File) uint64 {
	var workflows = make(map[string][]constraint, 0)
	var line string

	scanner := bufio.NewScanner(file)
	// Read Workflows
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			break
		}

		readWorkflow(line, workflows)
	}

	initialPart := partLimits{dest: "in", max: partRating{xValue: 4000, mValue: 4000, aValue: 4000, sValue: 4000},
		min: partRating{xValue: 0, mValue: 0, aValue: 0, sValue: 0}}

	// I will try without cycle detection first
	// With cycle detection I probably need something like this
	// map[string]bool visited
	// and reset whenever I change the current partLimit
	// that's a lot of resets :/

	acceptedParts := make([]partLimits, 0)

	partQueue := &utils.Queue[partLimits]{}
	partQueue.Enqueue(initialPart)

	var currentPart partLimits
	var currentWorkflow []constraint
	var ok bool
	var fieldToModify *uint64

DONE:
	for {
		if partQueue.Len() == 0 {
			break DONE
		}
		currentPart, ok = partQueue.Dequeue()

		if !ok {
			log.Fatalf("partQueue.Dequeue() failed")
		}

		currentWorkflow = workflows[currentPart.dest]

		// cmpr my current part to each value
		// If we do not match
		// For each comparator create both cases
		// continue until spent, whenever we hit an A save it in acceptedParts
		// Whenever we hit an R throw it away
	WORKFLOW:
		for _, c := range currentWorkflow {
			switch c.ax {
			case x:
				if c.op == greater {
					// Early abort
					if c.limit > currentPart.max.xValue {
						continue
					}
					// Create possible parts
					// New part that matches condition
					// Raise min

					newPartKeep := partLimits{dest: currentPart.dest, max: currentPart.max, min: currentPart.min}
					newPartKeep.min.xValue = c.limit + 1

					// New part that fails condition
					// Lower max
					newPartEnqueue := partLimits{dest: currentPart.dest, max: currentPart.max, min: currentPart.min}
					newPartEnqueue.max.xValue = c.limit
				} else {
					// c.op == lesser
					if c.limit < currentPart.min.xValue {
						continue
					}
					// Lower min
					// Raise max

				}

			case m:
			case a:
			case s:
			case none:
				if c.dest == "A" {
					acceptedParts = append(acceptedParts, currentPart)
				} else if c.dest != "R" {
					break WORKFLOW
				} else {
					currentPart.dest = c.dest
					partQueue.Enqueue(currentPart)
				}
			}
		}

	}

	return 0
}
func findDest(part partRating, constraints []constraint) string {
	dest := ""
	var limit uint64 = 0
DONE:
	for _, c := range constraints {
		switch c.ax {
		case x:
			limit = part.xValue
		case m:
			limit = part.mValue
		case a:
			limit = part.aValue
		case s:
			limit = part.sValue
		case none:
			dest = c.dest
			break DONE
		}
		dest = cmprLimit(limit, c.op, c.limit, c)
		if dest != "" {
			break DONE
		}
	}
	return dest
}

func cmprLimit(partValue uint64, op operator, conValue uint64, con constraint) string {
	dest := ""
	switch op {
	case greater:
		if partValue > conValue {
			dest = con.dest
		}
	case smaller:
		if partValue < conValue {
			dest = con.dest
		}
	}
	return dest
}

func readWorkflow(input string, workflows map[string][]constraint) {
	keySplits := strings.SplitAfterN(input, "{", 2)
	key, found := strings.CutSuffix(keySplits[0], "{")
	if !found {
		log.Fatalf("workflow does not begin with XXX{")
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
			limit, dest, found := strings.Cut(v[2:], ":")
			if !found {
				log.Fatalf(": not found for constraint")
				panic(0)
			}
			limitValue, err := strconv.ParseUint(limit, 10, 64)
			if err != nil {
				log.Fatalf("limitValue int conversion failed")
				panic(err)
			}

			mapValues = append(mapValues,
				constraint{ax: ax, op: op, limit: limitValue, dest: dest})
		}
	}
	workflows[key] = mapValues
}

func readPart(input string) partRating {
	splits := strings.Split(input, ",")
	splits[0] = strings.TrimPrefix(splits[0], "{")
	splits[3] = strings.TrimSuffix(splits[3], "}")
	part := partRating{}
	for _, v := range splits {
		val, err := strconv.ParseUint(v[2:], 10, 64)
		if err != nil {
			log.Fatalf("part axis int conversion failed")
		}
		switch v[0] {
		case 'x':
			part.xValue = val
		case 'm':
			part.mValue = val
		case 'a':
			part.aValue = val
		case 's':
			part.sValue = val
		}
	}
	return part

}
