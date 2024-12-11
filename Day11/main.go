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
	// Reading from a file using scanner.
	scanner := bufio.NewScanner(file)
	var line string
	var cnt = 1
	for scanner.Scan() {
		line = scanner.Text()
		returnInput = append(returnInput, line)
		cnt++
	}

	var array = strings.Split(returnInput[0], " ")
	return array

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

func Part1(input []string) {
	for c := 0; c < 25; c++ {
		var workingArray = []string{}
		for i := 0; i < len(input); i++ {
			if input[i] == "0" {
				workingArray = append(workingArray, "1")
			} else if math.Mod(float64(len(input[i])), 2) == 0 {
				res1 := input[i][0 : len(input[i])/2]
				res2Int := convertStringToInt(input[i][len(input[i])/2:])
				workingArray = append(workingArray, res1)
				workingArray = append(workingArray, convertIntToString(res2Int))
			} else {
				workingArray = append(workingArray, (convertIntToString(convertStringToInt(input[i]) * 2024)))
			}
		}
		input = workingArray
	}
	fmt.Println("Part 1 = ", len(input))
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 11")
	var input = ReadFile("./test.txt")
	fmt.Println(input)
	Part1(input)

}
