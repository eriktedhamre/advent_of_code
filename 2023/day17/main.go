package main

import (
	"bufio"
	"container/heap"
	"fmt"
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

var bestCost = make(map[string]int, 0)

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

func partOne(file *os.File) int {
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

	return searchGrid(grid)
}

func partTwo(file *os.File) int {
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

	return searchGrid2(grid)
}

func searchGrid(grid [][]int) int {

	var newConsec int
	var lowestHeat int
	pq := make(utils.PriorityQueue[types.Crucible], 0)
	pq.Init()

	heap.Push(&pq, &utils.Item[types.Crucible]{
		Value: types.Crucible{
			Row:              0,
			Col:              0,
			ConsecutiveMoves: 0,
			Direction:        types.Down,
		},
		Priority: 0})

	bestCost[fmt.Sprintf("%d,%d,%d,%d", 0, 0, 0, int(types.Down))] = 0

DONE:
	for {
		// We exit when we reach the goal
		curItem, ok := heap.Pop(&pq).(*utils.Item[types.Crucible])
		if !ok {
			fmt.Printf("type conversion failed for curItem : got %T\n", curItem)
			panic(nil)
		}

		if curItem.Value.Row == (len(grid)-1) && curItem.Value.Col == (len(grid[0])-1) {
			lowestHeat = curItem.Priority
			break DONE
		}
		newIndicies := calculateNewPosSlice(curItem, len(grid), len(grid[0]))
		for _, v := range newIndicies {
			// We are moving forward, not cash money solution
			if v.dir == curItem.Value.Direction {
				newConsec = curItem.Value.ConsecutiveMoves + 1
			} else {
				newConsec = 1
			}

			key := fmt.Sprintf("%d,%d,%d,%d", v.row, v.col, newConsec, v.dir)
			newCost := curItem.Priority + grid[v.row][v.col]

			if prevCost, exists := bestCost[key]; !exists || (newCost < prevCost) {
				bestCost[key] = newCost
				heap.Push(&pq, &utils.Item[types.Crucible]{
					Value: types.Crucible{
						Row:              v.row,
						Col:              v.col,
						ConsecutiveMoves: newConsec,
						Direction:        v.dir,
					},
					Priority: newCost,
				})
			}

		}
	}

	return lowestHeat
}

func searchGrid2(grid [][]int) int {

	var newConsec int
	var lowestHeat int
	pq := make(utils.PriorityQueue[types.Crucible], 0)
	pq.Init()

	heap.Push(&pq, &utils.Item[types.Crucible]{
		Value: types.Crucible{
			Row:              0,
			Col:              0,
			ConsecutiveMoves: 0,
			Direction:        types.Down,
		},
		Priority: 0})
	bestCost[fmt.Sprintf("%d,%d,%d,%d", 0, 0, 0, int(types.Down))] = 0

	heap.Push(&pq, &utils.Item[types.Crucible]{
		Value: types.Crucible{
			Row:              0,
			Col:              0,
			ConsecutiveMoves: 0,
			Direction:        types.Right,
		},
		Priority: 0})
	bestCost[fmt.Sprintf("%d,%d,%d,%d", 0, 0, 0, int(types.Right))] = 0

DONE:
	for {
		// We exit when we reach the goal
		curItem, ok := heap.Pop(&pq).(*utils.Item[types.Crucible])
		if !ok {
			fmt.Printf("type conversion failed for curItem : got %T\n", curItem)
			panic(nil)
		}

		// We need to have atleast 4 consecutive moves to stop
		if curItem.Value.Row == (len(grid)-1) &&
			curItem.Value.Col == (len(grid[0])-1) &&
			curItem.Value.ConsecutiveMoves > 3 {
			lowestHeat = curItem.Priority
			break DONE
		}
		newIndicies := calculateNewPosSlice2(curItem, len(grid), len(grid[0]))
		for _, v := range newIndicies {
			// We are moving forward, not cash money solution
			if v.dir == curItem.Value.Direction {
				newConsec = curItem.Value.ConsecutiveMoves + 1
			} else {
				newConsec = 1
			}

			key := fmt.Sprintf("%d,%d,%d,%d", v.row, v.col, newConsec, v.dir)
			newCost := curItem.Priority + grid[v.row][v.col]

			if prevCost, exists := bestCost[key]; !exists || (newCost < prevCost) {
				bestCost[key] = newCost
				heap.Push(&pq, &utils.Item[types.Crucible]{
					Value: types.Crucible{
						Row:              v.row,
						Col:              v.col,
						ConsecutiveMoves: newConsec,
						Direction:        v.dir,
					},
					Priority: newCost,
				})
			}

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
			res = append(res, Pos{newRow, newCol, v.dir})
		}
	}

	return res
}

func calculateNewPosSlice2(current *utils.Item[types.Crucible], rows, cols int) []Pos {
	// Look at consecutive moves and see what turns are allowed
	// Map direction we came from to allowed turns
	// See if new position is InBound
	var res []Pos = make([]Pos, 0)
	var mods []DirMod = make([]DirMod, 0)
	var newRow int
	var newCol int

	if current.Value.ConsecutiveMoves < 10 {
		mods = append(mods, ForwardMap[current.Value.Direction])
	}

	if current.Value.ConsecutiveMoves > 3 {
		mods = append(mods, LeftRightMap[current.Value.Direction]...)
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
