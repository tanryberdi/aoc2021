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
	path := "input2.txt"
	firstLine, matrix := readFromFile(path)

	fmt.Println(walkMatrix(firstLine, matrix))
}

func walkMatrix(numbers []int, matrix map[int][][]int) int {
	var boardsWin []int
	walked := make(map[int][][]int)
	for i := 0; i < len(matrix); i++ {
		walked[i] = [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
	}

	numWins := 0
	for i := 0; i < len(matrix); i++ {
		boardsWin = append(boardsWin, 0)
	}

	for i := 0; i < len(numbers); i++ {
		for k := 0; k < len(matrix); k++ {
			for l := 0; l < 5; l++ {
				for m := 0; m < 5; m++ {

					if matrix[k][l][m] == numbers[i] {
						walked[k][l][m] = 1

						// check for vertical
						nowX := l
						bb := true
						for nowX+1 <= 5 {
							bb = bb && (walked[k][nowX][m] == 1)
							nowX++
						}

						nowX = l
						for nowX >= 0 {
							bb = bb && (walked[k][nowX][m] == 1)
							nowX--
						}

						if bb {
							sumUnmarked := 0
							for ii := 0; ii < 5; ii++ {
								for jj := 0; jj < 5; jj++ {
									if walked[k][ii][jj] == 0 {
										sumUnmarked += matrix[k][ii][jj]
									}
								}
							}

							if boardsWin[k] == 0 {
								boardsWin[k] = 1
								numWins++
								if numWins == len(matrix) {
									return numbers[i] * sumUnmarked
								}
							}

						}

						// check for horizontal
						nowY := m
						bb = true
						for nowY+1 <= 5 {
							bb = bb && (walked[k][l][nowY] == 1)
							nowY++
						}

						nowY = m
						for nowY >= 0 {
							bb = bb && (walked[k][l][nowY] == 1)
							nowY--
						}

						if bb {
							sumUnmarked := 0
							for ii := 0; ii < 5; ii++ {
								for jj := 0; jj < 5; jj++ {
									if walked[k][ii][jj] == 0 {
										sumUnmarked += matrix[k][ii][jj]
									}
								}
							}

							if boardsWin[k] == 0 {
								boardsWin[k] = 1
								numWins++
								if numWins == len(matrix) {
									return numbers[i] * sumUnmarked
								}
							}

						}
					}

				}
			}
		}
	}

	fmt.Println(walked)
	return 0
}

// Converting strings to 2D int array for each matrix
func matrixStTo2DArrayInt(matrix []string) [][]int {
	newMatrix := make([][]int, 0)

	for _, st := range matrix {
		numbers := make([]int, 0)
		nums := strings.Split(st, " ")
		for _, num := range nums {
			numb, _ := strconv.Atoi(num)
			numbers = append(numbers, numb)
		}
		newMatrix = append(newMatrix, [][]int{numbers}...)
	}

	return newMatrix
}

// Reading input file ...
func readFromFile(path string) (numbers []int, matrix map[int][][]int) {
	var firstLine string
	matrix = make(map[int][][]int)
	numOfMatrix := 0

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	var matrixSt []string
	for scanner.Scan() {
		st := scanner.Text()

		if line == 0 {
			firstLine = strings.TrimSpace(st)
			nums := strings.Split(firstLine, ",")
			for _, num := range nums {
				numb, _ := strconv.Atoi(num)
				numbers = append(numbers, numb)
			}
		} else {
			if st != "" {
				st = strings.ReplaceAll(st, "  ", " ")
				st = strings.TrimLeft(st, " ")
				st = strings.TrimRight(st, " ")
				//fmt.Println(st)
				matrixSt = append(matrixSt, st)
				if len(matrixSt) == 5 {
					//fmt.Println(matrixStTo2DArrayInt(matrixSt))
					matrix[numOfMatrix] = matrixStTo2DArrayInt(matrixSt)
					numOfMatrix++
					matrixSt = nil
				}
			}
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers, matrix
}
