package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func main() {
	inputRange, err1 := utils.ComsumeInputRange("day-5/input_range.txt")
	inputInput, err2 := utils.ComsumeInput("day-5/input_input.txt")
	if err1 != nil {
		fmt.Println("main: error comsuming input range")
	}

	if err2 != nil {
		fmt.Println("main: error consuming input input")
	}
	fmt.Println(inputRange)
	fmt.Println(inputInput)

	// We could in theory sort intervals and then perform the merge intervals leetcode question where we check all id's against
	// the single interval
	// Fuck that it's 12:19am I want to go to bed, brute force all the way
	res := 0
	for _, i := range inputInput {
		ii, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("main: error converting input string to int")
		}
		invalid := true
		for _, nr := range inputRange {
			if ii >= nr.Start && ii <= nr.End {
				invalid = false
			}
		}
		if !invalid {
			res += 1
		}
	}
	fmt.Println(res)

	mergedIntervals := utils.Merge(inputRange)
	fmt.Println(mergedIntervals)
	totalIds := 0
	for _, interval := range mergedIntervals {
		totalIds += interval.End - interval.Start + 1
	}
	fmt.Println(totalIds)

}
