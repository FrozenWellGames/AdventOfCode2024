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
	var loopArray = []string{}
	loopArray = append(loopArray, input[0])
	for c := 0; c < 75; c++ {
		fmt.Println("Count = ", c)
		var workingArray = []string{}
		for i := 0; i < len(loopArray); i++ {
			if loopArray[i] == "0" {
				workingArray = append(workingArray, "1")
			} else if math.Mod(float64(len(loopArray[i])), 2) == 0 {
				res1 := loopArray[i][0 : len(loopArray[i])/2]
				res2Int := convertStringToInt(loopArray[i][len(loopArray[i])/2:])
				workingArray = append(workingArray, res1)
				workingArray = append(workingArray, convertIntToString(res2Int))
			} else {
				workingArray = append(workingArray, (convertIntToString(convertStringToInt(loopArray[i]) * 2024)))
			}
		}
		loopArray = workingArray
	}
	fmt.Println("Part 1 = ", len(loopArray))
	fmt.Println("Part 1 = ", len(loopArray))
}

func Part2(input []string) {
	stoneMap := make(map[int]int)

	for i := 0; i < len(input); i++ {
		value := convertStringToInt(input[i])
		stoneMap[value] = stoneMap[value] + 1
	}

	for i := 0; i < 75; i++ {
		//create map to store new stones
		newStoneMap := make(map[int]int)
		//loop over stone map
		for stone, count := range stoneMap {
			var newStones []int
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if digits := convertIntToString(stone); len(digits)%2 == 0 {
				a := convertStringToInt(digits[:len(digits)/2])
				b := convertStringToInt(digits[len(digits)/2:])
				newStones = append(newStones, a)
				newStones = append(newStones, b)
			} else {
				newStones = append(newStones, stone*2024)
			}
			for x := 0; x < len(newStones); x++ {
				newStoneMap[newStones[x]] = newStoneMap[newStones[x]] + count
			}
		}
		// update stoneMap for next loop
		stoneMap = newStoneMap
	}
	result := 0
	for _, count := range stoneMap {
		result += count
	}

	fmt.Println("Part 2 = ", result)
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 11")
	var input = ReadFile("./input.txt")
	fmt.Println(input)
	//Part1(input)
	Part2(input)
}
