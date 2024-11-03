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
		{"example_2.txt", 400},
		{"example_2_rev.txt", 400},
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

func TestCheckSpan(t *testing.T)  {
	tests := []struct {
		expected bool
		top int
		bottom int
		grid []string
	}{
		{true, 1, 6,
		[]string{"#...##..#",
                 "#....#..#",
                 "..##..###",
                 "#####.##.",
                 "#####.##.",
                 "..##..###",
                 "#....#..#",},
		},
		{false, 6, 5,
			[]string{"#....#..#",
					 "..##..###",
					 "#####.##.",
					 "#####.##.",
					 "..##..###",
					 "#....#..#",
					 "#...##..#",},
			},
	}

	for _, tt := range tests {
		result := checkSpan(tt.grid, tt.bottom, tt.top)
		if result != tt.expected {
			t.Errorf("Expected: %t, got: %t,",
				tt.expected, result,)
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
