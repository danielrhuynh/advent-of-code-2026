package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	input, err := comsumeInput("day-4/input.txt")
	if err != nil {
		fmt.Println("main: error consuming input")
	}

	var charSlice [][]string

	for _, s := range input {
		var runeSlice []string

		for _, r := range s {
			runeSlice = append(runeSlice, string(r))
		}

		charSlice = append(charSlice, runeSlice)
	}

	fmt.Println(charSlice)


	directions := [][]int{{1, 1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}}
	// Broadly:
		// check all directions for a given index
		// if the direction is in bounds += to current sum
		// if current sum < 4: add to larger sum
	rows := len(charSlice)
	cols := len(charSlice[0])
	res := 0
	flag := true
	for flag {
		flag = false
		for i:=0; i<rows; i++ {
			for j:=0; j<cols; j++ {
				if charSlice[i][j] != "@" {
					continue
				}
				surroundingTP := 0
				for _, direction := range directions {
					newI := i + direction[0]
					newJ := j + direction[1]

					if newI >= 0 && newI < rows && newJ >= 0 && newJ < cols && charSlice[newI][newJ] == "@" {
						surroundingTP++
					}
				}
				if surroundingTP < 4 {
					res++
					flag = true
					charSlice[i][j] = "."
				}
			}
		}
	}
	fmt.Println(res)
}
