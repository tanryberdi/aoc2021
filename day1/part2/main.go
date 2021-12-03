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

	scanner := bufio.NewScanner(file)
	a1, a2, a3, a4 := 0, 0, 0, 0
	count := 0
	for scanner.Scan() {
		a4, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum1 := a1 + a2 + a3
		sum2 := a2 + a3 + a4
		if sum2 > sum1 {
			count++
		}
		a1, a2, a3 = a2, a3, a4
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count - 3)

}
