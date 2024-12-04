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

	var inputGrid [][]string
	// Reading from a file using scanner.
	scanner := bufio.NewScanner(file)
	var line string
	var cnt = 1
	for scanner.Scan() {
		line = scanner.Text()
		lineInput = append(lineInput, line)
		cnt++
	}
	fmt.Println(lineInput)
	for x := 0; x < len(lineInput); x++ {
		inputGrid = append(inputGrid, strings.Split(lineInput[x], ""))
	}
	return inputGrid
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 4")
	var input = ReadFile("./test.txt")
	fmt.Println(input)

}
