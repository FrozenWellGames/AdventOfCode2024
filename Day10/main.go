package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func convertStringToInt(s string) int {
	returnInt, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}
	return returnInt
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) [][]int {
	file, err := os.Open(filepath)

	check(err)
	defer file.Close()

	var lineInput []string

	var inputMap [][]int
	// Reading from a file using scanner.
	scanner := bufio.NewScanner(file)
	var line string
	var cnt = 1
	for scanner.Scan() {
		line = scanner.Text()
		lineInput = append(lineInput, line)
		cnt++
	}
	for x := 0; x < len(lineInput); x++ {
		stringArray := strings.Split(lineInput[x], "")
		intArray := []int{}
		for i := 0; i < len(stringArray); i++ {
			intArray = append(intArray, convertStringToInt(stringArray[i]))
		}

		inputMap = append(inputMap, intArray)

	}
	return inputMap
}

type Cell struct {
	x, y int
}

func Part1(inputGrid [][]int) int {

	var score int
	var move func(int, int, int)
	visited := map[Cell]bool{}

	move = func(x int, y int, target int) {
		if x < 0 || x >= len(inputGrid) || y < 0 || y >= len(inputGrid[0]) || inputGrid[x][y] != target {
			return
		}
		if target == 9 && !visited[Cell{x, y}] {
			score++
			visited[Cell{x, y}] = true
			return
		}
		target++
		move(x-1, y, target)
		move(x+1, y, target)
		move(x, y-1, target)
		move(x, y+1, target)
	}
	for x := 0; x < len(inputGrid); x++ {
		for y := 0; y < len(inputGrid[0]); y++ {
			if inputGrid[x][y] != 0 {
				continue
			}
			move(x, y, 0)
			clear(visited)
		}
	}
	return score

}

func Part2(inputGrid [][]int) int {

	var score int
	var move func(int, int, int)
	visited := map[Cell]bool{}

	move = func(x int, y int, target int) {
		if x < 0 || x >= len(inputGrid) || y < 0 || y >= len(inputGrid[0]) || inputGrid[x][y] != target {
			return
		}
		if target == 9 && !visited[Cell{x, y}] {
			score++
			//visited[Cell{x, y}] = true
			return
		}
		target++
		move(x-1, y, target)
		move(x+1, y, target)
		move(x, y-1, target)
		move(x, y+1, target)
	}
	for x := 0; x < len(inputGrid); x++ {
		for y := 0; y < len(inputGrid[0]); y++ {
			if inputGrid[x][y] != 0 {
				continue
			}
			move(x, y, 0)
			//clear(visited)
		}
	}
	return score

}

func main() {

	start := time.Now()
	var inputGrid = ReadFile("./input.txt")
	//fmt.Println(inputGrid)
	//fmt.Println("Part 1 = ", Part1(inputGrid))
	fmt.Println("Part 2 = ", Part2(inputGrid))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
