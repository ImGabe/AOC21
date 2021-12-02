package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func first(scanner *bufio.Scanner) {
	hposition, depth := 0, 0

	for scanner.Scan() {
		split_text := strings.Split(scanner.Text(), " ")

		direction := split_text[0]
		number, err := strconv.Atoi(split_text[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "up":
			depth -= number
		case "down":
			depth += number
		case "forward":
			hposition += number
		}
	}

	fmt.Println("First:", hposition*depth)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
