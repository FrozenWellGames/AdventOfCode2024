package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func main() {
	fmt.Println("Advent Of Code 2024 - Day 3")
	var input = ReadFile("./test.txt")
	fmt.Println(input[0])

	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllString(input[0], -1)
	fmt.Println(matches)
}
