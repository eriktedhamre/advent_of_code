package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/eriktedhamre/advent_of_code/types"
	"github.com/eriktedhamre/advent_of_code/utils"
)

type Pos struct {
	row, col int
}

type DirMod struct {
	rowMod, colMod int
}

var (
UpMod = DirMod{-1, 0}
DownMod = DirMod{1, 0}
LeftMod = DirMod{0, -1}
RightMod = DirMod{0, 1}
)

var LeftRightMap = map[types.Direction][]DirMod {
	types.Up : {LeftMod, RightMod},
	types.Down : {LeftMod, RightMod},
	types.Left : {UpMod, DownMod},
	types.Right : {UpMod, DownMod},
}

var ForwardMap = map[types.Direction]DirMod {
	types.Up: UpMod,
	types.Down: DownMod,
	types.Left: LeftMod,
	types.Right: RightMod,
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

func partOne(file *os.File) int {
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

	return searchGrid(grid, cost)
}

func searchGrid(grid [][]int, cost [][]int) int {
	//var curItem utils.Item[Crucible]
	var newConsec int
	var newDir types.Direction
	var lowestHeat int
	pq := make(utils.PriorityQueue[types.Crucible], 0)
	pq.Init()

	pq.Push(&utils.Item[types.Crucible]{
		Value: types.Crucible{
			Row:              0,
			Col:              0,
			ConsecutiveMoves: 0,
			Direction:        types.Down,
		},
		Priority: 0})

DONE:
	for {
		// We exit when we reach the goal
		curItem, ok := pq.Pop().(*utils.Item[types.Crucible])
		if !ok {
			fmt.Printf("type conversion failed for curItem : got %T\n", pq.Pop())
			panic(nil)
		}

		if curItem.Value.Row == (len(cost) - 1) && curItem.Value.Col == (len(grid[0]) - 1) {
			lowestHeat = curItem.Priority
			break DONE
		}
		newIndicies := calculateNewPosSlice(curItem, len(cost), len(cost[0]))
		newConsec = 0
		for i, v := range newIndicies {
			// We are moving forward, not cash money solution
			if i == 2 {
				newConsec = curItem.Value.ConsecutiveMoves + 1
			}
			newDir = calcNewDir(curItem.Value.Direction, i)
			// if currentHeat + cost is lower than the previous cost add it to priority queue
			// Can this destroy the best solution by locking us out of it?
			// Don't think so, it is possible that we should simply add everything to the queue
			// I'm adding everything :)
			pq.Push(&utils.Item[types.Crucible]{
				Value: types.Crucible{
					Row: v.row,
					Col: v.col,
					ConsecutiveMoves: newConsec,
					Direction: newDir,
				},
				Priority: curItem.Priority + grid[v.row][v.col],
			})
		}
	}

	return lowestHeat
}

func calculateNewPosSlice(current *utils.Item[types.Crucible], rows, cols int) []Pos {
	// Look at consecutive moves and see what turns are allowed
	// Map direction we came from to allowed turns
	// See if new position is InBound
	var res []Pos = make([]Pos, 0)
	var mods []DirMod = make([]DirMod, 0)
	var newRow int
	var newCol int
	mods = append(mods, LeftRightMap[current.Value.Direction]...)

	// Add Forward if consecutiveMoves < 3
	if current.Value.ConsecutiveMoves != 3 {
		mods = append(mods, ForwardMap[current.Value.Direction])	
	}

	for _, v := range mods {
		newRow = current.Value.Row + v.rowMod
		newCol = current.Value.Col + v.colMod
		if utils.InBound(newRow, newCol, rows, cols) {
			res = append(res, Pos{newRow, newCol})
		}
	}

	return res
}

func calcNewDir(dir types.Direction, index int) types.Direction {
	var res types.Direction
	switch {
	case index == 2:
		res =  dir
	// Left or Up
	case index == 0:
		if dir == types.Up || dir == types.Down{
			res =  types.Left
		} else {
			res = types.Up
		}
	// Right or Down
	case index == 1:
		if dir == types.Up || dir == types.Down{
			res = types.Right
		} else {
			res = types.Down
		}
	}
	return res
}
