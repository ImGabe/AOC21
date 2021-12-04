package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateOnesFrequency(bitSize int, readings []uint64) []uint64 {
	onesFrequency := make([]uint64, bitSize)

	for _, reading := range readings {
		for i := 0; i < bitSize; i++ {
			onesFrequency[i] += reading >> i & 1
		}
	}

	return onesFrequency
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readings := make([]uint64, 0)
	bitSize := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > bitSize {
			bitSize = len(line)
		}

		reading, err := strconv.ParseUint(line, 2, len(line))
		if err != nil {
			log.Fatal(err)
		}

		readings = append(readings, reading)
	}

	onesFrequency := calculateOnesFrequency(bitSize, readings)

	var gamma uint64

	for i := 0; i < bitSize; i++ {
		zeroesFrequency := uint64(len(readings)) - onesFrequency[i]
		if onesFrequency[i] > zeroesFrequency {
			gamma |= (1 << i)
		}
	}

	oxygenGeneratorRatings := make([]uint64, len(readings))
	copy(oxygenGeneratorRatings, readings)

	for i := 0; i < bitSize; i++ {
		onesFrequency = calculateOnesFrequency(bitSize, oxygenGeneratorRatings)
		zeroesFrequency := uint64(len(oxygenGeneratorRatings)) - onesFrequency[bitSize-i-1]

		var bit uint64
		if onesFrequency[bitSize-i-1] >= zeroesFrequency {
			bit = 1
		}

		filteredRatings := make([]uint64, 0)
		for _, it := range oxygenGeneratorRatings {
			if it>>(bitSize-i-1)&1 == bit {
				filteredRatings = append(filteredRatings, it)
			}
		}
		oxygenGeneratorRatings = filteredRatings

		if len(oxygenGeneratorRatings) == 1 {
			break
		}
	}

	CO2ScrubberRatings := make([]uint64, len(readings))
	copy(CO2ScrubberRatings, readings)

	for i := 0; i < bitSize; i++ {
		onesFrequency = calculateOnesFrequency(bitSize, CO2ScrubberRatings)
		zeroesFrequency := uint64(len(CO2ScrubberRatings)) - onesFrequency[bitSize-i-1]

		var bit uint64
		if onesFrequency[bitSize-i-1] < zeroesFrequency {
			bit = 1
		}

		filteredRatings := make([]uint64, 0)
		for _, it := range CO2ScrubberRatings {
			if it>>(bitSize-i-1)&1 == bit {
				filteredRatings = append(filteredRatings, it)
			}
		}
		CO2ScrubberRatings = filteredRatings

		if len(CO2ScrubberRatings) == 1 {
			break
		}
	}

	epsilon := gamma ^ ((1 << bitSize) - 1)

	fmt.Println("Part 1:", gamma*epsilon)
	fmt.Println("Part 2:", oxygenGeneratorRatings[0]*CO2ScrubberRatings[0])
}
