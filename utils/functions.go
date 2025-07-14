package utils

import (
	"fmt"
	"math"

	"github.com/eriktedhamre/advent_of_code/types"
)

func StringToIntSlice(input string) ([]int, error) {
	var result []int
	for _, c := range input {
		if c < '0' || c > '9' {
			return nil, fmt.Errorf("Non digit character in string: %q", c)
		}
		result = append(result, int(c-'0'))
	}
	return result, nil
}

func InBound(row, col, rows, cols int) bool {
	return row > -1 && row < rows && col > -1 && col < cols
}

func MatrixRotate(grid []string) []string {
	tmp := make([][]rune, len(grid[0]))
	result := make([]string, 0)
	for row := range tmp {
		tmp[row] = make([]rune, len(grid))
	}

	for i, s := range grid {
		for j := range s {
			tmp[j][i] = rune(s[j])
		}
	}

	for i := range tmp {
		result = append(result, string(tmp[i]))
	}

	return result
}

func BoolToInt(b bool) uint64 {
	if b {
		return 1
	}
	return 0
}

func Contains(slice []rune, elem rune) bool {
	var result bool = false
DONE:
	for _, v := range slice {
		if v == elem {
			result = true
			break DONE
		}
	}
	return result
}

func FindFirst(slice []rune, elem rune) int64 {
	var result int64 = -1
DONE:
	for i := range slice {
		if slice[i] == elem {
			result = int64(i)
			break DONE
		}
	}
	return result
}

func PrefixSum(nums []uint64) []uint64 {
	prefixSums := make([]uint64, len(nums))
	if len(nums) == 0 {
		return prefixSums
	}

	prefixSums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSums[i] = prefixSums[i-1] + nums[i]
	}
	return prefixSums
}

func NaiveGCD(a uint64, b uint64) uint64 {
	if b == 0 {
		return a
	} else {
		return NaiveGCD(b, a%b)
	}
}

func NaiveLCM(a uint64, b uint64) uint64 {
	return (a * b) / NaiveGCD(a, b)
}

func AreaOfAPolygon(vertices []types.Coordinates) int64 {
	var res int64 = 0
	for i := range len(vertices) - 1 {
		res += int64(vertices[i].Row)*int64(vertices[i+1].Col) -
			int64(vertices[i+1].Row)*int64(vertices[i].Col)
	}
	res += int64(vertices[len(vertices)-1].Row)*int64(vertices[0].Col) -
		int64(vertices[0].Row)*int64(vertices[len(vertices)-1].Col)
	return int64(math.Abs(float64(res)) / 2)
}
