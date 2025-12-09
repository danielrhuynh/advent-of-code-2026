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
	rawInput, err := utils.ComsumeInputNewLines("day-8/test.txt")
	if err != nil {
		fmt.Println("main: error consuming input")
	}
	input := []types.XYZ{}
	for _, coor := range rawInput {
		split_slice := strings.Split(coor[0], ",")
		x, err := strconv.Atoi(split_slice[0])
		y, err := strconv.Atoi(split_slice[1])
		z, err := strconv.Atoi(split_slice[2])
		if err != nil {
			fmt.Println("main: Error converting strings to int")
		}
		input = append(input, types.XYZ{X: x, Y: y, Z: z})
	}
	fmt.Println(input)

	// var straightLineDistance func(x int, y int, z int)
	straightLineDistance := func(p1 types.XYZ, p2 types.XYZ) float64{
		dx := float64(p1.X - p2.X)
		dy := float64(p1.Y - p2.Y)
		dz := float64(p1.Z - p2.Z)
		return math.Sqrt(dx*dx + dy*dy + dz*dz)
	}

	edges := []types.Edge{}
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			dist := straightLineDistance(input[i], input[j])
			edges = append(edges, types.Edge{I: i, J: j, Dist: dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Dist < edges[j].Dist
	})

	dsu := types.InitDSU(len(input))
	maxEdges := 10
	// maxEdges := 1000
	if maxEdges > len(edges) {
		maxEdges = len(edges)
	}

	for i:=0; i<maxEdges; i++ {
		e := edges[i]
		merged, newSize := dsu.Union(e.I, e.J)
		fmt.Println(merged)
		fmt.Println(newSize)
	}

	sizes := []int{}
	for i := 0; i < len(input); i++ {
    	if dsu.Parent[i] == i {
        	sizes = append(sizes, dsu.Size[i])
    	}
	}
	fmt.Println(sizes)
	sort.Ints(sizes)
	k := len(sizes)
	product := sizes[k-1] * sizes[k-2] * sizes[k-3]
	fmt.Println(product)

}
