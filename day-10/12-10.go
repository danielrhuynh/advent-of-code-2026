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
        lcur := strings.IndexByte(line, '{')
        rcur := strings.LastIndexByte(line, '}')

        reqStr := strings.TrimSpace(line[lcur+1 : rcur])
        reqParts := strings.Split(reqStr, ",")

        req := make([]int, 0, len(reqParts))
        for _, p := range reqParts {
            p = strings.TrimSpace(p)
            v, err := strconv.Atoi(p)
            if err != nil {
                return nil, fmt.Errorf("bad req value %q in line %q: %w", p, line, err)
            }
            req = append(req, v)
        }
		machines = append(machines, types.Machine{Diagram: diagram, Buttons: buttons, Req: req})
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
		We can do this by making our matrix and vectors and putting our system into RREF to get our pivot variables
		and our free variables, but then we need to choose our free variables to minimize the overall number
		of 1's we choose (wtf we should use an optimization solver probably)
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

	// Actually doing the solving
	// Basically we want to put our matrix into RREF and solve for the value at each button
	// If a button is a free variable, we search for the minimum solution
	total := 0
	for _, machine := range machines {
		n := len(machine.Diagram)
		m := len(machine.Buttons)

		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = make([]int, m)
		}

		b := make([]int, n)
		for i, char := range machine.Diagram {
			if char == '#' {
				b[i] = 1
			}
		}

		for i, button := range machine.Buttons {
			for _, index := range button {
				A[index][i] = 1 - A[index][i]
			}
		}
		fmt.Println(A)
		fmt.Println(b)

		presses, err := minHammingWeight(A, b)
    	if err != nil {
     	   fmt.Println("main: unsatisfiable machine:", err)
    	}
    	total += presses
	}
	fmt.Println("answer:", total)
}

func minHammingWeight(A [][]int, b []int) (int, error) {
	n := len(A)
	m := len(A[0])

	M := cloneMatInt(A)
	rhs := cloneVecInt(b)

	where := make([]int, m)
	for j := range where {
		where[j] = -1
	}

	pivotRow := 0

	for col := 0; col < m && pivotRow < n; col++ {
		sel := -1
		for r := pivotRow; r < n; r++ {
			if M[r][col] == 1 {
				sel = r
				break
			}
		}

		// Free variable
		if sel == -1 {
			continue
		}

		M[pivotRow], M[sel] = M[sel], M[pivotRow]
		rhs[pivotRow], rhs[sel] = rhs[sel], rhs[pivotRow]

		where[col] = pivotRow

		for r := 0; r < n; r++ {
			if r == pivotRow {
				continue
			}
			if M[r][col] == 1 {
				for j := col; j < m; j++ {
					M[r][j] ^= M[pivotRow][j]
				}
				rhs[r] ^= rhs[pivotRow]
			}
		}
		pivotRow++
	}

	for i := 0; i < n; i++ {
		allZero := true
		for j := 0; j < m; j++ {
			if M[i][j] == 1 {
				allZero = false
				break
			}
		}
		if allZero && rhs[i] == 1 {
			return 0, fmt.Errorf("no solution (inconsistent constraints)")
		}
	}

	// I lose the solution about here... wtf
	freeCols := make([]int, 0)
	for col := 0; col < m; col++ {
		if where[col] == -1 {
			freeCols = append(freeCols, col)
		}
	}

	x0 := make([]int, m)
	for col := 0; col < m; col++ {
		if where[col] != -1 {
			x0[col] = rhs[where[col]]
		}
	}

	basis := make([][]int, 0, len(freeCols))
	for _, f := range freeCols {
		v := make([]int, m)
		v[f] = 1
		for p := 0; p < m; p++ {
			r := where[p]
			if r == -1 {
				continue
			}
			if M[r][f] == 1 {
				v[p] = 1
			}
		}
		basis = append(basis, v)
	}

	best := hammingWeightInt01(x0)
	d := len(basis)
	if d == 0 {
		return best, nil
	}

	limit := 1 << d
	for mask := 1; mask < limit; mask++ {
		x := cloneVecInt(x0)
		for k := 0; k < d; k++ {
			if (mask>>k)&1 == 1 {
				xorVecInPlaceInt01(x, basis[k])
			}
		}
		w := hammingWeightInt01(x)
		if w < best {
			best = w
		}
	}

	return best, nil
}

func xorVecInPlaceInt01(dst, src []int) {
	for i := range dst {
		dst[i] ^= src[i]
	}
}

func hammingWeightInt01(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func cloneVecInt(x []int) []int {
	y := make([]int, len(x))
	copy(y, x)
	return y
}

func cloneMatInt(A [][]int) [][]int {
	B := make([][]int, len(A))
	for i := range A {
		B[i] = cloneVecInt(A[i])
	}
	return B
}
