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
		{"example.txt", 8},
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

func Test_getGameId(t *testing.T) {
	tests := []struct {
		input         string
		expected_id   int
		expected_rest string
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			1, " 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
	}

	for _, tt := range tests {

		id, rest := getGameId(tt.input)
		if id != tt.expected_id {
			t.Errorf("Expected id: %d, got=%d", tt.expected_id, id)
		}
		if rest != tt.expected_rest {
			t.Errorf("Expected rest: %s \ngot: %s", tt.expected_rest, rest)
		}
	}
}
