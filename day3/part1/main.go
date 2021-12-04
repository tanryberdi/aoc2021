package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var numofBits [20]int
var lenOfLine, numOfLines int

func stringToByte(st string) {
	lenOfLine = len(st)

	for i, rune := range st {
		if rune == 49 { // if i-th char is 1
			numofBits[i]++
		}
	}
}

func calculateGamma(a [20]int) (string, string) {
	gamma := ""
	epsilon := ""
	for i := 0; i < lenOfLine; i++ {
		if a[i] > numOfLines/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	return gamma, epsilon
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numOfLines++
		stringToByte(scanner.Text())
		//os.Exit(1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var product int64
	product = 1
	gamma, epsilon := calculateGamma(numofBits)
	if g, err := strconv.ParseInt(gamma, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		product *= g
	}

	if e, err := strconv.ParseInt(epsilon, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		product *= e
	}

	fmt.Println(product)

}
