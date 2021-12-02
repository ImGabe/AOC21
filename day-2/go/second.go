package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func second(scanner *bufio.Scanner) {
	hposition, depth, aim := 0, 0, 0

	for scanner.Scan() {
		split_text := strings.Split(scanner.Text(), " ")

		direction := split_text[0]
		number, err := strconv.Atoi(split_text[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "up":
			aim -= number
		case "down":
			aim += number
		case "forward":
			hposition += number
			depth += aim * number
		}
	}

	fmt.Println("Second:", hposition*depth)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
