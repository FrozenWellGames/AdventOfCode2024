package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

// Set is a collection of unique elements
type antenna struct {
	key    string
	coOrds []coordinates
}

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

func indexOf(element string, data []antenna) int {
	for i := 0; i < len(data); i++ {
		if data[i].key == element {
			return i
		}
	}
	return -1
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

func formatData(input []string) [][]string {
	returnSlice := [][]string{}
	for i := 0; i < len(input); i++ {
		returnSlice = append(returnSlice, strings.Split(input[i], ""))
	}
	return returnSlice
}

func createSets(formattedData [][]string) []antenna {
	var uniqueAntennas = []string{}
	var antennas = []antenna{}

	for x := 0; x < len(formattedData); x++ {
		for y := 0; y < len(formattedData[x]); y++ {
			if formattedData[x][y] != "." {
				if !slices.Contains(uniqueAntennas, formattedData[x][y]) {
					uniqueAntennas = append(uniqueAntennas, formattedData[x][y])

					var test = antenna{key: formattedData[x][y], coOrds: []coordinates{}}
					test.coOrds = append(test.coOrds, coordinates{x: x, y: y})
					antennas = append(antennas, test)
				} else {
					localCoOrds := coordinates{x: x, y: y}
					indexOfAntenna := indexOf(formattedData[x][y], antennas)
					if indexOfAntenna != -1 {
						antennas[indexOfAntenna].coOrds = append(antennas[indexOfAntenna].coOrds, localCoOrds)
					}
				}

			}
		}
	}

	return antennas
}

func part1(antennas []antenna, sliceX int, sliceY int) {
	fmt.Println(sliceY, ",", sliceX)
	var sliceOfAntinodes = []coordinates{}
	fmt.Println("antennas = ", antennas)
	for keyIndex := 0; keyIndex < len(antennas); keyIndex++ {

		for a := 0; a < len(antennas[keyIndex].coOrds)-1; a++ {
			for b := a + 1; b < len(antennas[keyIndex].coOrds); b++ {
				x1 := antennas[keyIndex].coOrds[a].x
				y1 := antennas[keyIndex].coOrds[a].y
				x2 := antennas[keyIndex].coOrds[b].x
				y2 := antennas[keyIndex].coOrds[b].y

				stepX := x2 - x1
				stepY := y2 - y1

				antinode1 := coordinates{x1 - stepX, y1 - stepY}
				antinode2 := coordinates{x2 + stepX, y2 + stepY}

				if antinode1.x >= 0 && antinode1.x < sliceX && antinode1.y >= 0 && antinode1.y < sliceY {
					if !slices.Contains(sliceOfAntinodes, antinode1) {
						sliceOfAntinodes = append(sliceOfAntinodes, antinode1)
					}
				}
				if antinode2.x >= 0 && antinode2.x < sliceX && antinode2.y >= 0 && antinode2.y < sliceY {
					if !slices.Contains(sliceOfAntinodes, antinode2) {
						sliceOfAntinodes = append(sliceOfAntinodes, antinode2)
					}
				}
			}

		}
	}
	fmt.Println("sliceOfAntinodes ", sliceOfAntinodes)

	fmt.Println("len(sliceOfAntinodes) ", len(sliceOfAntinodes))
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 8")
	var input = ReadFile("./input.txt")
	formattedData := formatData(input)
	antennas := createSets(formattedData)

	part1(antennas, len(input), len(input[0]))

}
