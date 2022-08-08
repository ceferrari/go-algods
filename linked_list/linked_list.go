package linked_list

import (
	"fmt"

	"github.com/ceferrari/go-algods/utils"
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

type LinkedList struct {
	count int
	head  *Node
	tail  *Node
}

func (l *LinkedList) InsertHead(val int) {
	node := &Node{val: val}
	node.next = l.head
	if l.head != nil {
		l.head.prev = node
	}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.count++
}

func (l *LinkedList) InsertTail(val int) {
	node := &Node{val: val}
	node.prev = l.tail
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node
	if l.head == nil {
		l.head = node
	}
	l.count++
}

func (l *LinkedList) InsertPosition(val, pos int) {
	if !l.checkPosition(pos) {
		return
	}
	if pos == 0 {
		l.InsertHead(val)
		return
	}
	curr := l.head.next
	for i := 1; curr != nil; i++ {
		if i == pos {
			node := &Node{val: val}
			node.prev = curr.prev
			node.next = curr
			curr.prev.next = node
			curr.prev = node
			break
		}
		curr = curr.next
	}
	l.count++
}

func (l *LinkedList) UpdateHead(val int) {
	if !l.checkHead() {
		return
	}
	l.head.val = val
}

func (l *LinkedList) UpdateTail(val int) {
	if !l.checkTail() {
		return
	}
	l.tail.val = val
}

func (l *LinkedList) UpdatePosition(val, pos int) {
	if !l.checkPosition(pos) {
		return
	}
	curr := l.head
	for i := 0; curr != nil && i != pos; i++ {
		curr = curr.next
	}
	curr.val = val
}

func (l *LinkedList) RemoveHead() {
	if !l.checkHead() {
		return
	}
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	l.count--
}

func (l *LinkedList) RemoveTail() {
	if !l.checkTail() {
		return
	}
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}
	l.count--
}

func (l *LinkedList) RemovePosition(pos int) {
	if !l.checkPosition(pos) {
		return
	}
	curr := l.head
	for i := 0; curr != nil && i != pos; i++ {
		curr = curr.next
	}
	l.remove(curr)
}

func (l *LinkedList) RemoveValue(val int) {
	curr := l.head
	for i := 0; curr != nil; i++ {
		if curr.val == val {
			l.remove(curr)
		}
		curr = curr.next
	}
}

func (l *LinkedList) SearchPosition(pos int) {
	if !l.checkPosition(pos) {
		return
	}
	curr := l.head
	for i := 0; curr != nil && i != pos; i++ {
		curr = curr.next
	}
	utils.PrintDiv()
	fmt.Printf("Position %d has the value %d\n", pos, curr.val)
	fmt.Scanln()
}

func (l *LinkedList) SearchValue(val int) {
	poss := []int{}
	curr := l.head
	for i := 0; curr != nil; i++ {
		if curr.val == val {
			poss = append(poss, i)
		}
		curr = curr.next
	}
	utils.PrintDiv()
	if len(poss) > 0 {
		fmt.Printf("Value %d found at position(s): %+v\n", val, poss)
	} else {
		fmt.Printf("Value %d not found!\n", val)
	}
	fmt.Scanln()
}

func (l *LinkedList) Clear() {
	l.count = 0
	l.head = nil
	l.tail = nil
}

func (l *LinkedList) remove(node *Node) {
	if node == l.head {
		l.RemoveHead()
		return
	}
	if node == l.tail {
		l.RemoveTail()
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.count--
}

func (l *LinkedList) checkHead() bool {
	if l.head == nil {
		utils.PrintDiv()
		fmt.Println("[FAIL] Head is null!")
		fmt.Scanln()
		return false
	}
	return true
}

func (l *LinkedList) checkTail() bool {
	if l.tail == nil {
		utils.PrintDiv()
		fmt.Println("[FAIL] Tail is null!")
		fmt.Scanln()
		return false
	}
	return true
}

func (l *LinkedList) checkPosition(pos int) bool {
	if pos < 0 || pos >= l.count {
		utils.PrintDiv()
		fmt.Println("[FAIL] Invalid position!")
		fmt.Scanln()
		return false
	}
	return true
}

func (l *LinkedList) print() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.val, " ")
		curr = curr.next
	}
}

func Menu() {
	l := LinkedList{}
	op := -1
	for op != 0 {
		utils.ClearTerminal()
		utils.PrintDiv()
		fmt.Println("***          Linked List         ***")
		utils.PrintDiv()
		fmt.Printf("%2d | %s\n", 1, "Insert at head")
		fmt.Printf("%2d | %s\n", 2, "Insert at tail")
		fmt.Printf("%2d | %s\n", 3, "Insert at position")
		fmt.Printf("%2d | %s\n", 4, "Update at head")
		fmt.Printf("%2d | %s\n", 5, "Update at tail")
		fmt.Printf("%2d | %s\n", 6, "Update at position")
		fmt.Printf("%2d | %s\n", 7, "Remove at head")
		fmt.Printf("%2d | %s\n", 8, "Remove at tail")
		fmt.Printf("%2d | %s\n", 9, "Remove at position")
		fmt.Printf("%2d | %s\n", 10, "Remove by value")
		fmt.Printf("%2d | %s\n", 11, "Search at position")
		fmt.Printf("%2d | %s\n", 12, "Search by value")
		fmt.Printf("%2d | %s\n", 13, "Order ascending")
		fmt.Printf("%2d | %s\n", 14, "Order descending")
		fmt.Printf("%2d | %s\n", 15, "Clear list")
		fmt.Printf("%2d | %s\n", 0, "Back to main menu")
		utils.PrintDiv()
		fmt.Printf("%2d | ", l.count)
		l.print()
		fmt.Println()
		utils.PrintDiv()
		utils.ChooseOption()
		_, err := fmt.Scan(&op)
		utils.PrintDiv()
		if err != nil {
			utils.InvalidOption()
		} else {
			switch op {
			case 1:
				l.InsertHead(utils.ReadVal())
			case 2:
				l.InsertTail(utils.ReadVal())
			case 3:
				l.InsertPosition(utils.ReadVal(), utils.ReadPos())
			case 4:
				l.UpdateHead(utils.ReadVal())
			case 5:
				l.UpdateTail(utils.ReadVal())
			case 6:
				l.UpdatePosition(utils.ReadVal(), utils.ReadPos())
			case 7:
				l.RemoveHead()
			case 8:
				l.RemoveTail()
			case 9:
				l.RemovePosition(utils.ReadPos())
			case 10:
				l.RemoveValue(utils.ReadVal())
			case 11:
				l.SearchPosition(utils.ReadPos())
			case 12:
				l.SearchValue(utils.ReadVal())
			case 15:
				l.Clear()
			case 0:
			default:
				utils.InvalidOption()
			}
		}
	}
}
