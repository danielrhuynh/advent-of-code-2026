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
	maxLen := 0
	for _, line := range input {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	formattedInput := [][]string{}
	opSlice := []string{}
	isOpCol := []bool{}
	for col := maxLen - 1; col >= 0; col-- {
		columnChars := make([]string, 0, len(input))
		hasOp := false
		for row := 0; row < len(input); row++ {
			line := input[row]
			if col < len(line) {
				val := string(line[col])
				if val == "*" || val == "+" {
					 if !hasOp {
						opSlice = append(opSlice, val)
    				}
					hasOp = true
					continue
				}
				if val != " " {
					columnChars = append(columnChars, string(val))
				}
			}
		}
		if hasOp || len(columnChars) > 0 {
			formattedInput = append(formattedInput, columnChars)
        	isOpCol = append(isOpCol, hasOp)
		}
	}

	fmt.Println(formattedInput)
	fmt.Println(opSlice)

	preprocessedInput := [][]int{}
	localSlice := []int{}
	for i, col := range formattedInput {
		if len(col) > 0 {
			s := strings.Join(col, "")
			val, err := strconv.Atoi(s)
    		if err != nil {
    		    fmt.Println("main: error converting column to int:", s)
    		    return
    		}
    		localSlice = append(localSlice, val)
		}
		if isOpCol[i] {
			if len(localSlice) > 0 {
				preprocessedInput = append(preprocessedInput, localSlice)
				localSlice = []int{}
			}
		}
	}

	if len(localSlice) > 0 {
		preprocessedInput = append(preprocessedInput, localSlice)
	}

	fmt.Println(preprocessedInput)

	res := 0
	for i:=0; i < len(preprocessedInput); i++ {
		op := opSlice[i]
		col := preprocessedInput[i]
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
