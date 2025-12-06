package utils

import (
	"aoc/types"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ComsumeInputRange(path string) ([]types.NumRange, error ) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var res []types.NumRange

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
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

			res = append(res, types.NumRange{Start: start, End: end})
		}
	}

	return res, nil
}
