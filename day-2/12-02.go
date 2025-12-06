package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumRange struct {
	Start int
	End int
}

func comsumeInput(path string) ([]NumRange, error ) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var res []NumRange

	if ! scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	line := scanner.Text()
	ranges := strings.Split(line, ",")

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}

		pair := strings.SplitN(r, "-", 2)
		if len(pair) != 2 {
			continue
		}
		start, err := strconv.Atoi(pair[0])
		if err != nil {
			return nil, fmt.Errorf("bad start %q: %w", pair[0], err)
		}
		end, err := strconv.Atoi(pair[1])
		if err != nil {
			return nil, fmt.Errorf("bad end %q: %w", pair[1], err)
		}

		res = append(res, NumRange{Start: start, End: end})
	}

	return res, nil
}

func isInvalidId(id int) bool {
	s := strconv.Itoa(id)
	n := len(s)
	if n%2 != 0 {
        return false
    }
    half := n / 2
    return s[:half] == s[half:]
}

func isInvalidIdSlidingBlock(id int) bool {
    s := strconv.Itoa(id)
    n := len(s)

    for size := 1; size <= n/2; size++ {
        if n%size != 0 {
            continue
        }

        pattern := s[:size]
        ok := true

        for i := size; i < n; i += size {
            if s[i:i+size] != pattern {
                ok = false
                break
            }
        }

        if ok {
            return true
        }
    }

    return false
}


func main() {
	ranges, _ := comsumeInput("day-2/input.txt")
	// fmt.Println(ranges)

	/*
	Strat:
		Checking for an invalid ID:
			An invalid id is an id that has a sequence of digits repeated twice.
			Checking for a repeated substring per ID in the range?
			Naive: Can you split an ID in half and check if the first half and the last half are approximately equal?
			What about repeated substrings that don't span the entire string?

	 */
	 res := 0

	 for _, r := range ranges {
		for i:=r.Start; i <= r.End; i++ {
			if (isInvalidId(i)) {
				res += i
			}
		}
	 }
	 fmt.Println(res)

	 // In summary, this problem was simpler than I originally thought, the repeated block needs to span the entire string, we don't need a sliding window lol

	 // PART 2
	 // We want to now implement our sliding window approach
	 res2 := 0
	 for _, r := range ranges {
		for i:=r.Start; i <= r.End; i++ {
			if (isInvalidIdSlidingBlock(i)) {
				res2 += i
			}
		}
	 }
	 fmt.Println(res2)
}
