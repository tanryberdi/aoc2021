package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 0
	y := 0
	depth := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		st := strings.Split(scanner.Text(), " ")
		//fmt.Println(len(st), " ",st[0]," ",st[1])
		units, _ := strconv.Atoi(st[1])

		switch st[0] {
		case "forward":
			x += units
			depth += y * units
		case "down":
			y += units
		default:
			y -= units
		}

		//fmt.Println(x, " ", y, " ", depth)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(x * depth)
}
