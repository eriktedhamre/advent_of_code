package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eriktedhamre/advent_of_code/utils"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type DirMod struct {
	rowwMod, colMod int
}

var dirMods = []DirMod{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Crucible struct {
	row, col         int
	consecutiveMoves int
	origin           Direction
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

// Unique for a state
// position in grid
// current heat accumulated
// consecutive moves
// origin

// I will try large grid search with cache again....
// If accumulated heat in the new positions is higher than the current one
// and the other stats are the same/worse do not add it

// Or are we supposed to do a recursive solution

// Priority queue on heat??

func partOne(file *os.File) uint64 {
	var line string
	var grid [][]int = make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		lineSlice, err := utils.StringToIntSlice(line)
		if err != nil {
			fmt.Print(err)
			panic(err)
		}
		grid = append(grid, lineSlice)
	}

	return 0
}

func searchGrid(grid [][]int) {
	pq := make(utils.PriorityQueue[utils.Item[Crucible]], 0)
	pq.Init()

	pq.Push(utils.Item[Crucible]{Value: Crucible{row: 0, col: 0, consecutiveMoves: 0, origin: Up}, Priority: 0})

}
