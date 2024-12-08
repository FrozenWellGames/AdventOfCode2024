package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {
	fmt.Println("Advent Of Code 2024 - Day 9")
	var input = ReadFile("./input.txt")
	fmt.Println(input)

}
