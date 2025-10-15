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
	lesser operator = iota
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

	for scanner.Scan() {
		line = scanner.Text()
		parts = append(parts, readPart(line))
	}

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
		min: partRating{xValue: 1, mValue: 1, aValue: 1, sValue: 1}}

	acceptedParts := make([]partLimits, 0)

	partQueue := &utils.Queue[partLimits]{}
	partQueue.Enqueue(initialPart)

	var currentPart partLimits
	var currentWorkflow []constraint
	var ok bool

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
			newPartEnqueue := partLimits{dest: c.dest, max: currentPart.max, min: currentPart.min}
			newPartKeep := partLimits{dest: currentPart.dest, max: currentPart.max, min: currentPart.min}

			if c.ax == none {
				switch c.dest {
				case "A":
					acceptedParts = append(acceptedParts, currentPart)
				case "R":
					break WORKFLOW
				default:
					currentPart.dest = c.dest
					partQueue.Enqueue(currentPart)
				}
				continue
			}

			minVal := getMin(&currentPart, c.ax)
			maxVal := getMax(&currentPart, c.ax)

			switch c.op {
			case greater:
				// condition: value > limit
				if c.limit >= maxVal {
					continue
				}
				setMin(&newPartEnqueue, c.ax, c.limit+1)
				setMax(&newPartKeep, c.ax, c.limit)

			case lesser:
				// condition: value < limit
				if c.limit <= minVal {
					continue
				}
				setMax(&newPartEnqueue, c.ax, c.limit-1)
				setMin(&newPartKeep, c.ax, c.limit)
			}

			// Handle destinations
			if c.dest == "A" {
				acceptedParts = append(acceptedParts, newPartEnqueue)
			} else if c.dest != "R" {
				partQueue.Enqueue(newPartEnqueue)
			}

			currentPart = newPartKeep
		}
	}

	var sum uint64 = 0
	for _, v := range acceptedParts {
		sum += (v.max.xValue - v.min.xValue + 1) * (v.max.mValue - v.min.mValue + 1) * (v.max.aValue - v.min.aValue + 1) * (v.max.sValue - v.min.sValue + 1)
	}
	return sum
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
	case lesser:
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
				op = lesser
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

func getMin(p *partLimits, ax axis) uint64 {
	switch ax {
	case x:
		return p.min.xValue
	case m:
		return p.min.mValue
	case a:
		return p.min.aValue
	case s:
		return p.min.sValue
	default:
		return 0
	}
}

func getMax(p *partLimits, ax axis) uint64 {
	switch ax {
	case x:
		return p.max.xValue
	case m:
		return p.max.mValue
	case a:
		return p.max.aValue
	case s:
		return p.max.sValue
	default:
		return 0
	}
}

func setMin(p *partLimits, ax axis, val uint64) {
	switch ax {
	case x:
		p.min.xValue = val
	case m:
		p.min.mValue = val
	case a:
		p.min.aValue = val
	case s:
		p.min.sValue = val
	}
}

func setMax(p *partLimits, ax axis, val uint64) {
	switch ax {
	case x:
		p.max.xValue = val
	case m:
		p.max.mValue = val
	case a:
		p.max.aValue = val
	case s:
		p.max.sValue = val
	}
}
