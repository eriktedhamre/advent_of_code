package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

// The sum of revealed cubes of one color in one set
// cannot exceed the total number of cubes of that color
func partOne(file *os.File) uint64 {
	var cum_sum uint64 = 0
	scanner := bufio.NewScanner(file)

	red_limit := 12
	green_limit := 13
	blue_limit := 14
	var red int
	var green int
	var blue int
	var tmp int
	var invalid bool = false

	for scanner.Scan() {
		invalid = false
		red = 0
		green = 0
		blue = 0
		tmp = 0
		line := scanner.Text()
		id, rest := getGameId(line)

		for _, v := range strings.Split(rest, " ") {
			number, err := strconv.Atoi(v)

			if err == nil {
				tmp = number
				continue
			}

			switch {
			case strings.HasPrefix(v, "red"):
				red += tmp

			case strings.HasPrefix(v, "green"):
				green += tmp

			case strings.HasPrefix(v, "blue"):
				blue += tmp
			}

			if strings.HasSuffix(v, ";") {
				if validGame(red, red_limit, green, green_limit, blue, blue_limit) {
					red = 0
					green = 0
					blue = 0
				} else {
					invalid = true
					break
				}
			}
			tmp = 0
		}

		if !invalid && validGame(red, red_limit, green, green_limit, blue, blue_limit) {
			cum_sum += uint64(id)
		}
	}

	return cum_sum

}

func validGame(red, red_limit, green, green_limit, blue, blue_limit int) bool {
	return (red <= red_limit) && (green <= green_limit) && (blue <= blue_limit)
}

func getGameId(s string) (int, string) {

	start := strings.IndexAny(s, " ")
	end := strings.IndexAny(s, ":")

	value, _ := strconv.Atoi(s[start+1 : end])

	return value, s[end+1:]
}
