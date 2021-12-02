package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

func Second(scanner *bufio.Scanner) {
	arr := make([]int, 3)
	var err error

	for i := 0; i < 3; i++ {
		if !scanner.Scan() {
			log.Fatal("aaadsasadasas")
		}

		arr[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}

	previous, n := arr[0]+arr[1]+arr[2], 0
	arr = arr[1:]

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		sum := arr[0] + arr[1] + number

		if previous < sum {
			n += 1
		}

		arr = append(arr, number)
		arr = arr[1:]

		previous = sum
	}

	fmt.Println(n)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
