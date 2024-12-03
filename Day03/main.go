package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	// Reading from a file using scanner.
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

func getRawMulData(input []string) []string {
	var allMatches []string
	for i := 0; i < len(input); i++ {
		regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := regex.FindAllString(input[i], -1)

		for i := 0; i < len(matches); i++ {
			allMatches = append(allMatches, matches[i])
		}
		//	for i := 0; i < len(allMatches); i++ {
		//		fmt.Println(allMatches[i])
		//	}
	}
	return allMatches
}

func getRawMulDataTwo(input []string) []string {
	var allMatches []string
	//var active = true
	for i := 0; i < len(input); i++ {
		regex := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
		matches := regex.FindAllString(input[i], -1)

		for i := 0; i < len(matches); i++ {
			allMatches = append(allMatches, matches[i])
		}
	}
	return allMatches
}

func getFirstAndSecondNumbersRawMulData(rawMulData []string) [][]int {
	var firstAndSecondNumbers [][]int

	for i := 0; i < len(rawMulData); i++ {
		split := strings.Split(rawMulData[i], `,`)
		re := regexp.MustCompile("[0-9]+")
		var numberOne = re.FindAllString(split[0], -1)
		var numberTwo = re.FindAllString(split[1], -1)

		//
		int1, err := strconv.Atoi(numberOne[0])
		if err != nil {
		}
		int2, err := strconv.Atoi(numberTwo[0])
		if err != nil {
		}

		bothNumber := []int{int1, int2}
		firstAndSecondNumbers = append(firstAndSecondNumbers, bothNumber)
	}

	return firstAndSecondNumbers
}

func getFirstAndSecondNumbersRawMulDataPartTwo(rawMulData []string) [][]int {
	var firstAndSecondNumbers [][]int
	var active = true
	fmt.Println(active)
	for i := 0; i < len(rawMulData); i++ {
		fmt.Println(rawMulData[i])
		if rawMulData[i] == "don't()" {
			fmt.Println("It's Don't")
			active = false
		} else if rawMulData[i] == "do()" {
			fmt.Println("It's Do")
			active = true
		} else if active == true {

			split := strings.Split(rawMulData[i], `,`)
			re := regexp.MustCompile("[0-9]+")
			var numberOne = re.FindAllString(split[0], -1)
			var numberTwo = re.FindAllString(split[1], -1)

			//
			int1, err := strconv.Atoi(numberOne[0])
			if err != nil {
			}
			int2, err := strconv.Atoi(numberTwo[0])
			if err != nil {
			}

			bothNumber := []int{int1, int2}
			firstAndSecondNumbers = append(firstAndSecondNumbers, bothNumber)
		}
	}

	return firstAndSecondNumbers
}

func calculateTotalPartOne(numbers [][]int) int {
	var returnTotal int = 0

	for i := 0; i < len(numbers); i++ {
		returnTotal += numbers[i][0] * numbers[i][1]
	}

	return returnTotal
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 3")
	fmt.Println("Part 1")
	input := ReadFile("./test.txt")
	var rawMulData = getRawMulData(input)
	fmt.Println(rawMulData)
	var numbers = getFirstAndSecondNumbersRawMulData(rawMulData)
	fmt.Println(numbers)
	var totalPartOne = calculateTotalPartOne(numbers)
	fmt.Println(totalPartOne)
	fmt.Println("Part 2")
	input2 := ReadFile("./input.txt")
	fmt.Println(input2)
	var rawMulDataPartTwo = getRawMulDataTwo(input2)
	fmt.Println("part 2 data ", rawMulDataPartTwo)
	var numbersPartTwo = getFirstAndSecondNumbersRawMulDataPartTwo(rawMulDataPartTwo)
	fmt.Println(numbersPartTwo)
	var totalPartTwo = calculateTotalPartOne(numbersPartTwo)
	fmt.Println(totalPartTwo)
}
