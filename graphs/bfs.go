package graphs

import (
	"fmt"

	stackQueue "github.com/ceferrari/go-algods/queues"
)

var bfsAdj = map[int][]int{
	1: {2, 3, 4},
	2: {5, 6},
	3: {},
	4: {7, 8},
	5: {9, 10},
	6: {10},
	7: {11, 12},
	8: {},
}

func BFS(start int) {
	visited := map[int]bool{}
	stack := stackQueue.Queue[int]()
	stack.Enqueue(start)
	for !stack.IsEmpty() {
		node := stack.Dequeue()
		if !visited[node] {
			visited[node] = true
			for _, neighbor := range bfsAdj[node] {
				stack.Enqueue(neighbor)
			}
			fmt.Print(node, "->")
		}
	}
}
