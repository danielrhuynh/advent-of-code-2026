package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func comsumeInput(path string) ([]string, error ) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var res []string

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		line := scanner.Text()
		res = append(res, line)
	}

	return res, nil
}

func bestIndices(battery string, L int) []int {
    indices := []int{}

    start := 0

	// i < N - (L - k)
	// Where our current search space ends at 15-12-k where k is the current step we are on
	// This embodies the edge cases outlined:
		// If the first index is 0 we can search -> 3
		// If the first index is 3 -> 4 (basically the last possible choice is 3 so its 3 + k for the kth choice)
    for k := 0; k < L; k++ {
        end :=  len(battery) - (L - k)
        currMax := -1
        bestIdx := -1
        for i := start; i <= end; i++ {
            digit := int(battery[i] - '0')
            if digit > currMax {
                currMax = digit
                bestIdx = i
            }
        }
        indices = append(indices, bestIdx)
        start = bestIdx + 1
    }
    return indices
}

func main() {
	path := "day-3/input.txt"
	input, err := comsumeInput(path)
	if err != nil {
		fmt.Println("main: Error reading input")
	}
	fmt.Println(input)
	res := []int{}

	// Naive: We could just do 2 passes on every input
	// Store the indices of the 2 maxes and add to buffer in order
	// New idea: on the first pass iterate from 0 -> n-1, on the second pass iterate from firstMaxIndex -> n
	// For part 2 we need to modify our search space
		// For the first iteration the search space is from 0 to len(battery)-11
		// The next iteration is from lastIndex+1 to:
			// If the first index is 0 -> 5
			// If the first index is 4 -> 5
	locations := [][]int{}
	for _, battery := range input {
		indices := bestIndices(battery, 12)
		locations = append(locations, indices)
	}
	fmt.Println(locations)
	for i, location := range locations {
		result := ""
		for _, j := range location {
			byte := string(input[i][j])
			result += byte
		}
		iRes, err := strconv.Atoi(result)
		if err != nil {
			fmt.Println(("main: Error converting result string to an int"))
		}

		res = append(res, iRes)
	}
	fmt.Println(res)
	sum := 0
	for _, quant := range res {
		sum += quant
	}
	fmt.Println(sum)
}
