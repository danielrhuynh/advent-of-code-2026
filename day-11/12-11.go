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
	lineInput, err := utils.ComsumeInput("day-11/input.txt")
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
	// res := 0
	// seen := map[string]struct{}{}
	// var dfs func(node string)
	//  dfs = func(node string) {
	// 	if _, ok := seen[node]; ok {
	// 		return
	// 	}
	// 	if node == "out" {
	// 		res++
	// 		return
	// 	}

	// 	seen[node] = struct{}{}
	// 	for _, neighbour := range adj[node] {
	// 		dfs(neighbour)
	// 	}
	// 	delete(seen, node)
	// }
	// dfs("you")
	// fmt.Println(res)

	// ########## PART 2 ##########
	// reverse adjacency list
	rev := make(map[string][]string)

	for u, nbrs := range adj {
 	   for _, v := range nbrs {
        	rev[v] = append(rev[v], u)
        	if _, ok := rev[v]; !ok {
            	rev[v] = nil
        	}
    	}
	}

	fmt.Println(rev)
	canReachOut := map[string]bool{}
	canReachDAC := map[string]bool{}
	canReachFFT := map[string]bool{}
	var reverseDFS func(node string, reach map[string]bool)
	reverseDFS = func(node string, reach map[string]bool) {
		if _, ok := reach[node]; ok {
			return
		}
		reach[node] = true
		for _, neighbour := range rev[node] {
			reverseDFS(neighbour, reach)
		}
	}

	var forwardDFS func(node string, reach map[string]bool)
	forwardDFS = func(node string, reach map[string]bool) {
		if _, ok := reach[node]; ok {
			return
		}
		reach[node] = true
		for _, neighbour := range adj[node] {
			forwardDFS(neighbour, reach)
		}
	}
	reverseDFS("out", canReachOut)
	reverseDFS("dac", canReachDAC)
	reverseDFS("fft", canReachFFT)

	reachableFromSVR := map[string]bool{}
	forwardDFS("svr", reachableFromSVR)

	allNodes := map[string]struct{}{}
	for u, nbrs := range adj {
		allNodes[u] = struct{}{}
		for _, v := range nbrs {
			allNodes[v] = struct{}{}
		}
	}

	active := map[string]bool{}
	for n := range allNodes {
		active[n] = canReachOut[n] && reachableFromSVR[n]
	}

	// 0 unvisited (missing key), 1 pending, 2 visited
	state := map[string]uint8{}
	order := make([]string, 0)

	var topoDFS func(u string) bool
	topoDFS = func(u string) bool {
		if !active[u] {
			return true
		}
		switch state[u] {
		case 1:
			return false
		case 2:
			return true
		}
		state[u] = 1
		for _, v := range adj[u] {
			if !active[v] {
				continue
			}
			if !topoDFS(v) {
				return false
			}
		}
		state[u] = 2
		order = append(order, u)
		return true
}

	for n := range allNodes {
		if active[n] && state[n] == 0 {
			if !topoDFS(n) {
				fmt.Println("cycle detected in relevant subgraph; DAG DP not applicable")
				return
			}
		}
	}

	updateState := func(node string, s int) int {
		if node == "dac" {
			switch s {
			case 0:
				s = 1
			case 2:
				s = 3
			}
		}
		if node == "fft" {
			switch s {
			case 0:
				s = 2
			case 1:
				s = 3
			}
		}
		return s
	}

	dp0 := map[string]int64{}
	dp1 := map[string]int64{}
	dp2 := map[string]int64{}
	dp3 := map[string]int64{}
	for n := range allNodes {
		if active[n] {
			dp0[n], dp1[n], dp2[n], dp3[n] = 0, 0, 0, 0
		}
	}

	startState := updateState("svr", 0)
	switch startState {
	case 0:
		dp0["svr"] = 1
	case 1:
		dp1["svr"] = 1
	case 2:
		dp2["svr"] = 1
	case 3:
		dp3["svr"] = 1
	}

	// Iterate backwards bc of topological sort
	for i := len(order) - 1; i >= 0; i-- {
		u := order[i]
		if !active[u] {
			continue
		}
		w0, w1, w2, w3 := dp0[u], dp1[u], dp2[u], dp3[u]
		if w0 == 0 && w1 == 0 && w2 == 0 && w3 == 0 {
			continue
		}
		for _, v := range adj[u] {
			if !active[v] {
				continue
			}
			if w0 != 0 {
				ns := updateState(v, 0)
				if ns == 0 { dp0[v] += w0 } else if ns == 1 { dp1[v] += w0 } else if ns == 2 { dp2[v] += w0 } else { dp3[v] += w0 }
			}
			if w1 != 0 {
				ns := updateState(v, 1)
				if ns == 0 { dp0[v] += w1 } else if ns == 1 { dp1[v] += w1 } else if ns == 2 { dp2[v] += w1 } else { dp3[v] += w1 }
			}
			if w2 != 0 {
				ns := updateState(v, 2)
				if ns == 0 { dp0[v] += w2 } else if ns == 1 { dp1[v] += w2 } else if ns == 2 { dp2[v] += w2 } else { dp3[v] += w2 }
			}
			if w3 != 0 {
				ns := updateState(v, 3)
				if ns == 0 { dp0[v] += w3 } else if ns == 1 { dp1[v] += w3 } else if ns == 2 { dp2[v] += w3 } else { dp3[v] += w3 }
			}
		}
	}
	fmt.Println(dp3["out"])
}
