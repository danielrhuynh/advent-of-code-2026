package main

import (
	"aoc/types"
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	inputRaw, err := utils.ComsumeInputNewLines("day-9/input.txt")
	if err != nil {
		fmt.Println(("main: Cannot consume input"))
	}

	input := []types.XY{}
	for _, i := range inputRaw {
		split_slice := strings.Split(i[0], ",")
		x, err := strconv.Atoi(split_slice[0])
		if err != nil {
			fmt.Println("main: Error converting x to int")
		}
		y, err := strconv.Atoi(split_slice[1])
		if err != nil {
			fmt.Println("main: Error converting y to int")
		}
		input = append(input, types.XY{X: x, Y: y})
	}

	fmt.Println(input)

	// Strat:
		// The problem that we want to solve, is to find the largest rectangle we can make bounded by some choice of top left corner and bottom right corner.
		// Naive: The largest area possible will be made by the top left most corner and the bottom right most corner or the top right most corner and the bottom left most corner
		// There are various cases where this is not necessarily true
		// We could just scan and check the area of one point with every other point
	res := 0.0
	for i, p1 := range input {
		for _, p2 := range input[i:] {
			area := (math.Abs(float64(p1.X)-float64(p2.X))+1)*(math.Abs(float64(p1.Y)-float64(p2.Y))+1)
			res = math.Max(area, res)
		}
	}

	fmt.Println(int(res))
}
