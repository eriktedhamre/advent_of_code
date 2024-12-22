package main

import (
	"bufio"
	"fmt"
	"math"
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

type Pos struct {
	row, col int
}

type DirMod struct {
	rowwMod, colMod int
}

var (
UpMod = DirMod{-1, 0}
DownMod = DirMod{1, 0}
LeftMod = DirMod{0, -1}
RightMod = DirMod{0, 1}
)

var dirMods = []DirMod{UpMod, DownMod, LeftMod, RightMod}

type Crucible struct {
	row, col         int
	consecutiveMoves int
	direction           Direction
}

var myMap = map[Direction][]DirMod {
	Up : {LeftMod, RightMod},
	Down : {LeftMod, RightMod},
	Left : {UpMod, DownMod},
	Right : {UpMod, DownMod},
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
	var cost [][]int

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

	cost = make([][]int, len(grid))
	cols := len(grid[0])

	for i := range cost {
		cost[i] = make([]int, cols)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}

	return 0
}

func searchGrid(grid [][]int, cost [][]int) {
	//var curItem utils.Item[Crucible]
	var rows int = len(grid)
	var cols int = len(grid[0])
	var newRow int
	var newCol int
	pq := make(utils.PriorityQueue[utils.Item[Crucible]], 0)
	pq.Init()

	pq.Push(utils.Item[Crucible]{
		Value: Crucible{
			row:              0,
			col:              0,
			consecutiveMoves: 0,
			direction:        Up},
		Priority: 0})

	for {
		// We exit when we reach the goal
		curItem, ok := pq.Pop().(utils.Item[Crucible])
		if !ok {
			fmt.Printf("type conversion failed for curItem")
			panic(nil)
		}

		for _, v := range dirMods {
			newRow = curItem.Value.row + v.rowwMod
			newCol = curItem.Value.col + v.colMod

			
			// Check if the heat for the new pos is lower than currently recorded

			if utils.InBound(newRow, newCol, rows, cols) {
				
			}
		}


	}

	

}

func calculateNewPosSlice(current *utils.Item[Crucible], cost [][]int) []Pos {
	// Look at consecutive moves and see what turns are allowed
	// Map direction we came from to allowed turns
	// See if new position is InBound

	// Add Left

	// Add Right

	// Add Forward if consecutiveMoves < 3
	if current.Value.consecutiveMoves == 3 {
		
	} else {

	}
}
