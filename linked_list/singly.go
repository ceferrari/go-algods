package linked_list

import (
	"fmt"

	"github.com/ceferrari/go-algods/utils"
)

type SinglyNode struct {
	val  int
	next *SinglyNode
}

type SinglyLinkedList struct {
	count int
	head  *SinglyNode
}

func (l *SinglyLinkedList) InsertHead(val int) {
	node := &SinglyNode{val: val}
	node.next = l.head
	l.head = node
	l.count++
}

func (l *SinglyLinkedList) InsertTail(val int) {
	if l.head == nil {
		l.InsertHead(val)
		return
	}
	node := &SinglyNode{val: val}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = node
	l.count++
}

func (l *SinglyLinkedList) InsertPosition(val, pos int) {
	if !l.checkPosition(pos) {
		return
	}
	if pos == 0 {
		l.InsertHead(val)
		return
	}
	curr := l.head
	for i := 1; curr.next != nil; i++ {
		if i == pos {
			node := &SinglyNode{val: val}
			node.next = curr.next
			curr.next = node
			break
		}
		curr = curr.next
	}
	l.count++
}

func (l *SinglyLinkedList) UpdateHead(val int) {
	if !l.checkHead() {
		return
	}
	l.head.val = val
}

func (l *SinglyLinkedList) UpdateTail(val int) {
	if !l.checkHead() {
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.val = val
}

func (l *SinglyLinkedList) UpdatePosition(val, pos int) {
	if !l.checkPosition(pos) {
		return
	}
	curr := l.head
	for i := 0; curr != nil && i != pos; i++ {
		curr = curr.next
	}
	curr.val = val
}

func (l *SinglyLinkedList) RemoveHead() {
	if !l.checkHead() {
		return
	}
	l.head = l.head.next
	l.count--
}

func (l *SinglyLinkedList) RemoveTail() {
	if !l.checkHead() {
		return
	}
	curr := l.head
	for curr.next != nil && curr.next.next != nil {
		curr = curr.next
	}
	curr.next = nil
	l.count--
}

func (l *SinglyLinkedList) RemovePosition(pos int) {
	if !l.checkPosition(pos) {
		return
	}
	if pos == 0 {
		l.RemoveHead()
		return
	}
	curr := l.head
	for i := 1; curr.next != nil && i != pos; i++ {
		if i != pos {
			curr.next = curr.next.next
		}
		curr = curr.next
	}
	l.count--
}

func (l *SinglyLinkedList) RemoveValue(val int) {
	curr := l.head
	for curr.next != nil {
		if curr.next.val == val {
			curr.next = curr.next.next
			l.count--
		}
		curr = curr.next
	}
	if l.head.val == val {
		l.RemoveHead()
	}
}

func (l *SinglyLinkedList) SearchPosition(pos int) {
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

func (l *SinglyLinkedList) SearchValue(val int) {
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

func (l *SinglyLinkedList) order(dir string) {
	for i := l.head; i != nil; i = i.next {
		for j := i.next; j != nil; j = j.next {
			if (dir == "asc" && i.val > j.val) || (dir == "desc" && i.val < j.val) {
				temp := j.val
				j.val = i.val
				i.val = temp
			}
		}
	}
}

func (l *SinglyLinkedList) OrderAscending() {
	l.order("asc")
}

func (l *SinglyLinkedList) OrderDescending() {
	l.order("desc")
}

func (l *SinglyLinkedList) Reverse() {
	a := make([]int, l.count)
	curr := l.head
	for i := 0; curr != nil; i++ {
		a[i] = curr.val
		curr = curr.next
	}
	curr = l.head
	for i := len(a) - 1; curr != nil; i-- {
		curr.val = a[i]
		curr = curr.next
	}
}

func (l *SinglyLinkedList) Clear() {
	l.count = 0
	l.head = nil
}

func (l *SinglyLinkedList) checkHead() bool {
	if l.head == nil {
		utils.PrintDiv()
		fmt.Println("[FAIL] Head is null!")
		fmt.Scanln()
		return false
	}
	return true
}

func (l *SinglyLinkedList) checkPosition(pos int) bool {
	if pos < 0 || pos >= l.count {
		utils.PrintDiv()
		fmt.Println("[FAIL] Invalid position!")
		fmt.Scanln()
		return false
	}
	return true
}

func (l *SinglyLinkedList) print() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.val, " ")
		curr = curr.next
	}
}

func SinglyMenu() {
	l := SinglyLinkedList{}
	l.InsertTail(6)
	l.InsertTail(3)
	l.InsertTail(1)
	l.InsertTail(4)
	l.InsertTail(9)
	l.InsertTail(5)
	l.InsertTail(8)
	l.InsertTail(7)
	l.InsertTail(2)
	op := -1
	for op != 0 {
		utils.ClearTerminal()
		utils.PrintDiv()
		fmt.Println("***      Singly Linked List      ***")
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
		fmt.Printf("%2d | %s\n", 15, "Reverse list")
		fmt.Printf("%2d | %s\n", 16, "Clear list")
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
			case 13:
				l.OrderAscending()
			case 14:
				l.OrderDescending()
			case 15:
				l.Reverse()
			case 16:
				l.Clear()
			case 0:
			default:
				utils.InvalidOption()
			}
		}
	}
}
