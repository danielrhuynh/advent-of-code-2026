package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input blocks are all 3x3s
	// this problem is np hard so we want a polynomial approximation (that I found on reddit...)
	// let S = the total number of # per region that we need to fill up sum of (num presents[i] * area of present[i])
	// let N = the total number of presents per region
	// let A = the area of a region

	// The idea here is that:
		// we have three cases:
		// if s > a it is impossible to fill up the area since we need more area than we have
		// if area/3 >= n it is guarenteed that we can fill the area
			// basically the logic here, is that a present is enclosed to a 3x3 space. area / 3 is the number of 3x3 spots we can put a
			// present in, regardless of orientation or trying to fit it
			// we can basically just put these presents next to each other
			// therefore is area/3 >= n we are guarenteed to be able to fit the region
		// else we don't know
	// Running the solver on our actual input, we end up having no unknown cases (by coincidence or by design, idk)
	// Therefore, we don't actually need to try to solve the np problem.
	presentMap := map[int]int{}
	// Test input
	// presentMap[0] = 7
	// presentMap[1] = 7
	// presentMap[2] = 7
	// presentMap[3] = 7
	// presentMap[4] = 7
	// presentMap[5] = 7

	// Actual input
	presentMap[0] = 7
	presentMap[1] = 5
	presentMap[2] = 7
	presentMap[3] = 7
	presentMap[4] = 7
	presentMap[5] = 6

	rawInput, err := utils.ComsumeInputNewLines("day-12/input.txt")
	if err != nil {
		fmt.Println("main: error parsing input")
	}
	fmt.Println(rawInput)

	guarenteed := 0
	unknown := 0
	for _, region := range rawInput {
		xIndex := strings.IndexByte(region[0], 'x')
		cIndex := strings.IndexByte(region[0], ':')
		l := strings.TrimSpace(region[0][:xIndex])
		r := strings.TrimSpace(region[0][xIndex+1:cIndex])
		li, _ := strconv.Atoi(l)
		ri, _ := strconv.Atoi(r)
		a := li*ri
		sum := 0
		s := 0
		counts := strings.Fields(region[0][cIndex+1:])
		for i := 0; i < len(counts); i++ {
			j, _ := strconv.Atoi(counts[i])
			s += j*presentMap[i]
			sum += j
		 }
		fmt.Println(a)
		fmt.Println(sum)
		if s > a {
			continue
		} else if a/3 >= sum {
			guarenteed += 1
		} else {
			unknown += 1
		}
	}

	fmt.Println(guarenteed)
	fmt.Println(unknown)
}
