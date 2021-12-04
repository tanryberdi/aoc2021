package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Read input from file;
func readFromFile(path string) []string {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// Eliminate for every digit of bitCriteria to handle Oxygen
// generator rating;
func chooseO2(step int, bitCr []string) []string {
	var bitCrNew []string
	n0 := 0
	n1 := 0
	for i := 0; i < len(bitCr); i++ {
		if bitCr[i][step] == 48 {
			n0++
		} else {
			n1++
		}
	}

	if n1 >= n0 {
		for i := 0; i < len(bitCr); i++ {
			if bitCr[i][step] == 49 {
				bitCrNew = append(bitCrNew, bitCr[i])
			}
		}
	} else {
		for i := 0; i < len(bitCr); i++ {
			if bitCr[i][step] == 48 {
				bitCrNew = append(bitCrNew, bitCr[i])
			}
		}
	}

	return bitCrNew
}

// Eliminate for every digit of bitCriteria to handle CO2
// scrubber rating;
func chooseCO2(step int, bitCr []string) []string {
	var bitCrNew []string
	n0 := 0
	n1 := 0
	for i := 0; i < len(bitCr); i++ {
		if bitCr[i][step] == 48 {
			n0++
		} else {
			n1++
		}
	}

	if n1 >= n0 {
		for i := 0; i < len(bitCr); i++ {
			if bitCr[i][step] == 48 {
				bitCrNew = append(bitCrNew, bitCr[i])
			}
		}
	} else {
		for i := 0; i < len(bitCr); i++ {
			if bitCr[i][step] == 49 {
				bitCrNew = append(bitCrNew, bitCr[i])
			}
		}
	}

	return bitCrNew
}

func binaryToDecimal(st string) int64 {

	g, err := strconv.ParseInt(st, 2, 64)
	if err != nil {
		return 0
	}
	return g

}

func main() {
	path := "input2.txt"
	bitCriteria := readFromFile(path)

	newList := bitCriteria
	for i := 0; i < len(bitCriteria[0]); i++ {
		newList = chooseO2(i, newList)
		//fmt.Println(i, "-->", newList)
	}
	oxygen := binaryToDecimal(newList[0])
	//fmt.Println(oxygen)

	newList = bitCriteria
	for i := 0; i < len(bitCriteria[0]); i++ {
		newList = chooseCO2(i, newList)
		//fmt.Println(i, "-->", newList)
		if len(newList) == 1 {
			break
		}
	}
	co2 := binaryToDecimal(newList[0])
	//fmt.Println(co2)

	fmt.Println("Result is ===>", oxygen*co2)
}
