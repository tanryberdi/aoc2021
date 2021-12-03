package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	prev := 0
	depth := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		now, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if now > prev {
			depth++
		}

		prev = now
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(depth - 1)
}
