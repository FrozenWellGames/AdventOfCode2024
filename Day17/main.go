package main

import (
	"fmt"
	"strconv"
)

func convertIntToString(s int) string {
	returnString := strconv.Itoa(s)

	return returnString
}

func partOne(program []int, value int) string {
	var res []int
	var instructionPointer int
	registers := map[rune]int{
		'A': value,
		'B': 0,
		'C': 0,
	}

	for instructionPointer < len(program)-1 {
		operand := program[instructionPointer+1]

		switch operator := program[instructionPointer]; operator {
		case 0: // opcode  adv
			registers['A'] >>= getComboOperandValue(operand, registers)
		case 1: // opcode  bxl
			registers['B'] ^= operand
		case 2: // opcode  bst
			registers['B'] = getComboOperandValue(operand, registers) & 7
		case 3: // opcode jnz
			if registers['A'] != 0 {
				instructionPointer = operand
				continue
			}
		case 4: // opcode bxc
			registers['B'] ^= registers['C']
		case 5: // opcode out
			val := getComboOperandValue(operand, registers) & 7
			res = append(res, val)
		case 6: // opcode bdv
			registers['B'] = registers['A'] >> getComboOperandValue(operand, registers)
		case 7: // opcode cdv
			registers['C'] = registers['A'] >> getComboOperandValue(operand, registers)
		}
		instructionPointer += 2
	}
	var returnString string = ""
	for i := 0; i < len(res); i++ {
		returnString += convertIntToString(res[i])
		if i < len(res)-1 {
			returnString += ","
		}
	}

	return returnString
}

func getComboOperandValue(comboOperand int, registers map[rune]int) (value int) {
	if comboOperand >= 0 && comboOperand <= 3 {
		value = comboOperand
	} else if comboOperand == 4 {
		value = registers['A']
	} else if comboOperand == 5 {
		value = registers['B']
	} else if comboOperand == 6 {
		value = registers['C']
	}
	return
}

func main() {
	fmt.Println("Advent Of Code 2024 - Day 17")
	//fmt.Println("Part One:", partOne([]int{0, 1, 5, 4, 3, 0}, 729))
	fmt.Println("Part One:", partOne([]int{2, 4, 1, 1, 7, 5, 1, 5, 0, 3, 4, 4, 5, 5, 3, 0}, 23999685))

}
