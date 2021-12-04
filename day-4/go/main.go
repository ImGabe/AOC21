package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	bingoGridSize = 5
)

func isGridWon(grid [][]int) bool {
	// Line
	for _, ys := range grid {
		jackpot := true

		for _, y := range ys {
			if y != -1 {
				jackpot = false
				break
			}
		}

		if jackpot {
			return true
		}
	}

	// Column
	for i := 0; i < bingoGridSize; i++ {
		jackpot := true

		for _, ys := range grid {
			if ys[i] != -1 {
				jackpot = false
				break
			}
		}

		if jackpot {
			return true
		}
	}

	return false
}

func sumUnmarkedNumbers(grid [][]int) (sum int) {
	for _, line := range grid {
		for _, number := range line {
			if number != -1 {
				sum += number
			}
		}
	}

	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	numbersString := strings.Split(scanner.Text(), ",")
	numbers := make([]int, len(numbersString))
	for i, number := range numbersString {
		numbers[i], err = strconv.Atoi(number)
	}

	grids := make([][][]int, 0)

	for scanner.Scan() {
		// Grid
		grid := make([][]int, bingoGridSize)

		for i := 0; i < bingoGridSize; i++ {
			scanner.Scan()

			columns := strings.Fields(scanner.Text())

			if len(columns) != bingoGridSize {
				log.Fatal("Invalid column size")
			}

			columnNumbers := make([]int, bingoGridSize)

			for j := 0; j < bingoGridSize; j++ {
				columnNumbers[j], err = strconv.Atoi(columns[j])
				if err != nil {
					log.Fatal(err)
				}
			}

			grid[i] = columnNumbers
		}

		grids = append(grids, grid)
	}

	gridsWonScores := make([]int, 0)

	for _, number := range numbers {
		// Grids
		for _, grid := range grids {
			for _, gridLine := range grid {
				for idx, gridNumber := range gridLine {
					if gridNumber == number {
						gridLine[idx] = -1
					}
				}
			}
		}

		// Won?
		newGrids := make([][][]int, 0)
		for _, grid := range grids {
			if isGridWon(grid) {
				score := sumUnmarkedNumbers(grid)
				gridsWonScores = append(gridsWonScores, score*number)
			} else {
				newGrids = append(newGrids, grid)
			}
		}

		grids = newGrids
	}

	fmt.Println("Part 1:", gridsWonScores[0])
	fmt.Println("Part 2:", gridsWonScores[len(gridsWonScores)-1])
}
