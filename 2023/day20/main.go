package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"
	"log"
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
	lowSignals, highSignals, iteration uint64
	snapshot                           uint64
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
	fmt.Print(partOne(file))
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
			s.modules[trimmedName] = &flipFlop{name: name, state: low}
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

	fmt.Println(s)
	return 0
}
