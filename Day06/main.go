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
	} else if inputMap[currentPosition[0]-1][currentPosition[1]] == "#" || inputMap[currentPosition[0]-1][currentPosition[1]] == "o" {
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
	} else if inputMap[currentPosition[0]][currentPosition[1]+1] == "#" || inputMap[currentPosition[0]][currentPosition[1]+1] == "o" {
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
	} else if inputMap[currentPosition[0]+1][currentPosition[1]] == "#" || inputMap[currentPosition[0]+1][currentPosition[1]] == "o" {
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
	} else if inputMap[currentPosition[0]][currentPosition[1]-1] == "#" || inputMap[currentPosition[0]][currentPosition[1]-1] == "o" {
		newPosition = append(newPosition, currentPosition[0])
		newPosition = append(newPosition, currentPosition[1])
		changeDir = true
	}
	return newPosition, changeDir
}

func part1(inputMap [][]string) [][]string {
	var startPosition []int
	var currentPosition []int
	var currentDirection = "Up"
	for x := 0; x < len(inputMap); x++ {
		for y := 0; y < len(inputMap[x]); y++ {
			if inputMap[x][y] == "^" {
				startPosition = append(currentPosition, x, y)
				currentPosition = append(currentPosition, x, y)

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

	inputMap[startPosition[0]][startPosition[1]] = "^"
	// return input map with all visited locations on
	return inputMap
}

// PART 2
func part2(inputMap [][]string, updatedInputMap [][]string) {
	// Create array of unique visited locations
	var uniqueVisitedLocations [][]int
	for x := 0; x < len(updatedInputMap); x++ {
		for y := 0; y < len(updatedInputMap[x]); y++ {
			if updatedInputMap[x][y] == "x" {
				uniqueVisitedLocations = append(uniqueVisitedLocations, []int{x, y})
			}
		}
	}
	fmt.Println("uniqueVisitedLocations = ", uniqueVisitedLocations)
	var loopCount = 0
	var InfiniteLoopsFound = 0
	// Loop Over All visited locations and try adding a 0 to each and see if the previous code gets stuck in a loop
	for i := 0; i < len(uniqueVisitedLocations); i++ {
		loopCount = 0
		var workingInputMap [][]string
		workingInputMap = ReadFile("./input.txt")
		workingInputMap[uniqueVisitedLocations[i][0]][uniqueVisitedLocations[i][1]] = "o"
		//workingInputMap[8][3] = "o"
		var currentPosition []int
		var currentDirection = "Up"
		for x := 0; x < len(workingInputMap); x++ {
			for y := 0; y < len(workingInputMap[x]); y++ {
				if workingInputMap[x][y] == "^" {
					currentPosition = append(currentPosition, x, y)

					workingInputMap[x][y] = "x"
				}
			}
		}

		var stillOnMap = true
		for stillOnMap {

			if currentDirection == "Up" {
				if currentPosition[0] == 0 {
					break
				}
				currentPos, changeDir := checkUp(workingInputMap, currentPosition)
				if changeDir {
					currentDirection = changeDirection(currentDirection)
				}
				currentPosition = currentPos

				workingInputMap[currentPosition[0]][currentPosition[1]] = "x"

			} else if currentDirection == "Down" {
				if currentPosition[0] == len(workingInputMap)-1 {
					break
				}

				currentPos, changeDir := checkDown(workingInputMap, currentPosition)
				if changeDir {
					currentDirection = changeDirection(currentDirection)
				}
				currentPosition = currentPos

				workingInputMap[currentPosition[0]][currentPosition[1]] = "x"

			} else if currentDirection == "Left" {
				if currentPosition[1] == 0 {
					break
				}
				currentPos, changeDir := checkLeft(workingInputMap, currentPosition)
				if changeDir {
					currentDirection = changeDirection(currentDirection)
				}
				currentPosition = currentPos

				workingInputMap[currentPosition[0]][currentPosition[1]] = "x"
			} else if currentDirection == "Right" {

				if currentPosition[1] == len(workingInputMap[0])-1 {
					break
				}
				currentPos, changeDir := checkRight(workingInputMap, currentPosition)
				if changeDir {
					currentDirection = changeDirection(currentDirection)
				}
				currentPosition = currentPos

				workingInputMap[currentPosition[0]][currentPosition[1]] = "x"
			}
			loopCount++

			if loopCount > 10000 {
				fmt.Println("infiniteLoopFound", loopCount)
				InfiniteLoopsFound++
				break
			}
		}

	}
	fmt.Println("InfiniteLoopsFound ", InfiniteLoopsFound)
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 6")
	var inputMap = ReadFile("./input.txt")
	var updatedInputMap = part1(inputMap)
	var part2InputMap = ReadFile("./input.txt")
	part2(part2InputMap, updatedInputMap)

}
