package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func convertStringToInt(s string) int {
	returnInt, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}
	return returnInt
}

func convertIntToString(s int) string {
	returnString := strconv.Itoa(s)

	return returnString
}

func formatData(input []string) ([]int, [][]int) {
	answerArray := []int{}
	numbersArray := [][]int{}
	for i := 0; i < len(input); i++ {
		firstSplit := strings.Split(input[i], ":")
		answerArray = append(answerArray, convertStringToInt(firstSplit[0]))
		trimmedNumbers := strings.TrimSpace(firstSplit[1])
		secondSplit := strings.Split(trimmedNumbers, " ")
		numbers := []int{}
		for i := 0; i < len(secondSplit); i++ {
			numbers = append(numbers, convertStringToInt(secondSplit[i]))
		}
		numbersArray = append(numbersArray, numbers)
	}

	return answerArray, numbersArray
}

func calculationsPart1(previousNumbers []int, currentNumber int) []int {
	returnNumbers := []int{}

	for i := 0; i < len(previousNumbers); i++ {
		plusAnswer := previousNumbers[i] + currentNumber
		multiplyAnswer := previousNumbers[i] * currentNumber
		returnNumbers = append(returnNumbers, plusAnswer)
		returnNumbers = append(returnNumbers, multiplyAnswer)
	}

	return returnNumbers

}

func calculationsPart2(previousNumbers []int, currentNumber int) []int {
	returnNumbers := []int{}

	for i := 0; i < len(previousNumbers); i++ {
		plusAnswer := previousNumbers[i] + currentNumber
		multiplyAnswer := previousNumbers[i] * currentNumber
		combineAnswer := convertStringToInt(convertIntToString(previousNumbers[i]) + convertIntToString(currentNumber))
		returnNumbers = append(returnNumbers, plusAnswer)
		returnNumbers = append(returnNumbers, multiplyAnswer)
		returnNumbers = append(returnNumbers, combineAnswer)
	}

	return returnNumbers

}

func part1(answersArray []int, numbersArray [][]int) int {
	var TotalToReturn int = 0

	for i := 0; i < len(answersArray); i++ {
		workingNumbersArray := []int{numbersArray[i][0]}
		for x := 1; x < len(numbersArray[i]); x++ {
			var returnNumbers = []int{}
			returnNumbers = calculationsPart1(workingNumbersArray, numbersArray[i][x])
			workingNumbersArray = nil
			workingNumbersArray = append(workingNumbersArray, returnNumbers...)
		}
		if slices.Contains(workingNumbersArray, answersArray[i]) {
			TotalToReturn += answersArray[i]
		}
	}
	return TotalToReturn
}

func part2(answersArray []int, numbersArray [][]int) int {
	var TotalToReturn int = 0

	for i := 0; i < len(answersArray); i++ {
		workingNumbersArray := []int{numbersArray[i][0]}
		for x := 1; x < len(numbersArray[i]); x++ {
			var returnNumbers = []int{}
			returnNumbers = calculationsPart2(workingNumbersArray, numbersArray[i][x])
			workingNumbersArray = nil
			workingNumbersArray = append(workingNumbersArray, returnNumbers...)
		}
		if slices.Contains(workingNumbersArray, answersArray[i]) {
			TotalToReturn += answersArray[i]
		}
	}
	return TotalToReturn
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 7")
	var input = ReadFile("./input.txt")
	answersArray, numbersArray := formatData(input)
	fmt.Println("Part 1 = ", part1(answersArray, numbersArray))
	fmt.Println("Part 2 = ", part2(answersArray, numbersArray))

}
