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
		{"example.txt", 21},
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

func TestMatches(t *testing.T) {

	tests := []struct {
		springs    []rune
		conditions []int
		expected   uint64
	}{
		{[]rune{'#', '.', '#', '.', '#', '#', '#'}, []int{1, 1, 3}, 1},
		{[]rune{'#', '.', '#', '.', '#', '.', '#'}, []int{1, 1, 3}, 0},
		{[]rune{'.', '.', '.', '.', '.', '.', '.'}, []int{1, 1, 3}, 0},
		{[]rune{'#', '#', '#', '#', '#', '#', '#'}, []int{1, 1, 3}, 0},
		{[]rune{'.', '#', '#', '#', '.', '#', '#', '.', '#', '.', '.', '.'}, []int{3, 2, 1}, 1},
		{[]rune{'.', '#', '#', '#', '.', '#', '#', '.', '.', '#', '.', '.'}, []int{3, 2, 1}, 1},
		{[]rune{'.', '#', '#', '#', '.', '#', '#', '.', '.', '.', '#', '.'}, []int{3, 2, 1}, 1},
		{[]rune{'.', '#', '#', '#', '.', '#', '#', '.', '.', '.', '.', '#'}, []int{3, 2, 1}, 1},
		{[]rune{'.', '#', '#', '#', '.', '.', '#', '#', '.', '#', '.', '.'}, []int{3, 2, 1}, 1},
		{[]rune{'.', '#', '#', '#', '.', '.', '#', '#', '.', '.', '#', '.'}, []int{3, 2, 1}, 1},
		{[]rune{'.', '#', '#', '#', '.', '.', '#', '#', '.', '.', '.', '#'}, []int{3, 2, 1}, 1},
		{[]rune{'.', '#', '#', '#', '.', '.', '.', '#', '#', '.', '#', '.'}, []int{3, 2, 1}, 1},
		{[]rune{'.', '#', '#', '#', '.', '.', '.', '#', '#', '.', '.', '#'}, []int{3, 2, 1}, 1},
		{[]rune{'.', '#', '#', '#', '.', '.', '.', '.', '#', '#', '.', '#'}, []int{3, 2, 1}, 1},
	}

	for i, tt := range tests {

		result := matches(tt.springs, tt.conditions)
		if result != tt.expected {
			t.Errorf("Expected: %d, got: %d Test: %d",
				tt.expected, result, i)
		}
	}
}

func TestSolver(t *testing.T) {

	tests := []struct {
		springs    []rune
		conditions []int
		expected   uint64
	}{
		{[]rune{'#', '?', '#', '.', '#', '#', '#'}, []int{1, 1, 3}, 1},
		{[]rune{'?', '.', '#', '.', '#', '#', '#'}, []int{1, 1, 3}, 1},
		{[]rune{'?', '.', '?', '.', '#', '#', '#'}, []int{1, 1, 3}, 1},

		{[]rune{'?', '?', '?', '.', '#', '#', '#'}, []int{1, 1, 3}, 1},
		{[]rune{'.', '?', '?', '.', '.', '?', '?', '.', '.', '.', '?', '#', '#', '.'}, []int{1, 1, 3}, 4},
		{[]rune{'?', '#', '#', '#', '?', '?', '?', '?', '?', '?', '?', '?'}, []int{3, 2, 1}, 10},
	}

	for i, tt := range tests {

		result := solver(tt.springs, tt.conditions)
		if result != tt.expected {
			t.Errorf("Expected: %d, got: %d Test: %d",
				tt.expected, result, i)
		}
	}

}

// func Test_partTwo(t *testing.T) {

// 	tests := []struct {
// 		filepath string
// 		expected uint64
// 	}{
// 		{"example.txt", 1030},
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
