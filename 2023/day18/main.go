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

type instruction struct {
	dir  types.Direction
	dist uint
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
	var splits []string
	//var instructions []instruction = make([]instruction, 0)
	var points []types.Coordinates = make([]types.Coordinates, 0)
	var current types.Coordinates = types.Coordinates{Row: 0, Col: 0}
	points = append(points, types.Coordinates{Row: current.Row, Col: current.Col})

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
			current.Row = current.Row + mod
		case "D":
			current.Row = current.Row - mod
		case "R":
			current.Col = current.Col + mod
		case "L":
			current.Col = current.Col - mod
		}
		points = append(points, types.Coordinates{Row: current.Row, Col: current.Col})
	}
	points = points[:len(points)-1] // not sure if needed
	for _, v := range points {
		fmt.Printf("(%d, %d)\n", v.Row, v.Col)
	}
	printGrid(points)
	return uint64(utils.AreaOfAPolygon(points))
}

func printGrid(points []types.Coordinates) {
	var maxCol int = math.MinInt
	var maxRow int = math.MinInt
	var minCol int = math.MaxInt
	var minRow int = math.MaxInt

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

	fmt.Println("I got so far")

	var rowOffset = 0
	if minRow < 0 {
		rowOffset = int(math.Abs(float64(minRow)))
	}
	var colOffset = 0
	if minCol < 0 {
		colOffset = int(math.Abs(float64(minCol)))
	}

	var grid [][]rune = make([][]rune, maxRow+rowOffset)
	for i := range maxRow + rowOffset {
		grid[i] = make([]rune, maxCol+colOffset)
	}

	for row := 0; row < maxRow+rowOffset; row++ {
		for col := 0; col < maxCol+colOffset; col++ {
			grid[row][col] = '.'
		}
	}

	fmt.Println("but in the end")

	for i := 0; i < len(points)-1; i++ {
		drawLine(&points[i], &points[i+1], grid, rowOffset, colOffset)
	}
	drawLine(&points[len(points)-1], &points[0], grid, rowOffset, colOffset)

	for row := 0; row < maxRow+rowOffset; row++ {
		fmt.Println(grid[row])
	}

}

func drawLine(start, end *types.Coordinates, grid [][]rune, rowOffset, colOffset int) {
	var rowDiff int = 0
	var colDiff int = 0
	var modifier int = 1

	grid[start.Row+rowOffset][start.Col+colOffset] = '#'
	rowDiff = 0
	colDiff = 0
	modifier = 1

	// Could do a combined calculation here :)
	rowDiff = start.Row - end.Row
	if rowDiff != 0 {
		// Step Row Direction
		if rowDiff > 0 {
			// New Row value is lower step with -1
			modifier = -1
		} else {
			// New Row value is higher step with 1
			modifier = 1
		}
		for i := 0; i < int(math.Abs(float64(rowDiff))); i++ {
			grid[start.Row+rowOffset+modifier][start.Col+colOffset] = '#'
			modifier += modifier
		}
	} else {
		// Step Col Direction
		colDiff = start.Col - end.Col
		if colDiff > 0 {
			// new Col value is lower step with -1
			modifier = -1
		} else {
			// new Col value is higher step with 1
			modifier = 1
		}
		for i := 0; i < int(math.Abs(float64(colDiff))); i++ {
			grid[start.Row+rowOffset][start.Col+colOffset+modifier] = '#'
			modifier += modifier
		}
	}

}

// func partTwo(file *os.File) int {
// 	var line string
// 	var grid [][]int = make([][]int, 0)

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line = scanner.Text()
// 		lineSlice, err := utils.StringToIntSlice(line)
// 		if err != nil {
// 			fmt.Print(err)
// 			panic(err)
// 		}
// 		grid = append(grid, lineSlice)
// 	}

// 	return searchGrid2(grid)
// }
