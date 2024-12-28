package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"

	"github.com/eriktedhamre/advent_of_code/types"
	"github.com/eriktedhamre/advent_of_code/utils"
)

type Pos struct {
	row, col int
	dir      types.Direction
}

type DirMod struct {
	dir            types.Direction
	rowMod, colMod int
}

var (
	UpMod    = DirMod{types.Up, -1, 0}
	DownMod  = DirMod{types.Down, 1, 0}
	LeftMod  = DirMod{types.Left, 0, -1}
	RightMod = DirMod{types.Right, 0, 1}
)

var LeftRightMap = map[types.Direction][]DirMod{
	types.Up:    {LeftMod, RightMod},
	types.Down:  {LeftMod, RightMod},
	types.Left:  {UpMod, DownMod},
	types.Right: {UpMod, DownMod},
}

var ForwardMap = map[types.Direction]DirMod{
	types.Up:    UpMod,
	types.Down:  DownMod,
	types.Left:  LeftMod,
	types.Right: RightMod,
}

var InQueue = make(map[string]bool, 0)

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
	var lowestHeat int
	//var best *utils.Item[types.Crucible]
	pq := make(utils.PriorityQueue[types.Crucible], 0)
	pq.Init()

	heap.Push(&pq, &utils.Item[types.Crucible]{
		Value: types.Crucible{
			Row:              0,
			Col:              0,
			ConsecutiveMoves: 0,
			Direction:        types.Down,
			//Parents:          []types.Coordinates{{Row: 0, Col: 0}},
		},
		Priority: 0})

	InQueue[fmt.Sprintf("%d%d%d%d%d", 0, 0, 0, int(types.Down), 0)] = true

DONE:
	for {
		// We exit when we reach the goal
		curItem, ok := heap.Pop(&pq).(*utils.Item[types.Crucible])
		//fmt.Println(curItem.Priority)
		if !ok {
			fmt.Printf("type conversion failed for curItem : got %T\n", pq.Pop())
			panic(nil)
		}

		if curItem.Value.Row == (len(cost)-1) && curItem.Value.Col == (len(grid[0])-1) {
			lowestHeat = curItem.Priority
			//best = curItem
			break DONE
		}
		newIndicies := calculateNewPosSlice(curItem, len(cost), len(cost[0]))
		for _, v := range newIndicies {
			// We are moving forward, not cash money solution
			if v.dir == curItem.Value.Direction {
				newConsec = curItem.Value.ConsecutiveMoves + 1
			} else {
				newConsec = 1
			}
			//newDir = calcNewDir(curItem.Value.Direction, i)

			key := fmt.Sprintf("%d%d%d%d%d", v.row, v.col, newConsec, v.dir, curItem.Priority+grid[v.row][v.col])

			if _, exists := InQueue[key]; exists {
				// Item is already in the queue
				continue
			}

			// newParents := deepCopy(curItem.Value.Parents)
			// newParents = append(newParents, types.Coordinates{Row: v.row, Col: v.col})

			// if currentHeat + cost is lower than the previous cost add it to priority queue
			// Can this destroy the best solution by locking us out of it?
			// Don't think so, it is possible that we should simply add everything to the queue
			// I'm adding everything :)
			heap.Push(&pq, &utils.Item[types.Crucible]{
				Value: types.Crucible{
					Row:              v.row,
					Col:              v.col,
					ConsecutiveMoves: newConsec,
					Direction:        v.dir,
					//Parents:          newParents,
				},
				Priority: curItem.Priority + grid[v.row][v.col],
			})
			InQueue[key] = true
		}
	}

	// for _, v := range best.Value.Parents {
	// 	grid[v.Row][v.Col] = 0
	// }
	// for i, row := range grid {
	// 	for j := range row {
	// 		fmt.Print(grid[i][j])
	// 	}
	// 	fmt.Print("\n")
	// }

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
			res = append(res, Pos{newRow, newCol, v.dir})
		}
	}

	return res
}

func calcNewDir(dir types.Direction, index int) types.Direction {
	var res types.Direction
	switch {
	case index == 2:
		res = dir
	// Left or Up
	case index == 0:
		if dir == types.Up || dir == types.Down {
			res = types.Left
		} else {
			res = types.Up
		}
	// Right or Down
	case index == 1:
		if dir == types.Up || dir == types.Down {
			res = types.Right
		} else {
			res = types.Down
		}
	}
	return res
}

func deepCopy(parents []types.Coordinates) []types.Coordinates {
	newParents := make([]types.Coordinates, len(parents))
	copy(newParents, parents)
	return newParents
}
