package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	first(scanner)

	file.Seek(0, io.SeekStart)

	scanner = bufio.NewScanner(file)
	second(scanner)
}
