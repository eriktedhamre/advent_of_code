package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type mapping struct {
	dstRangeStart uint64
	srcRangeStart uint64
	length        uint64
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
	fmt.Print(partTwo(file))
}

func partOne(file *os.File) uint64 {
	var row string
	var seeds []uint64
	var seedToSoil []mapping
	var soilToFert []mapping
	var fertToWater []mapping
	var waterToLight []mapping
	var lightToTemp []mapping
	var tempToHumid []mapping
	var humidityToLoc []mapping

	var lowest uint64 = math.MaxInt64
	var tmp uint64

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	row = scanner.Text()
	for _, v := range strings.Split(row, " ") {
		value, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		seeds = append(seeds, uint64(value))
	}

	scanner.Scan() // Empty Line
	scanner.Scan() // Seed To Soil
	seedToSoil = readMapping(scanner)
	scanner.Scan() // Soil to Fert
	soilToFert = readMapping(scanner)
	scanner.Scan() // Fert-to-Water
	fertToWater = readMapping(scanner)
	scanner.Scan() // Water to Light
	waterToLight = readMapping(scanner)
	scanner.Scan() // Light to Temp
	lightToTemp = readMapping(scanner)
	scanner.Scan() // Temp to humid
	tempToHumid = readMapping(scanner)
	scanner.Scan() // humid to loc
	humidityToLoc = readMapping(scanner)

	for _, seed := range seeds {
		tmp = useMapping(seed, seedToSoil)
		tmp = useMapping(tmp, soilToFert)
		tmp = useMapping(tmp, fertToWater)
		tmp = useMapping(tmp, waterToLight)
		tmp = useMapping(tmp, lightToTemp)
		tmp = useMapping(tmp, tempToHumid)
		tmp = useMapping(tmp, humidityToLoc)
		lowest = min(lowest, tmp)
	}

	return lowest
}

func partTwo(file *os.File) uint64 {
	var row string
	var seeds []uint64
	var seedToSoil []mapping
	var soilToFert []mapping
	var fertToWater []mapping
	var waterToLight []mapping
	var lightToTemp []mapping
	var tempToHumid []mapping
	var humidityToLoc []mapping

	var lowest uint64 = math.MaxInt64
	var tmp uint64

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	row = scanner.Text()
	for _, v := range strings.Split(row, " ") {
		value, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		seeds = append(seeds, uint64(value))
	}

	scanner.Scan() // Empty Line
	scanner.Scan() // Seed To Soil
	seedToSoil = readMapping(scanner)
	scanner.Scan() // Soil to Fert
	soilToFert = readMapping(scanner)
	scanner.Scan() // Fert-to-Water
	fertToWater = readMapping(scanner)
	scanner.Scan() // Water to Light
	waterToLight = readMapping(scanner)
	scanner.Scan() // Light to Temp
	lightToTemp = readMapping(scanner)
	scanner.Scan() // Temp to humid
	tempToHumid = readMapping(scanner)
	scanner.Scan() // humid to loc
	humidityToLoc = readMapping(scanner)

	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			tmp = useMapping(j, seedToSoil)
			tmp = useMapping(tmp, soilToFert)
			tmp = useMapping(tmp, fertToWater)
			tmp = useMapping(tmp, waterToLight)
			tmp = useMapping(tmp, lightToTemp)
			tmp = useMapping(tmp, tempToHumid)
			tmp = useMapping(tmp, humidityToLoc)
			lowest = min(lowest, tmp)
		}
	}

	return lowest
}

func readMapping(scanner *bufio.Scanner) []mapping {
	var row string
	var numbers []string
	var result []mapping
	for scanner.Scan() {
		row = scanner.Text()

		if len(row) == 0 {
			break
		}

		numbers = strings.Split(row, " ")
		dstRangeStart, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Printf("numbers[%d] wrong conversion", 0)
		}

		srcRangeStart, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Printf("numbers[%d] wrong conversion", 0)
		}
		length, err := strconv.Atoi(numbers[2])
		if err != nil {
			fmt.Printf("numbers[%d] wrong conversion", 0)
		}

		result = append(result, mapping{dstRangeStart: uint64(dstRangeStart),
			srcRangeStart: uint64(srcRangeStart),
			length:        uint64(length)})
	}
	return result
}

func printMapping(m []mapping) {
	for _, v := range m {
		fmt.Printf("%d, %d, %d\n",
			v.dstRangeStart,
			v.srcRangeStart,
			v.length)
	}
}

func useMapping(input uint64, m []mapping) uint64 {
	for _, mapp := range m {
		if (mapp.srcRangeStart <= input) && (input < (mapp.srcRangeStart + mapp.length)) {
			return mapp.dstRangeStart + (input - mapp.srcRangeStart)
		}
	}
	return input
}

func contains[T comparable](slice []T, item T) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
