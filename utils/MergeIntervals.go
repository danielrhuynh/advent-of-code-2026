package utils

import (
	"aoc/types"
	"sort"
)

func Merge(intervals []types.NumRange) []types.NumRange {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})

	res := make([]types.NumRange, 0, len(intervals))
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {

		if intervals[i].Start <= (&res[len(res)-1]).End {
			if intervals[i].Start < (&res[len(res)-1]).Start {
				(&res[len(res)-1]).Start = intervals[i].Start
			}
			if intervals[i].End > (&res[len(res)-1]).End {
				(&res[len(res)-1]).End = intervals[i].End
			}
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}
