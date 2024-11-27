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
		{"example.txt", 136},
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

func Test_moveRocksinRowWest(t *testing.T) {
	tests := []struct {
		input    [][]rune
		expected [][]rune
	}{
		{
			[][]rune{
				{'O', 'O', 'O', 'O', '.', '#', '.', 'O', '.', '.'},
				{'O', 'O', '.', '.', '#', '.', '.', '.', '.', '#'},
				{'O', 'O', '.', '.', 'O', '#', '#', '.', '.', 'O'},
				{'O', '.', '.', '#', '.', 'O', 'O', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
				{'.', '.', '#', '.', '.', '.', '.', '#', '.', '#'},
				{'.', '.', 'O', '.', '.', '#', '.', 'O', '.', 'O'},
				{'.', '.', 'O', '.', '.', '.', '.', '.', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
			},
			[][]rune{
				{'O', 'O', 'O', 'O', '.', '#', 'O', '.', '.', '.'},
				{'O', 'O', '.', '.', '#', '.', '.', '.', '.', '#'},
				{'O', 'O', 'O', '.', '.', '#', '#', 'O', '.', '.'},
				{'O', '.', '.', '#', 'O', 'O', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
				{'.', '.', '#', '.', '.', '.', '.', '#', '.', '#'},
				{'O', '.', '.', '.', '.', '#', 'O', 'O', '.', '.'},
				{'O', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
			}},
	}

	for _, tt := range tests {
		for i := 0; i < len(tt.input); i++ {
			moveRocksinRowWest(tt.input, i)
		}

		for i, row := range tt.input {
			for j, r := range row {
				if r != tt.expected[i][j] {
					t.Fatalf("Element mistmach for i: %d, j: %d", i, j)
				}
			}
		}
	}
}

func Test_moveRocksinColumnSouth(t *testing.T) {
	tests := []struct {
		input    [][]rune
		expected [][]rune
	}{
		{
			[][]rune{
				{'O', 'O', 'O', 'O', '.', '#', '.', 'O', '.', '.'},
				{'O', 'O', '.', '.', '#', '.', '.', '.', '.', '#'},
				{'O', 'O', '.', '.', 'O', '#', '#', '.', '.', 'O'},
				{'O', '.', '.', '#', '.', 'O', 'O', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
				{'.', '.', '#', '.', '.', '.', '.', '#', '.', '#'},
				{'.', '.', 'O', '.', '.', '#', '.', 'O', '.', 'O'},
				{'.', '.', 'O', '.', '.', '.', '.', '.', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
			},
			[][]rune{
				{'.', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '#', '.', '.', '.', '.', '#'},
				{'.', '.', '.', 'O', '.', '#', '#', '.', '.', '.'},
				{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
				{'O', '.', 'O', '.', '.', '.', '.', 'O', '#', 'O'},
				{'O', '.', '#', '.', '.', 'O', '.', '#', '.', '#'},
				{'O', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
				{'O', 'O', '.', '.', '.', '.', 'O', 'O', '.', '.'},
				{'#', 'O', 'O', '.', '.', '#', '#', '#', '.', '.'},
				{'#', 'O', 'O', '.', 'O', '#', '.', '.', '.', 'O'},
			}},
	}

	for _, tt := range tests {
		for i := 0; i < len(tt.input[0]); i++ {
			moveRocksInColumnSouth(tt.input, i)
		}

		for i, row := range tt.input {
			for j, r := range row {
				if r != tt.expected[i][j] {
					t.Fatalf("Element mistmach for i: %d, j: %d", i, j)
				}
			}
		}
	}
}

func Test_moveRocksinRowEast(t *testing.T) {
	tests := []struct {
		input    [][]rune
		expected [][]rune
	}{
		{
			[][]rune{
				{'O', 'O', 'O', 'O', '.', '#', '.', 'O', '.', '.'},
				{'O', 'O', '.', '.', '#', '.', '.', '.', '.', '#'},
				{'O', 'O', '.', '.', 'O', '#', '#', '.', '.', 'O'},
				{'O', '.', '.', '#', '.', 'O', 'O', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
				{'.', '.', '#', '.', '.', '.', '.', '#', '.', '#'},
				{'.', '.', 'O', '.', '.', '#', '.', 'O', '.', 'O'},
				{'.', '.', 'O', '.', '.', '.', '.', '.', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
			},
			[][]rune{
				{'.', 'O', 'O', 'O', 'O', '#', '.', '.', '.', 'O'},
				{'.', '.', 'O', 'O', '#', '.', '.', '.', '.', '#'},
				{'.', '.', 'O', 'O', 'O', '#', '#', '.', '.', 'O'},
				{'.', '.', 'O', '#', '.', '.', '.', '.', 'O', 'O'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
				{'.', '.', '#', '.', '.', '.', '.', '#', '.', '#'},
				{'.', '.', '.', '.', 'O', '#', '.', '.', 'O', 'O'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'O'},
				{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
				{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
			}},
	}

	for _, tt := range tests {
		for i := 0; i < len(tt.input); i++ {
			moveRocksinRowEast(tt.input, i)
		}

		for i, row := range tt.input {
			for j, r := range row {
				if r != tt.expected[i][j] {
					t.Fatalf("Element mistmach for i: %d, j: %d", i, j)
				}
			}
		}
	}
}

func Test_updateGrid(t *testing.T) {
	tests := []struct {
		key      string
		input    [][]rune
		expected [][]rune
	}{
		{
			"OOOO....#.O.#O.O",
			[][]rune{
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
			},
			[][]rune{
				{'O', 'O', 'O', 'O'},
				{'.', '.', '.', '.'},
				{'#', '.', 'O', '.'},
				{'#', 'O', '.', 'O'},
			}},
	}

	for _, tt := range tests {

		updateGrid(tt.input, tt.key)
		for i, row := range tt.input {
			for j, r := range row {
				if r != tt.expected[i][j] {
					t.Fatalf("Element mistmach for i: %d, j: %d", i, j)
				}
			}
		}
	}
}