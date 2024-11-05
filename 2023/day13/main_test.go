package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_partOne(t *testing.T) {

	tests := []struct {
		filepath string
		expected uint64
	}{
		{"example_1.txt", 400},
		{"example_1_rev.txt", 300},
		{"example.txt", 405},
	}

	for _, tt := range tests {
		file, err := os.Open(tt.filepath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		result := partOne(file)
		if result != tt.expected {
			t.Errorf("Expected: %d, got: %d, file: %s",
				tt.expected, result, tt.filepath)
		}
	}
}

func TestCheckSpan(t *testing.T) {
	tests := []struct {
		expected bool
		top      int
		bottom   int
		grid     []string
	}{
		{true, 1, 6,
			[]string{"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#"},
		},
		{true, 0, 5,
			[]string{"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
				"#...##..#"},
		},
	}

	for _, tt := range tests {
		result := checkSpan(tt.grid, tt.bottom, tt.top)
		if result != tt.expected {
			t.Errorf("Expected: %t, got: %t,",
				tt.expected, result)
		}
	}
}
func TestCheckHorizontal(t *testing.T) {
	tests := []struct {
		expected int
		grid     []string
	}{
		{3,
			[]string{"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
				"#...##..#"},
		},
	}

	for _, tt := range tests {
		result := checkHorizontal(tt.grid)
		if result != tt.expected {
			t.Errorf("Expected: %d, got: %d,",
				tt.expected, result)
		}
	}
}

func TestMatrixRotate(t *testing.T) {
	tests := []struct {
		grid    []string
		rotated []string
	}{
		{
			[]string{"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
				"#...##..#"},

			[]string{"##.##.#",
				"...##..",
				"..####.",
				"..####.",
				"#..##..",
				"##....#",
				"..####.",
				"..####.",
				"###..##"},
		},
	}

	for _, tt := range tests {
		result := matrixRotate(tt.grid)
		for i, row := range tt.rotated {
			for j := range row {
				if result[i][j] != tt.rotated[i][j] {
					t.Errorf("elem[%d][%d] mismatch", i, j)
				}
			}
		}
	}
}

// func Test_partTwo(t *testing.T) {

// 	tests := []struct {
// 		filepath string
// 		expected uint64
// 	}{
// 		{"example.txt", 525152},
// 	}

// 	for _, tt := range tests {
// 		file, err := os.Open(tt.filepath)
// 		if err != nil {
// 			fmt.Println("Error opening file:", err)
// 			return
// 		}
// 		defer file.Close()
// 		result := partTwo(file)
// 		if result != tt.expected {
// 			t.Errorf("Expected: %d, got: %d, file: %s",
// 				tt.expected, result, tt.filepath)
// 		}
// 	}
// }
