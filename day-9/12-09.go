package main

import (
	"aoc/types"
	"aoc/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputRaw, err := utils.ComsumeInputNewLines("day-9/input.txt")
	if err != nil {
		fmt.Println(("main: Cannot consume input"))
	}

	input := []types.XY{}
	for _, i := range inputRaw {
		split_slice := strings.Split(i[0], ",")
		x, err := strconv.Atoi(split_slice[0])
		if err != nil {
			fmt.Println("main: Error converting x to int")
		}
		y, err := strconv.Atoi(split_slice[1])
		if err != nil {
			fmt.Println("main: Error converting y to int")
		}
		input = append(input, types.XY{X: x, Y: y})
	}

	fmt.Println(input)

	// Strat:
		// The problem that we want to solve, is to find the largest rectangle we can make bounded by some choice of top left corner and bottom right corner.
		// Naive: The largest area possible will be made by the top left most corner and the bottom right most corner or the top right most corner and the bottom left most corner
		// There are various cases where this is not necessarily true
		// We could just scan and check the area of one point with every other point
	res := 0.0
	for i, p1 := range input {
		for _, p2 := range input[i:] {
			area := (math.Abs(float64(p1.X)-float64(p2.X))+1)*(math.Abs(float64(p1.Y)-float64(p2.Y))+1)
			res = math.Max(area, res)
		}
	}

	fmt.Println(int(res))

	// Part 2:
		// Now, we can only switch out tiles that are red, or in the span of red tiles
		// The choice of red tiles need to be bounded by four corners that are either red tiles, or in one of the spans of red times (green)
		// You can choose any red tile to be the left position but the right tile needs to have a red tile in the same row before it
		// Strat:
			// We want to form a polygon with edges given these points (technically we don't do anything since our points bound the polygon)
			// We want to check that for a rectangle of [x_min, x_max]*[y_min, y_max], that for all x in the range [x_min, x_max] (all x's in the rectangle)...
				// y_min >= min_y[x] and y_max <= max_y[x] meaning that y's of the rectangle are in the vertical span of that x of the polygon
			// We need to do
				// Raycasting (PIP) to find the interior spans of the polygon
				// Still form every possible rectangle like we do in part 1 but then:
				// Use AABB to find if our rectangle is fully contained in our polygon
	minX := math.MaxInt
	maxX := math.MinInt

	for _, p := range input {
		minX = min(minX, p.X)
		maxX = max(maxX, p.X)
	}

	fmt.Println(minX, maxX)

	// Do the raytracting and determine the polygon boundaries
	intersections := map[int][]int{}
	// For all columns of the polygon
	for x:=minX; x <= maxX; x++ {
		// for all points in the polygon (which are in order of the order of which the edges should be connected)
		for i:=0; i<len(input); i++ {
			// if two adjacent points have the same Y there is a horizontal edge
			if input[i].Y == input[(i+1)%len(input)].Y {
				// We make the upper bound half open to avoid double counting vertices, this is fine since our AABB will check for interior points not the boundaries of the bb and polygon
				if min(input[i].X, input[(i+1)%len(input)].X) <= x && x < max(input[i].X, input[(i+1)%len(input)].X) {
					intersections[x] = append(intersections[x], input[i].Y)
				}
			}
		}
	}

	// Sort intersections
	for key := range intersections {
		sort.Slice(intersections[key], func(i, j int) bool {
			return intersections[key][i] < intersections[key][j]
		})
	}

	fmt.Println(intersections)
	newRes := 0.0
	// AABB
	for i, p1 := range input {
		for _, p2 := range input[i:] {
			isValid := true
			// For all x in the range of our rectangle
			for x:=min(p1.X, p2.X); x < max(p1.X, p2.X); x++ {
				ys := intersections[x]
				if len(ys) < 2 {
					isValid = false
					break
				}

				// Are we valid for any band?
				ok := false
				for j:=0; j<=len(ys); j+=2 {
					if min(p1.Y, p2.Y) >= intersections[x][j] && max(p1.Y, p2.Y) <= intersections[x][j+1] {
						ok = true
						break
					}
				}
				if !ok {
					isValid = false
					break
				}
			}
			// If we are valid see if this is a new best area
			if isValid {
				area := (math.Abs(float64(p1.X)-float64(p2.X))+1)*(math.Abs(float64(p1.Y)-float64(p2.Y))+1)
				newRes = math.Max(area, newRes)
			}
		}
	}

	fmt.Println(int(newRes))
}
