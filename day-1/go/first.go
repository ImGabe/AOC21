package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
)

func First(scanner *bufio.Scanner) {
	previous, n := math.MaxInt32, 0

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		if previous < number {
			n += 1
		}

		previous = number
	}

	fmt.Println(n)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
