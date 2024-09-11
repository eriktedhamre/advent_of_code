package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_Part_Two(t *testing.T) {

	tests := []struct {
		filepath string
		expected uint64
	}{
		{"example.txt", 142},
		{"example_2.txt", 281},
		{"tc_1.txt", 98},
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
