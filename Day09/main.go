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

type coordinates struct {
	x int
	y int
}

func isInTheGrid(xLength int, yLength int, position coordinates) bool {
	return position.x >= 0 && position.x < xLength && position.y >= 0 && position.y < yLength
}

func findIndex(s []string, v string) int {
	for i, vs := range s {
		if vs == v {
			return i
		}
	}

	return -1
}

func findLastIndex(s []string, v string) int {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}

	return -1
}

func findLastNonPeriodIndex(s []string, v string) int {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != v {
			return i
		}
	}

	return -1
}

type dataObject struct {
	key         string
	firstIndex  int
	secondIndex int
	length      int
}

func generateEmpties(blocks []string) []dataObject {
	var empties []dataObject
	var previousValue = "0"
	var tempFirst = 0
	var tempSecond = 1

	for i := 0; i < len(blocks); i++ {
		if blocks[i] == "." {

			if previousValue != "." {
				tempFirst = i
			}

			previousValue = blocks[i]
		} else if blocks[i] != "." {
			if previousValue == "." {
				tempSecond = i - 1
				empties = append(empties, dataObject{key: ".", firstIndex: tempFirst, secondIndex: tempSecond, length: (tempSecond - tempFirst) + 1})
				previousValue = blocks[i]
			}
		}
	}

	return empties

}

func Part1(input string) int {

	var inputSlice []string
	for _, slice := range strings.Split(input, "") {
		inputSlice = append(inputSlice, slice)
	}

	var blocks []string

	for i := 0; i < len(inputSlice); i++ {
		mod := math.Mod(float64(i), 2)
		if mod == 0 {
			for j := 0; j < convertStringToInt(inputSlice[i]); j++ {
				blocks = append(blocks, convertIntToString(i/2))
			}
		}
		if mod == 1 {
			for j := 0; j < convertStringToInt(inputSlice[i]); j++ {
				blocks = append(blocks, ".")
			}
		}
	}
	for i := 0; i < len(blocks); i++ {

		indexFirst := findIndex(blocks, ".")
		indexLast := findLastNonPeriodIndex(blocks, ".")
		if indexFirst < indexLast {
			blocks[indexFirst] = blocks[indexLast]
			blocks[indexLast] = "."
		}

	}

	total := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] != "." {
			total += convertStringToInt(blocks[i]) * i
		}
	}

	return total
}

func Part2(input string) int {

	var inputSlice []string
	for _, slice := range strings.Split(input, "") {
		inputSlice = append(inputSlice, slice)
	}

	var blocks []string

	for i := 0; i < len(inputSlice); i++ {
		mod := math.Mod(float64(i), 2)
		if mod == 0 {
			for j := 0; j < convertStringToInt(inputSlice[i]); j++ {
				blocks = append(blocks, convertIntToString(i/2))
			}
		}
		if mod == 1 {
			for j := 0; j < convertStringToInt(inputSlice[i]); j++ {
				blocks = append(blocks, ".")
			}
		}
	}

	fmt.Println("Part 2 = blocks ", blocks)

	// Iterate downwards through all values
	fmt.Println("blocks length ", len(blocks))

	var files []dataObject
	for i := 9999; i > -1; i-- {
		firstIndexOfCurrentNumber := findIndex(blocks, convertIntToString(i))
		secondIndexOfCurrentNumber := findLastIndex(blocks, convertIntToString(i))
		files = append(files, dataObject{key: convertIntToString(i), firstIndex: firstIndexOfCurrentNumber, secondIndex: secondIndexOfCurrentNumber, length: (secondIndexOfCurrentNumber - firstIndexOfCurrentNumber) + 1})

	}
	fmt.Println("files ", files)

	empties := generateEmpties(blocks)
	fmt.Println("empties ", empties)

	for i := 0; i < len(files); i++ {
	out:
		for e := 0; e < len(empties); e++ {
			if i == 8 {
				fmt.Println()
			}

			if empties[e].length >= files[i].length {
				if empties[e].firstIndex < files[i].firstIndex {
					//copy files to empty  starting at first index of empties
					for f := 0; f < files[i].length; f++ {
						blocks[empties[e].firstIndex+f] = files[i].key
						//blocks[files[i].firstIndex+f] = "."

					}
					//copy "." to blocks  starting at first index of files
					for g := 0; g < files[i].length; g++ {

						blocks[files[i].firstIndex+g] = "."

					}
				}

				empties = generateEmpties(blocks)
				break out
			}

		}

	}
	fmt.Println("Part 2 ", blocks)
	total := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] != "." {
			total += convertStringToInt(blocks[i]) * i
		}
	}

	return total

}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 9")
	var input = ReadFile("./input.txt")
	fmt.Println("Part 1 = ", Part1(input[0]))
	fmt.Println("Part 2 = ", Part2(input[0]))
}
