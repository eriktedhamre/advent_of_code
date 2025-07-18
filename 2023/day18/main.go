package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/eriktedhamre/advent_of_code/types"
	"github.com/eriktedhamre/advent_of_code/utils"
)

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
	var splits []string
	//var instructions []instruction = make([]instruction, 0)
	var points []types.Coordinates = make([]types.Coordinates, 0)
	var current types.Coordinates = types.Coordinates{Row: 0, Col: 0}
	points = append(points, types.Coordinates{Row: current.Row, Col: current.Col})
	var boundaryPoints uint64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		splits = strings.Split(line, " ")
		mod, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Print(err)
			panic(err)
		}
		switch splits[0] {
		case "U":
			current.Row = current.Row + int64(mod)
		case "D":
			current.Row = current.Row - int64(mod)
		case "R":
			current.Col = current.Col + int64(mod)
		case "L":
			current.Col = current.Col - int64(mod)
		}
		boundaryPoints += uint64(mod)
		points = append(points, types.Coordinates{Row: current.Row, Col: current.Col})
	}
	points = points[:len(points)-1] // not sure if needed
	boundaryPoints += uint64(math.Abs(float64(current.Row))) + uint64(math.Abs(float64(current.Col)))
	// for _, v := range points {
	// 	fmt.Printf("(%d, %d)\n", v.Row, v.Col)
	// }

	// I realized that my solution would be impossible
	// for the second part due to the magnitude of the grid.
	// Apparently you can use Picket's theorem if you think a bit :/
	// A = i + b/2 - 1
	// i = A - b/2 + 1
	// Sought = i + b

	A := utils.AreaOfAPolygon(points)
	i := A - float64(boundaryPoints)/2 + 1
	return boundaryPoints + uint64(i)
}

func partTwo(file *os.File) uint64 {
	var line string
	var hexadecimalString string
	var splits []string
	var points []types.Coordinates = make([]types.Coordinates, 0)
	var current types.Coordinates = types.Coordinates{Row: 0, Col: 0}
	points = append(points, types.Coordinates{Row: current.Row, Col: current.Col})
	var boundaryPoints uint64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		splits = strings.Split(line, " ")
		// (#70c710)
		hexadecimalString = splits[2]

		mod, err := strconv.ParseUint(hexadecimalString[2:7], 16, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return 1
		}

		switch hexadecimalString[7] {
		case '3':
			current.Row = current.Row + int64(mod)
		case '1':
			current.Row = current.Row - int64(mod)
		case '0':
			current.Col = current.Col + int64(mod)
		case '2':
			current.Col = current.Col - int64(mod)
		}
		boundaryPoints += uint64(mod)
		points = append(points, types.Coordinates{Row: current.Row, Col: current.Col})
	}
	points = points[:len(points)-1] // not sure if needed
	boundaryPoints += uint64(math.Abs(float64(current.Row))) + uint64(math.Abs(float64(current.Col)))
	// for _, v := range points {
	// 	fmt.Printf("(%d, %d)\n", v.Row, v.Col)
	// }

	// I realized that my solution would be impossible
	// for the second part due to the magnitude of the grid.
	// Apparently you can use Picket's theorem if you think a bit :/
	// A = i + b/2 - 1
	// i = A - b/2 + 1
	// Sought = i + b

	A := utils.AreaOfAPolygon(points)
	i := A - float64(boundaryPoints)/2 + 1
	return boundaryPoints + uint64(i)

}

func createGrid(points []types.Coordinates) [][]rune {
	var maxCol int64 = math.MinInt
	var maxRow int64 = math.MinInt
	var minCol int64 = math.MaxInt
	var minRow int64 = math.MaxInt

	for i := range len(points) - 1 {
		if maxCol < points[i].Col {
			maxCol = points[i].Col
		}
		if maxRow < points[i].Row {
			maxRow = points[i].Row
		}
		if minCol > points[i].Col {
			minCol = points[i].Col
		}
		if minRow > points[i].Row {
			minRow = points[i].Row
		}
	}

	var rowOffset int64 = 0
	if minRow < 0 {
		rowOffset = int64(math.Abs(float64(minRow)))
	}
	var colOffset int64 = 0
	if minCol < 0 {
		colOffset = int64(math.Abs(float64(minCol)))
	}

	var grid [][]rune = make([][]rune, maxRow+rowOffset+1)
	for i := range maxRow + rowOffset + 1 {
		grid[i] = make([]rune, maxCol+colOffset+1)
	}

	for row := int64(0); row < maxRow+rowOffset+1; row++ {
		for col := int64(0); col < maxCol+colOffset+1; col++ {
			grid[row][col] = '.'
		}
	}

	for i := 0; i < len(points)-1; i++ {
		drawLine(&points[i], &points[i+1], grid, rowOffset, colOffset)
	}
	drawLine(&points[len(points)-1], &points[0], grid, rowOffset, colOffset)

	return grid
}

// Flood-fill from all '.' cells connected to the grid edges
func floodFillOutside(grid [][]rune, target rune, marker rune) uint64 {
	rows := int64(len(grid))
	cols := int64(len(grid[0]))
	directions := []types.Coordinates{{Row: -1, Col: 0}, {Row: 1, Col: 0}, {Row: 0, Col: -1}, {Row: 0, Col: 1}}

	queue := []types.Coordinates{}

	// Start from borders
	for r := int64(0); r < rows; r++ {
		if grid[r][0] == target {
			queue = append(queue, types.Coordinates{Row: r, Col: 0})
		}
		if grid[r][cols-1] == target {
			queue = append(queue, types.Coordinates{Row: r, Col: cols - 1})
		}
	}
	for c := int64(0); c < cols; c++ {
		if grid[0][c] == target {
			queue = append(queue, types.Coordinates{Row: 0, Col: c})
		}
		if grid[rows-1][c] == target {
			queue = append(queue, types.Coordinates{Row: rows - 1, Col: c})
		}
	}

	var cellsFilled uint64 = 0
	// Flood all edge-connected '.'
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if grid[p.Row][p.Col] != target {
			continue
		}

		grid[p.Row][p.Col] = marker
		cellsFilled++

		for _, d := range directions {
			nr := p.Row + d.Row
			nc := p.Col + d.Col
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == target {
				queue = append(queue, types.Coordinates{Row: nr, Col: nc})
			}
		}
	}
	return cellsFilled
}

func drawLine(start, end *types.Coordinates, grid [][]rune, rowOffset, colOffset int64) {
	var rowDiff int64 = 0
	var colDiff int64 = 0
	var modifier int64 = 0
	var modifier_step int64 = 1

	rowDiff = 0
	colDiff = 0
	modifier = 0

	// Could do a combined calculation here :)
	rowDiff = start.Row - end.Row
	if rowDiff != 0 {
		// Step Row Direction
		if rowDiff > 0 {
			// New Row value is lower step with -1
			modifier_step = -1
		} else {
			// New Row value is higher step with 1
			modifier_step = 1
		}
		for i := 0; i <= int(math.Abs(float64(rowDiff))); i++ {
			grid[start.Row+rowOffset+modifier][start.Col+colOffset] = '#'
			modifier += modifier_step
		}
	} else {
		// Step Col Direction
		colDiff = start.Col - end.Col
		if colDiff > 0 {
			// new Col value is lower step with -1
			modifier_step = -1
		} else {
			// new Col value is higher step with 1
			modifier_step = 1
		}
		for i := 0; i <= int(math.Abs(float64(colDiff))); i++ {
			grid[start.Row+rowOffset][start.Col+colOffset+modifier] = '#'
			modifier += modifier_step
		}
	}

}

func printGrid(grid [][]rune) {
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Println(string(grid[i]))
	}
}
