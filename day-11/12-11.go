package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	// Strat:
		// 1D Graph question
		// We want to find every path starting from node you to node out
		// We can parse the input as an adjacency list (node -> [connections])
		// Then run a dfs on the graph starting at you, incrementing res everytime we end up at out
	lineInput, err := utils.ComsumeInput("day-11/test.txt")
	if err != nil {
		fmt.Println("main: failed to read input")
	}
	fmt.Println(lineInput)

	// Parse line input into a map
	adj := make(map[string][]string)
	for _, ns := range lineInput {
		colonIndex := strings.IndexByte(ns, ':')
		node := strings.TrimSpace(ns[:colonIndex])
		neighbours := strings.Fields(ns[colonIndex+1:])
		adj[node] = neighbours
	}
	fmt.Println(adj)
	// Run a DFS starting at you until you find out
	// we want to maintain a seen set to prevent cycles?
	res := 0
	seen := map[string]struct{}{}
	var dfs func(node string)
	 dfs = func(node string) {
		if _, ok := seen[node]; ok {
			return
		}
		if node == "out" {
			res++
			return
		}

		seen[node] = struct{}{}
		for _, neighbour := range adj[node] {
			dfs(neighbour)
		}
		delete(seen, node)
	}
	dfs("you")
	fmt.Println(res)

	// reverse adjacency list

}
