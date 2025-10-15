package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {

	tests := []struct {
		filepath string
		expected uint64
	}{
		{"example_test.txt", 19114},
		{"input.txt", 480738},
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

func TestPartTwo(t *testing.T) {

	tests := []struct {
		filepath string
		expected uint64
	}{
		{"easy.txt", 10000},
		{"easy2.txt", 256000000000000},
		{"medium.txt", 4000000},
		{"example.txt", 167409079868000},
		{"input.txt", 131550418841958},
	}

	for _, tt := range tests {
		file, err := os.Open(tt.filepath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		result := partTwo(file)
		if result != tt.expected {
			t.Errorf("Expected: %d, got: %d, file: %s",
				tt.expected, result, tt.filepath)
		}
	}
}
