package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	// Strat:
		// We want the number of times a beam originiating in the middle top of the input is split
		// an input is split if a splitter = "^" is below a beam
		// This creates two beams to the left and right (indices i-1 and i+1)
		//
		// What this means in practice:
		// Assuming we have our input in matrix form already
		// Keep track of indices where a beam should be to determine if a beam coincides with a splitter
		// beams = [i_0, ..., i_n]
		// for j in range of cols:
			// If input[i][j] == "^"
			// beams.append(i-1)
			// beams.append(i+1)
			// res += 1
		// return res

	input, err := utils.ConsumeInputMatrix("day-7/input.txt")
	if err != nil {
		fmt.Println("main: error consuming input as matrix")
	}
	sIndex := 0
	for i, s := range input[0] {
		if s == "S" {
			sIndex = i
		}
	}

	// cols := len(input[0])
	rows := len(input)
	// beams := []int{} // Use a set
	// beams = append(beams, sIndex)
	beams := map[int]struct{}{}
	beams[sIndex] = struct{}{}
	res := 0
	for j:=0; j<rows; j++ {
		for k := range beams {
			if input[j][k] == "^" {
				// check for index in slice here
				// beams = append(beams, i-1)
				// beams = append(beams, i+1)
				beams[k-1] = struct{}{}
				beams[k+1] = struct{}{}
				delete(beams, k)
				res += 1
			}
		}
	}

	fmt.Println(res)
}
