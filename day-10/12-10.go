package main

import (
	"aoc/types"
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseMachine(lines []string) ([]types.Machine, error) {
	machines := make([]types.Machine, 0, len(lines))
	for _, line := range lines {
		lb := strings.IndexByte(line, '[')
		rb := strings.IndexByte(line, ']')

		diagram := line[lb+1:rb]

		rest := line[rb+1:]
		cb := strings.IndexByte(rest, '{')
		rest = rest[:cb]

		buttons := [][]int{}
		for {
			ob := strings.IndexByte(rest, '(')
			if ob == -1 {
				break
			}
			cr := strings.IndexByte(rest[ob+1:], ')')
			close := ob + 1 + cr
			contents := strings.TrimSpace(rest[ob+1:close])

			parts := strings.Split(contents, ",")
			btn := make([]int, 0, len(parts))
			for _, p := range parts {
				p = strings.TrimSpace(p)
				idx, err := strconv.Atoi(p)
				if err != nil {
					fmt.Println("main: error parsing button contents")
				}
				btn =  append(btn, idx)
			}
			buttons = append(buttons, btn)
			rest = rest[close+1:]
		}
		machines = append(machines, types.Machine{Diagram: diagram, Buttons: buttons})
	}
	return machines, nil
}

func main() {
	/*
		For any button only two cases matter
		if a button is pressed an even amount of times it is 0
		if a button is pressed an odd amount of times it is 1
		so each button is a binary decision
		Therefore, each line is a linear system over a vector space of the set {0, 1}
		In this vector space, addition is XOR and multiplication is AND
		Creating a linear system...
		 	x = 0 or 1 depending on the amount of times the button is pressed
			n = # lights which is the len(square brackets)
			m = # buttons for that machine (number of parens)
			b = target vector where # is a 1 and . is a 0
			A = matrix where A_ij = 1 if button j toggles light 1
		In additon to a solution, we also need the minimum number of pressed
		where our cost is the sum of x over j
		So we want to solve Ax=b and minimize the hamming weight
	*/
	path := "day-10/test.txt"
	lines, err := utils.ComsumeInput(path)
	if err != nil {
		fmt.Println("main: error parsing input")
	}
	fmt.Println(lines)
	machines, err := parseMachine(lines)
	if err != nil {
		fmt.Println("main: error parsing machines")
	}
	fmt.Println(machines)

	// Convert to bit masks
	compiledMachines := []types.CompiledMachine{}
	for _, m := range machines {
		n := len(m.Diagram)
		target := types.NewMask(n)
		buttonMasks := []types.Mask{}
		for i, char := range m.Diagram {
			if char == '#' {
				target.Toggle(i)
			}
		}

		for _, button := range m.Buttons {
			bm := types.NewMask(n)
			for _, index := range button {
				bm.Toggle(index)
			}
			buttonMasks = append(buttonMasks, bm)
		}
		compiledMachines = append(compiledMachines, types.CompiledMachine{N: n, Target: target, Buttons: buttonMasks})
	}
	fmt.Println(compiledMachines)

	// Actually doing the solving

}

