package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eriktedhamre/advent_of_code/types"
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

func partOne(file *os.File) int {
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
	// points = points[:len(points)-1] // not sure if needed
	for _, v := range points {
		fmt.Printf("(%d, %d)\n", v.Row, v.Col)
	}
	return 0
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
