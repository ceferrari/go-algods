package main

import (
	"fmt"

	"github.com/ceferrari/go-algods/graphs"
	"github.com/ceferrari/go-algods/lists"
	"github.com/ceferrari/go-algods/maps"
	"github.com/ceferrari/go-algods/queues"
	"github.com/ceferrari/go-algods/stacks"
	"github.com/ceferrari/go-algods/trees"
	"github.com/ceferrari/go-algods/utils"
)

var printGraph = func(i int) { fmt.Print(i, "->") }

func main() {
	op := -1
	for op != 0 {
		utils.ClearTerminal()
		utils.PrintDiv()
		fmt.Println("*** Algorithms & Data Structures ***")
		utils.PrintDiv()
		fmt.Printf("%2d | %s\n", 1, "[Maps] Double Hashing")
		fmt.Printf("%2d | %s\n", 2, "[Trees] Binary Search Tree")
		fmt.Printf("%2d | %s\n", 3, "[Lists] Singly Linked List")
		fmt.Printf("%2d | %s\n", 4, "[Lists] Doubly Linked List")
		fmt.Printf("%2d | %s\n", 5, "[Stacks] Stack")
		fmt.Printf("%2d | %s\n", 6, "[Queues] Queue")
		fmt.Printf("%2d | %s\n", 7, "[Graphs] DFS (Iterative)")
		fmt.Printf("%2d | %s\n", 8, "[Graphs] DFS (Recursive)")
		fmt.Printf("%2d | %s\n", 9, "[Graphs] BFS")
		fmt.Printf("%2d | %s\n", 0, "Exit")
		utils.PrintDiv()
		utils.ChooseOption()
		_, err := fmt.Scan(&op)
		utils.PrintDiv()
		if err != nil {
			utils.InvalidOption()
		} else {
			switch op {
			case 1:
				maps.DoubleHashing()
				fmt.Scanln()
			case 2:
				trees.BST()
				fmt.Scanln()
			case 3:
				lists.SinglyMenu()
			case 4:
				lists.DoublyMenu()
			case 5:
				stacks.RunStack()
				fmt.Scanln()
			case 6:
				queues.RunQueue()
				fmt.Scanln()
			case 7:
				graphs.DFS(1, printGraph)
				fmt.Scanln()
			case 8:
				graphs.DFSRecursive(1, map[int]bool{}, printGraph)
				fmt.Scanln()
			case 9:
				graphs.BFS(1, printGraph)
				fmt.Scanln()
			case 0:
				utils.Exit()
			default:
				utils.InvalidOption()
			}
		}
	}
}
