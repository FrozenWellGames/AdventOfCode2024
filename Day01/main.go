package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func CreateData(input []string) ([]int, []int) {

	var returnArrayOne []int
	var returnArrayTwo []int
	for i := 0; i < len(input); i++ {

		var stringArray = strings.Fields(input[i])

		returnArrayOne = append(returnArrayOne, convertStringToInt(stringArray[0]))
		returnArrayTwo = append(returnArrayTwo, convertStringToInt(stringArray[1]))

	}

	sort.Ints(returnArrayOne[:])
	sort.Ints(returnArrayTwo[:])

	return returnArrayOne, returnArrayTwo
}

func part1(arrayOne []int, arrayTwo []int) int {
	var totalPartOne int
	for i := 0; i < len(arrayOne); i++ {
		totalPartOne += int(math.Abs(float64(arrayOne[i]) - float64(arrayTwo[i])))
	}
	return totalPartOne
}

func part2(arrayOne []int, arrayTwo []int) int {
	var totalPartTwo int
	for x := 0; x < len(arrayOne); x++ {
		var currentArrayOneValue = arrayOne[x]
		var countInArrayTwo = 0

		for y := 0; y < len(arrayTwo); y++ {
			if currentArrayOneValue == arrayTwo[y] {
				countInArrayTwo++
			}
		}
		totalPartTwo += currentArrayOneValue * countInArrayTwo
	}
	return totalPartTwo
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 1")
	var input = ReadFile("./input.txt")
	arrayOne, arrayTwo := CreateData(input)
	fmt.Println("Part 1 = ", part1(arrayOne, arrayTwo))
	fmt.Println("Part 2 = ", part2(arrayOne, arrayTwo))

}
