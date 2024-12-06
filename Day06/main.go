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

	var inputMap [][]string
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
		inputMap = append(inputMap, strings.Split(lineInput[x], ""))
	}
	return inputMap
}

func changeDirection(currentDirection string) string {
	if currentDirection == "Up" {
		return "Right"
	}
	if currentDirection == "Down" {
		return "Left"
	}
	if currentDirection == "Left" {
		return "Up"
	}
	if currentDirection == "Right" {
		return "Down"
	}
	return "Up"
}

func checkUp(inputMap [][]string, currentPosition []int) ([]int, bool) {
	var newPosition []int
	var changeDir = false
	if inputMap[currentPosition[0]-1][currentPosition[1]] == "." || inputMap[currentPosition[0]-1][currentPosition[1]] == "x" {
		newPosition = append(newPosition, currentPosition[0]-1)
		newPosition = append(newPosition, currentPosition[1])
	} else if inputMap[currentPosition[0]-1][currentPosition[1]] == "#" {
		newPosition = append(newPosition, currentPosition[0])
		newPosition = append(newPosition, currentPosition[1])
		changeDir = true
	}
	return newPosition, changeDir
}

func checkRight(inputMap [][]string, currentPosition []int) ([]int, bool) {
	var newPosition []int
	var changeDir = false
	if inputMap[currentPosition[0]][currentPosition[1]+1] == "." || inputMap[currentPosition[0]][currentPosition[1]+1] == "x" {
		newPosition = append(newPosition, currentPosition[0])
		newPosition = append(newPosition, currentPosition[1]+1)
	} else if inputMap[currentPosition[0]][currentPosition[1]+1] == "#" {
		newPosition = append(newPosition, currentPosition[0])
		newPosition = append(newPosition, currentPosition[1])
		changeDir = true
	}
	return newPosition, changeDir
}

func checkDown(inputMap [][]string, currentPosition []int) ([]int, bool) {
	var newPosition []int
	var changeDir = false
	if inputMap[currentPosition[0]+1][currentPosition[1]] == "." || inputMap[currentPosition[0]+1][currentPosition[1]] == "x" {
		newPosition = append(newPosition, currentPosition[0]+1)
		newPosition = append(newPosition, currentPosition[1])
	} else if inputMap[currentPosition[0]+1][currentPosition[1]] == "#" {
		newPosition = append(newPosition, currentPosition[0])
		newPosition = append(newPosition, currentPosition[1])
		changeDir = true
	}
	return newPosition, changeDir
}

func checkLeft(inputMap [][]string, currentPosition []int) ([]int, bool) {
	var newPosition []int
	var changeDir = false
	if inputMap[currentPosition[0]][currentPosition[1]-1] == "." || inputMap[currentPosition[0]][currentPosition[1]-1] == "x" {
		newPosition = append(newPosition, currentPosition[0])
		newPosition = append(newPosition, currentPosition[1]-1)
	} else if inputMap[currentPosition[0]][currentPosition[1]-1] == "#" {
		newPosition = append(newPosition, currentPosition[0])
		newPosition = append(newPosition, currentPosition[1])
		changeDir = true
	}
	return newPosition, changeDir
}

func part1(inputMap [][]string) {
	var visitedPositions [][]int
	var currentPosition []int
	var currentDirection = "Up"
	for x := 0; x < len(inputMap); x++ {
		for y := 0; y < len(inputMap[x]); y++ {
			if inputMap[x][y] == "^" {
				currentPosition = append(currentPosition, x, y)
				visitedPositions = append(visitedPositions, currentPosition)
				inputMap[x][y] = "x"
			}
		}
	}
	var stillOnMap = true
	for stillOnMap {

		if currentDirection == "Up" {
			if currentPosition[0] == 0 {
				break
			}
			currentPos, changeDir := checkUp(inputMap, currentPosition)
			if changeDir {
				currentDirection = changeDirection(currentDirection)
			}
			currentPosition = currentPos
			visitedPositions = append(visitedPositions, currentPosition)
			inputMap[currentPosition[0]][currentPosition[1]] = "x"

		} else if currentDirection == "Down" {
			if currentPosition[0] == len(inputMap)-1 {
				break
			}

			currentPos, changeDir := checkDown(inputMap, currentPosition)
			if changeDir {
				currentDirection = changeDirection(currentDirection)
			}
			currentPosition = currentPos
			visitedPositions = append(visitedPositions, currentPosition)
			inputMap[currentPosition[0]][currentPosition[1]] = "x"

		} else if currentDirection == "Left" {
			if currentPosition[1] == 0 {
				break
			}
			currentPos, changeDir := checkLeft(inputMap, currentPosition)
			if changeDir {
				currentDirection = changeDirection(currentDirection)
			}
			currentPosition = currentPos
			visitedPositions = append(visitedPositions, currentPosition)
			inputMap[currentPosition[0]][currentPosition[1]] = "x"
		} else if currentDirection == "Right" {

			if currentPosition[1] == len(inputMap[0])-1 {
				break
			}
			currentPos, changeDir := checkRight(inputMap, currentPosition)
			if changeDir {
				currentDirection = changeDirection(currentDirection)
			}
			currentPosition = currentPos
			visitedPositions = append(visitedPositions, currentPosition)
			inputMap[currentPosition[0]][currentPosition[1]] = "x"
		}

	}
	var count = 0
	for x := 0; x < len(inputMap); x++ {
		for y := 0; y < len(inputMap[x]); y++ {
			if inputMap[x][y] == "x" {
				count++
			}
		}
	}
	fmt.Println("Part1 = ", count)

}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 6")
	var inputMap = ReadFile("./input.txt")
	part1(inputMap)

}
