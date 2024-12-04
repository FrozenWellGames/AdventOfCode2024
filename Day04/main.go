package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) [][]string {
	file, err := os.Open(filepath)

	check(err)
	defer file.Close()

	var lineInput []string

	var inputGrid [][]string
	// Reading from a file using scanner.
	scanner := bufio.NewScanner(file)
	var line string
	var cnt = 1
	for scanner.Scan() {
		line = scanner.Text()
		lineInput = append(lineInput, line)
		cnt++
	}
	for x := 0; x < len(lineInput); x++ {
		inputGrid = append(inputGrid, strings.Split(lineInput[x], ""))
	}
	return inputGrid
}

// Part 1 Functions
func checkRight(x int, y int, inputGrid [][]string) bool {
	var validXMAS = false
	if len(inputGrid[y])-y > 3 {
		if inputGrid[x][y+1] == "M" && inputGrid[x][y+2] == "A" && inputGrid[x][y+3] == "S" {
			validXMAS = true
		}
	}
	return validXMAS
}

func checkLeft(x int, y int, inputGrid [][]string) bool {
	var validXMAS = false
	if y > 2 {
		if inputGrid[x][y-1] == "M" && inputGrid[x][y-2] == "A" && inputGrid[x][y-3] == "S" {
			validXMAS = true
		}
	}
	return validXMAS
}

func checkUp(x int, y int, inputGrid [][]string) bool {
	var validXMAS = false
	if x > 2 {
		if inputGrid[x-1][y] == "M" && inputGrid[x-2][y] == "A" && inputGrid[x-3][y] == "S" {
			validXMAS = true
		}
	}
	return validXMAS
}

func checkDown(x int, y int, inputGrid [][]string) bool {
	var validXMAS = false
	if len(inputGrid[x])-x > 3 {
		if inputGrid[x+1][y] == "M" && inputGrid[x+2][y] == "A" && inputGrid[x+3][y] == "S" {
			validXMAS = true
		}
	}
	return validXMAS
}

func upRight(x int, y int, inputGrid [][]string) bool {
	var validXMAS = false
	if x > 2 && len(inputGrid[y])-y > 3 {
		if inputGrid[x-1][y+1] == "M" && inputGrid[x-2][y+2] == "A" && inputGrid[x-3][y+3] == "S" {
			validXMAS = true
		}
	}
	return validXMAS
}

func downRight(x int, y int, inputGrid [][]string) bool {
	var validXMAS = false
	if len(inputGrid[x])-x > 3 && len(inputGrid[y])-y > 3 {
		if inputGrid[x+1][y+1] == "M" && inputGrid[x+2][y+2] == "A" && inputGrid[x+3][y+3] == "S" {
			validXMAS = true
		}
	}
	return validXMAS
}

func downLeft(x int, y int, inputGrid [][]string) bool {
	var validXMAS = false
	if len(inputGrid[x])-x > 3 && y > 2 {
		if inputGrid[x+1][y-1] == "M" && inputGrid[x+2][y-2] == "A" && inputGrid[x+3][y-3] == "S" {
			validXMAS = true
		}
	}
	return validXMAS
}

func upLeft(x int, y int, inputGrid [][]string) bool {
	var validXMAS = false
	if x > 2 && y > 2 {
		if inputGrid[x-1][y-1] == "M" && inputGrid[x-2][y-2] == "A" && inputGrid[x-3][y-3] == "S" {
			validXMAS = true
		}
	}
	return validXMAS
}

func checkAllDirectionsX(x int, y int, inputGrid [][]string) int {
	var count = 0
	if checkRight(x, y, inputGrid) {
		count++
	}
	if checkLeft(x, y, inputGrid) {
		count++
	}
	if checkUp(x, y, inputGrid) {
		count++
	}
	if checkDown(x, y, inputGrid) {
		count++
	}
	if upRight(x, y, inputGrid) {
		count++
	}
	if downRight(x, y, inputGrid) {
		count++
	}
	if downLeft(x, y, inputGrid) {
		count++
	}
	if upLeft(x, y, inputGrid) {
		count++
	}
	return count
}

func checkXs(inputGrid [][]string) int {
	var count = 0
	for x := 0; x < len(inputGrid); x++ {
		for y := 0; y < len(inputGrid[x]); y++ {
			if inputGrid[x][y] == "X" {
				count += checkAllDirectionsX(x, y, inputGrid)
			}
		}
	}
	return count
}

// Part 2 Functions
func upDiagonal(x int, y int, inputGrid [][]string) bool {
	var validMAS = false

	if (inputGrid[x-1][y+1] == "S" && inputGrid[x+1][y-1] == "M") || (inputGrid[x-1][y+1] == "M" && inputGrid[x+1][y-1] == "S") {
		validMAS = true
	}

	return validMAS
}
func downDiagonal(x int, y int, inputGrid [][]string) bool {
	var validMAS = false

	if (inputGrid[x-1][y-1] == "S" && inputGrid[x+1][y+1] == "M") || (inputGrid[x-1][y-1] == "M" && inputGrid[x+1][y+1] == "S") {
		validMAS = true
	}

	return validMAS
}

func checkAllDirectionsA(x int, y int, inputGrid [][]string) int {
	var count = 0

	if x > 0 && len(inputGrid[y])-y > 1 && len(inputGrid[x])-x > 1 && y > 0 {

		if upDiagonal(x, y, inputGrid) && downDiagonal(x, y, inputGrid) {
			count++
		}
	}

	return count
}

func checkAs(inputGrid [][]string) int {
	var count = 0
	for x := 0; x < len(inputGrid); x++ {
		for y := 0; y < len(inputGrid[x]); y++ {
			if inputGrid[x][y] == "A" {
				count += checkAllDirectionsA(x, y, inputGrid)
			}
		}
	}
	return count
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 4")
	var input = ReadFile("./input.txt")
	fmt.Println("Part 1 = ", checkXs(input))
	fmt.Println("Part 2 = ", checkAs(input))

}
