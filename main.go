package main

import (
	"fmt"

	binaryTree "github.com/ceferrari/go-algods/binary_tree"
	hashTable "github.com/ceferrari/go-algods/hash_table"
	linkedList "github.com/ceferrari/go-algods/linked_list"
	"github.com/ceferrari/go-algods/utils"
)

func main() {
	op := -1
	for op != 0 {
		utils.ClearTerminal()
		utils.PrintDiv()
		fmt.Println("*** Algorithms & Data Structures ***")
		utils.PrintDiv()
		fmt.Printf("%2d | %s\n", 1, "Binary Tree")
		fmt.Printf("%2d | %s\n", 2, "Hash Table")
		fmt.Printf("%2d | %s\n", 3, "Linked List")
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
				binaryTree.Traverse()
				fmt.Scanln()
			case 2:
				hashTable.DoubleHashing()
				fmt.Scanln()
			case 3:
				linkedList.Menu()
			case 0:
				utils.Exit()
			default:
				utils.InvalidOption()
			}
		}
	}
}
