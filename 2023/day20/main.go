package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/eriktedhamre/advent_of_code/utils"
)

type pulse bool

const (
	low  pulse = false
	high pulse = true
)

type system struct {
	modules         map[string]module
	connections     map[string][]string
	signalQueue     utils.Queue[signal]
	sortedNames     []string
	lowSignalsSent  uint64
	highSignalsSent uint64
}

type signal struct {
	p   pulse
	src string
	dst string
}

type snapshot struct {
	lowSignals, highSignals, snapshot uint64
	iteration                         int
}

type module interface {
	process(s *signal, connections map[string][]string) ([]signal, bool)
	snapshot() string
}

type flipFlop struct {
	name  string
	state pulse
}

func (f *flipFlop) process(s *signal, connections map[string][]string) ([]signal, bool) {
	if s.p != low {
		return nil, false

	}

	if s.dst != f.name {
		log.Fatalf("Destination mismatch for module: %+v, got signal: %+v", f, s)
	}

	f.state = !f.state

	return generateOutputSignals(f.name, f.state, connections), true
}

func (f *flipFlop) snapshot() string {
	var snapshot = "0"
	if f.state {
		snapshot = "1"
	}
	return snapshot
}

type conjunction struct {
	name        string
	index       map[string]int
	inputMemory []pulse
}

func (c *conjunction) process(s *signal, connections map[string][]string) ([]signal, bool) {
	if i, ok := c.index[s.src]; ok {
		c.inputMemory[i] = s.p
	}

	var p = low
	for _, v := range c.inputMemory {
		if v != high {
			p = high
		}
	}

	return generateOutputSignals(c.name, p, connections), true
}

func (c *conjunction) snapshot() string {
	var sb strings.Builder

	for _, v := range c.inputMemory {
		if v {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return sb.String()
}

func generateOutputSignals(src string, p pulse, connections map[string][]string) []signal {

	dsts := connections[src]
	signals := make([]signal, 0, len(dsts))
	for _, v := range dsts {
		signals = append(signals, signal{p: p, src: src, dst: v})
	}

	return signals
}

func snapshotHash(modules map[string]module, sortedNames []string) uint64 {
	h := fnv.New64a()

	for _, name := range sortedNames {
		mod := modules[name]
		io.WriteString(h, name)
		io.WriteString(h, mod.snapshot())
	}

	return h.Sum64()
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

	var inputs map[string][]string = make(map[string][]string, 0)
	var moduleWithType []string = make([]string, 0)

	//var flipFlops []flipFlop
	//var conjunctions []conjunction

	s := system{
		modules:     make(map[string]module),
		connections: make(map[string][]string),
		sortedNames: make([]string, 0),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		// ["broadcaster", "a,b,c"]
		// ["%a", "b"]
		// ["&inv", "b,cls,d"]
		src, dest, ok := strings.Cut(line, "->")
		if !ok {
			log.Fatalf("No -> in line %s", line)
		}

		src = strings.TrimSpace(src)
		moduleWithType = append(moduleWithType, src)
		src = strings.Trim(src, "&%")
		dest = strings.TrimSpace(dest)

		dests := strings.Split(dest, ",")
		for i := range dests {
			dests[i] = strings.TrimSpace(dests[i])
		}

		s.connections[src] = dests

		for _, dst := range dests {
			inputs[dst] = append(inputs[dst], src)
		}
	}

	var trimmedName string
	for _, name := range moduleWithType {
		switch {
		case name == "broadcaster":
			// Do nothing
		case strings.HasPrefix(name, "%"):
			trimmedName = strings.Trim(name, "%")
			s.modules[trimmedName] = &flipFlop{name: trimmedName, state: low}
			s.sortedNames = append(s.sortedNames, trimmedName)
		case strings.HasPrefix(name, "&"):
			trimmedName = strings.Trim(name, "&")
			s.sortedNames = append(s.sortedNames, trimmedName)
			srcs := inputs[trimmedName]
			index := make(map[string]int, len(srcs))
			for i, src := range srcs {
				index[src] = i
			}
			inputMemory := make([]pulse, len(srcs))
			s.modules[trimmedName] = &conjunction{
				name:        trimmedName,
				index:       index,
				inputMemory: inputMemory,
			}

		default:
			log.Fatalf("Undefined moduleWithType: %s", name)
		}
	}

	sort.Strings(s.sortedNames)

	//fmt.Println(s)

	var snapshots = make(map[uint64]snapshot, 0)
	var sortedSnaps = make([]snapshot, 0)

	// Add starting state hash
	startState := snapshot{lowSignals: 0,
		highSignals: 0,
		snapshot:    snapshotHash(s.modules, s.sortedNames),
		iteration:   -1}
	snapshots[startState.snapshot] = startState

	const totalIterations uint64 = 1000
	var currentIter snapshot
	var startOfCycle snapshot
	var cycle bool

CYCLE:
	for i := 0; i < 1000; i++ {
		currentIter = runIteration(&s, len(snapshots))

		sortedSnaps = append(sortedSnaps, currentIter)
		startOfCycle, cycle = snapshots[currentIter.snapshot]
		if !cycle {
			snapshots[currentIter.snapshot] = currentIter
		} else {
			break CYCLE
		}
	}

	var startOfCycleIndex int
	var endOfCycleIndex int
	var remainder uint64
	var numberOfCycles uint64
	var cycleLowSignal uint64 = 0
	var cycleHighSignal uint64 = 0
	var totalLowSignals uint64 = 0
	var totalHighSignals uint64 = 0

	// Maybe transform snapsshots into a slice sorted on iteration
	endOfCycleIndex = currentIter.iteration
	startOfCycleIndex = startOfCycle.iteration
	// If Cycle Start == 0
	if startOfCycle.iteration == -1 {
		numberOfCycles = uint64(math.Floor(float64(totalIterations) / float64(endOfCycleIndex)))
		remainder = totalIterations % uint64(endOfCycleIndex)
		for _, v := range sortedSnaps {
			cycleLowSignal += v.lowSignals
			cycleHighSignal += v.highSignals
		}
		totalLowSignals = cycleLowSignal * numberOfCycles
		totalHighSignals = cycleHighSignal * numberOfCycles
		for i := range remainder {
			totalLowSignals += sortedSnaps[i].lowSignals
			totalHighSignals += sortedSnaps[i].highSignals
		}
	} else {
		cycleSize := endOfCycleIndex - startOfCycleIndex
		numberOfCycles = uint64(math.Floor(float64(totalIterations) / float64(cycleSize)))
		remainder = (totalIterations - uint64(startOfCycleIndex)) % uint64(cycleSize)
		// prefix
		for i := 0; i < startOfCycleIndex; i++ {
			totalLowSignals += sortedSnaps[i].lowSignals
			totalHighSignals += sortedSnaps[i].highSignals
		}
		// cycles
		for i := startOfCycleIndex; i < (startOfCycleIndex + cycleSize); i++ {
			cycleLowSignal += sortedSnaps[i].lowSignals
			cycleHighSignal += sortedSnaps[i].highSignals
		}
		totalLowSignals = cycleLowSignal * numberOfCycles
		totalHighSignals = cycleHighSignal * numberOfCycles
		// suffix
		for i := startOfCycleIndex; i < (startOfCycleIndex + int(remainder)); i++ {
			totalLowSignals += sortedSnaps[i].lowSignals
			totalHighSignals += sortedSnaps[i].highSignals
		}
	}
	//	simply calculate number of steps in the cycle
	// 	calculate the remainder
	//  calculate signals for the remainder
	// else
	// 	calculate cycle start
	// 	Add signal counts from non cyclical steps
	// 	calculate number of steps in the cycle
	//  calculate the remainder
	//  calculate signals for the remainder

	return totalHighSignals * totalLowSignals
}

func partTwo(file *os.File) uint64 {
	var line string

	var inputs map[string][]string = make(map[string][]string, 0)
	var moduleWithType []string = make([]string, 0)

	//var flipFlops []flipFlop
	//var conjunctions []conjunction

	s := system{
		modules:     make(map[string]module),
		connections: make(map[string][]string),
		sortedNames: make([]string, 0),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		// ["broadcaster", "a,b,c"]
		// ["%a", "b"]
		// ["&inv", "b,cls,d"]
		src, dest, ok := strings.Cut(line, "->")
		if !ok {
			log.Fatalf("No -> in line %s", line)
		}

		src = strings.TrimSpace(src)
		moduleWithType = append(moduleWithType, src)
		src = strings.Trim(src, "&%")
		dest = strings.TrimSpace(dest)

		dests := strings.Split(dest, ",")
		for i := range dests {
			dests[i] = strings.TrimSpace(dests[i])
		}

		s.connections[src] = dests

		for _, dst := range dests {
			inputs[dst] = append(inputs[dst], src)
		}
	}

	var trimmedName string
	for _, name := range moduleWithType {
		switch {
		case name == "broadcaster":
			// Do nothing
		case strings.HasPrefix(name, "%"):
			trimmedName = strings.Trim(name, "%")
			s.modules[trimmedName] = &flipFlop{name: trimmedName, state: low}
			s.sortedNames = append(s.sortedNames, trimmedName)
		case strings.HasPrefix(name, "&"):
			trimmedName = strings.Trim(name, "&")
			s.sortedNames = append(s.sortedNames, trimmedName)
			srcs := inputs[trimmedName]
			index := make(map[string]int, len(srcs))
			for i, src := range srcs {
				index[src] = i
			}
			inputMemory := make([]pulse, len(srcs))
			s.modules[trimmedName] = &conjunction{
				name:        trimmedName,
				index:       index,
				inputMemory: inputMemory,
			}

		default:
			log.Fatalf("Undefined moduleWithType: %s", name)
		}
	}

	sort.Strings(s.sortedNames)

	var iterationCount uint64 = 1
	var cycleCount = make(map[string][]uint64, 0)
	var keys = make([]string, 0)
	keys = append(keys, "ct", "xc", "kp", "ks")

	for {
		runIterationTwo(&s, iterationCount, cycleCount)
		iterationCount++

		allKeysSat := true

		for _, k := range keys {
			iterations, ok := cycleCount[k]

			if !ok || len(iterations) < 2 {
				allKeysSat = false
				break
			}
		}
		if allKeysSat {
			break
		}
	}

	var answer = cycleCount[keys[0]][0]
	for i := 1; i < len(keys); i++ {
		answer = utils.NaiveLCM(answer, cycleCount[keys[i]][0])
	}

	return answer
}

func runIteration(s *system, iterCount int) snapshot {

	// Add broadcast -> connections signals
	broadcast := generateOutputSignals("broadcaster", low, s.connections)
	for _, signal := range broadcast {
		s.signalQueue.Enqueue(signal)
	}

	s.lowSignalsSent = 0
	s.highSignalsSent = 0

	// Process Signals in SignalQueue

	// Add one low signal sent for signal from button module
	s.lowSignalsSent++

DONE:
	for {
		if s.signalQueue.Len() == 0 {
			break DONE
		}

		curSignal, ok := s.signalQueue.Dequeue()

		if !ok {
			log.Fatalf("s.signalQueue.Dequeue() failed")
		}

		if curSignal.p {
			s.highSignalsSent++
		} else {
			s.lowSignalsSent++
		}

		dstModule, ok := s.modules[curSignal.dst]
		if ok {
			signals, update := dstModule.process(&curSignal, s.connections)

			if update {
				for _, signal := range signals {
					s.signalQueue.Enqueue(signal)
				}
			}
		}
	}

	res := snapshot{lowSignals: s.lowSignalsSent,
		highSignals: s.highSignalsSent,
		iteration:   iterCount,
		snapshot:    snapshotHash(s.modules, s.sortedNames)}

	//fmt.Printf("lowSignalsSent %d, highSignalsSent %d, iteration %d, hash %d\n", res.lowSignals, res.highSignals, res.iteration, res.snapshot)

	return res

	// Calculate Hash

	// Something with the cycles
}

func runIterationTwo(s *system, iteration uint64, cycles map[string][]uint64) {

	// Add broadcast -> connections signals
	broadcast := generateOutputSignals("broadcaster", low, s.connections)
	for _, signal := range broadcast {
		s.signalQueue.Enqueue(signal)
	}

	s.lowSignalsSent = 0
	s.highSignalsSent = 0

DONE:
	for {
		if s.signalQueue.Len() == 0 {
			break DONE
		}

		curSignal, ok := s.signalQueue.Dequeue()

		if !ok {
			log.Fatalf("s.signalQueue.Dequeue() failed")
		}

		if curSignal.dst == "ct" && curSignal.p == low {
			cycles["ct"] = append(cycles["ct"], iteration)
		}

		if curSignal.dst == "xc" && curSignal.p == low {
			cycles["xc"] = append(cycles["xc"], iteration)
		}

		if curSignal.dst == "kp" && curSignal.p == low {
			cycles["kp"] = append(cycles["kp"], iteration)
		}

		if curSignal.dst == "ks" && curSignal.p == low {
			cycles["ks"] = append(cycles["ks"], iteration)
		}

		dstModule, ok := s.modules[curSignal.dst]
		if ok {
			signals, update := dstModule.process(&curSignal, s.connections)

			if update {
				for _, signal := range signals {
					s.signalQueue.Enqueue(signal)
				}
			}
		}
	}
}
