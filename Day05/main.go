package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) ([]string, []string) {
	file, err := os.Open(filepath)

	check(err)
	defer file.Close()

	var returnInput []string
	// Reading from a file using scanner.
	scanner := bufio.NewScanner(file)
	var line string
	var cnt = 1
	for scanner.Scan() {
		line = scanner.Text()
		returnInput = append(returnInput, line)
		cnt++
	}
	var emptyIndex int
	for i := 0; i < len(returnInput); i++ {
		if returnInput[i] == "" {
			emptyIndex = i
		}
	}
	rulesInput := returnInput[0:emptyIndex]
	updateInput := returnInput[emptyIndex+1 : len(returnInput)]
	return rulesInput, updateInput
}

func convertStringToInt(s string) int {
	returnInt, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}
	return returnInt
}

func CreateRulesData(rulesInput []string) [][]string {
	var returnArray [][]string
	for i := 0; i < len(rulesInput); i++ {
		stringArray := strings.Split(rulesInput[i], "|")

		returnArray = append(returnArray, stringArray)
	}
	return returnArray
}

func CreatePagesData(pagesInput []string) [][]string {
	var returnArray [][]string
	for i := 0; i < len(pagesInput); i++ {
		stringArray := strings.Split(pagesInput[i], ",")

		returnArray = append(returnArray, stringArray)
	}

	return returnArray
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func part1(rulesInput [][]string, pagesInput [][]string) ([][]string, [][]string) {
	var validLines [][]string
	var invalidLines [][]string
	for p := 0; p < len(pagesInput); p++ {
		var isValid = true
	out:
		for i := 0; i < len(pagesInput[p]); i++ {
			for r := 0; r < len(rulesInput); r++ {
				if rulesInput[r][0] == pagesInput[p][i] {
					indexOne := indexOf(rulesInput[r][0], pagesInput[p])
					indexTwo := indexOf(rulesInput[r][1], pagesInput[p])
					if indexTwo != -1 {
						if indexOne > indexTwo {
							isValid = false
							break out
						}
					}
				}
			}
		}
		if isValid {
			validLines = append(validLines, pagesInput[p])
		} else {
			invalidLines = append(invalidLines, pagesInput[p])
		}
	}
	return validLines, invalidLines
}

func getMiddleNumbers(pagesInput [][]string) int {
	var total = 0
	for i := 0; i < len(pagesInput); i++ {
		var middleIndex = math.Floor(float64(len(pagesInput[i]) / 2))
		middleIndexValue := pagesInput[i][int(middleIndex)]
		total += convertStringToInt(middleIndexValue)
	}
	return total
}

func reformatData(rulesInput [][]string, pagesInput [][]string) [][]string {
	var formattedData [][]string
	for p := 0; p < len(pagesInput); p++ {
		var workingRow []string
		workingRow = append(workingRow, pagesInput[p]...)
		var isNotCorrect = true
		for isNotCorrect {
			isNotCorrect = false
			for r := 0; r < len(rulesInput); r++ {
				indexOne := indexOf(rulesInput[r][0], workingRow)
				indexTwo := indexOf(rulesInput[r][1], workingRow)
				if indexOne != -1 && indexTwo != -1 && indexOne > indexTwo {
					var temp = workingRow[indexOne]
					workingRow[indexOne] = workingRow[indexTwo]
					workingRow[indexTwo] = temp
					isNotCorrect = true
				}
			}
		}
		formattedData = append(formattedData, workingRow)
	}

	return formattedData
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 5")
	var rulesInput, pagesInput = ReadFile("./input.txt")
	var rulesData = CreateRulesData(rulesInput)
	var pagesData = CreatePagesData(pagesInput)
	validLines, invalidLines := part1(rulesData, pagesData)
	fmt.Println("Part 1 = ", getMiddleNumbers(validLines))
	var reformattedData = reformatData(rulesData, invalidLines)
	fmt.Println("Part 2 = ", getMiddleNumbers(reformattedData))
}
