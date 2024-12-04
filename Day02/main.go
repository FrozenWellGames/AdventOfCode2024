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

func ReadFile(filepath string) []string {
	file, err := os.Open(filepath)

	check(err)
	defer file.Close()

	var returnInput []string
	scanner := bufio.NewScanner(file)
	var line string
	var cnt = 1
	for scanner.Scan() {
		line = scanner.Text()
		returnInput = append(returnInput, line)
		cnt++
	}
	return returnInput
}

func convertStringToInt(s string) int {
	returnInt, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}
	return returnInt
}

func CreateData(input []string) [][]int {
	var returnArray [][]int
	for i := 0; i < len(input); i++ {
		var stringArray = strings.Fields(input[i])
		var intArray []int
		for i := 0; i < len(stringArray); i++ {
			intArray = append(intArray, convertStringToInt(stringArray[i]))
		}
		returnArray = append(returnArray, intArray)
	}

	return returnArray
}

func isSafeLine(report []int) bool {
	var returnBool = true
	var increasing = report[0] < report[1]

	if increasing {
		for i := 1; i < len(report); i++ {
			if report[i-1] > report[i] {
				return false
			} else if int(math.Abs(float64(report[i-1])-float64(report[i]))) == 0 {
				return false
			} else if int(math.Abs(float64(report[i-1])-float64(report[i]))) > 3 {
				return false
			}
			returnBool = true
		}
	} else {
		for i := 1; i < len(report); i++ {
			if report[i-1] < report[i] {
				return false
			} else if int(math.Abs(float64(report[i-1])-float64(report[i]))) == 0 {
				return false
			} else if int(math.Abs(float64(report[i-1])-float64(report[i]))) > 3 {
				return false
			}
			returnBool = true
		}
	}
	return returnBool
}

func createDifferentArrays(report []int) [][]int {
	var returnArray [][]int
	for i := 0; i < len(report); i++ {
		var tempArray []int
		tempArray = append(tempArray, report...)
		tempArray = append(tempArray[:i], tempArray[i+1:]...)
		returnArray = append(returnArray, tempArray)
	}
	return returnArray
}

// Part 1
func part1(data [][]int) int {
	var partOneTrueValues = 0
	for i := 0; i < len(data); i++ {
		if isSafeLine(data[i]) {
			partOneTrueValues += 1
		}
	}
	return partOneTrueValues
}

// Part 2

func part2(data [][]int) int {
	var partTwoTrueValues = 0
	for i := 0; i < len(data); i++ {
		if isSafeLine(data[i]) {
			partTwoTrueValues += 1
		} else {
			var array = createDifferentArrays(data[i])
		out:
			for x := 0; x < len(array); x++ {
				if isSafeLine(array[x]) {
					partTwoTrueValues += 1
					break out
				}
			}
		}
	}
	return partTwoTrueValues
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 1")
	var input = ReadFile("./input.txt")
	data := CreateData(input)
	fmt.Println("data = ", data)
	fmt.Println("Part 1 = ", part1(data))
	fmt.Println("Part 2 = ", part2(data))
}
