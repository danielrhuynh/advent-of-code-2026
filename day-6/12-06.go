package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ComsumeInput("day-6/input.txt")
	if err != nil {
		fmt.Println("main: error consuming input")
	}
	formattedInput := [][]string{}
	for _, row := range input {
		nums := strings.Fields(row)
		formattedInput = append(formattedInput, nums)
	}
	fmt.Println(formattedInput)

	numEquations := len(formattedInput[0])
	numberSlice := make([][]int, numEquations)
	opSlice := make([]string, 0, numEquations)

	for _, row := range formattedInput {
		for i:= 0; i < numEquations; i++ {
			val, err := strconv.Atoi(row[i])
			if err != nil {
				fmt.Println("main: we are on an operator")
				opSlice = append(opSlice, row[i])
				continue
			}
			numberSlice[i] = append(numberSlice[i], val)
		}
	}

	fmt.Println(numberSlice)
	fmt.Println(opSlice)

	res := 0
	for i:=0; i < numEquations; i++ {
		op := opSlice[i]
		col := numberSlice[i]
		switch op {
		case "+":
			localRes := 0
			for _, num := range col {
				localRes += num
			}
			res += localRes
		case "*" :
			localRes := 1
			for _, num := range col {
				localRes *= num
			}
			res += localRes
		}
	}
	fmt.Println(res)
}
