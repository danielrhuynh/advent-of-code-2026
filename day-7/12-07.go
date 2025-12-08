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
		// We actually need to use a set or something unique here such that we dont count the same beam twice
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

	rows := len(input)
	cols := len(input[0])
	beams := map[int]struct{}{}
	beams[sIndex] = struct{}{}
	res := 0
	for j:=0; j<rows; j++ {
		for k := range beams {
			if input[j][k] == "^" {
				beams[k-1] = struct{}{}
				beams[k+1] = struct{}{}
				delete(beams, k)
				res += 1
			}
		}
	}

	fmt.Println(res)

	// Part 2:
		// This is idealistically 2^splits but not in practice
		// Write a dfs where we count the amount of times we finish the frame... fuck
		// dfs(i, j):
			// Base case:
			// if j > rows:
				// res += 1
				// return
			// if input[i][j] == "^":
				// dfs(i-1, j+1)
				// dfs(i+1, j+1)
			// else:
				// dfs(i, j+1)

	// We need to memoize our input since this takes way too long to run (omg I get to use DP)
	timelines := 0
	cache := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		cache[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(i int, j int) int
	 dfs = func(i int, j int) int {
		if i >= rows {
			return 1
		}
		if cache[i][j] != -1 {
        	return cache[i][j]
    	}
		var res int
		if input[i][j] == "^" {
			res = dfs(i+1, j+1) + dfs(i+1, j-1)
		} else {
			res = dfs(i+1, j)
		}
		cache[i][j] = res
		return res
	}

	timelines = dfs(0, sIndex)
	fmt.Println(timelines)
}


