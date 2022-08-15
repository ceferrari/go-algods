package graphs

import (
	stackQueue "github.com/ceferrari/go-algods/stacks"
)

var dfsAdj = map[int][]int{
	1: {2, 3, 4},
	2: {5, 6},
	3: {},
	4: {7, 8},
	5: {9, 10},
	6: {10},
	7: {11, 12},
	8: {},
}

func DFS(start int, function func(int)) {
	visited := map[int]bool{}
	stack := stackQueue.Stack[int]()
	stack.Push(start)
	for !stack.IsEmpty() {
		node := stack.Pop()
		if !visited[node] {
			visited[node] = true
			function(node)
			for _, neighbor := range dfsAdj[node] {
				stack.Push(neighbor)
			}
		}
	}
}

func DFSRecursive(node int, visited map[int]bool, function func(int)) {
	if !visited[node] {
		visited[node] = true
		function(node)
		for _, neighbor := range dfsAdj[node] {
			DFSRecursive(neighbor, visited, function)
		}
	}
}
