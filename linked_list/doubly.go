package linked_list

import (
	"fmt"

	"github.com/ceferrari/go-algods/utils"
)

type DoublyNode struct {
	val  int
	prev *DoublyNode
	next *DoublyNode
}

type DoublyLinkedList struct {
	count int
	head  *DoublyNode
	tail  *DoublyNode
}

func (l *DoublyLinkedList) InsertHead(val int) {
	node := &DoublyNode{val: val}
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

func (l *DoublyLinkedList) InsertTail(val int) {
	node := &DoublyNode{val: val}
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

func (l *DoublyLinkedList) InsertPosition(val, pos int) {
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
			node := &DoublyNode{val: val}
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

func (l *DoublyLinkedList) UpdateHead(val int) {
	if !l.checkHead() {
		return
	}
	l.head.val = val
}

func (l *DoublyLinkedList) UpdateTail(val int) {
	if !l.checkTail() {
		return
	}
	l.tail.val = val
}

func (l *DoublyLinkedList) UpdatePosition(val, pos int) {
	if !l.checkPosition(pos) {
		return
	}
	curr := l.head
	for i := 0; curr != nil && i != pos; i++ {
		curr = curr.next
	}
	curr.val = val
}

func (l *DoublyLinkedList) RemoveHead() {
	if !l.checkHead() {
		return
	}
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	l.count--
}

func (l *DoublyLinkedList) RemoveTail() {
	if !l.checkTail() {
		return
	}
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}
	l.count--
}

func (l *DoublyLinkedList) RemovePosition(pos int) {
	if !l.checkPosition(pos) {
		return
	}
	curr := l.head
	for i := 0; curr != nil && i != pos; i++ {
		curr = curr.next
	}
	l.remove(curr)
}

func (l *DoublyLinkedList) RemoveValue(val int) {
	curr := l.head
	for curr != nil {
		if curr.val == val {
			l.remove(curr)
		}
		curr = curr.next
	}
}

func (l *DoublyLinkedList) SearchPosition(pos int) {
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

func (l *DoublyLinkedList) SearchValue(val int) {
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

func (l *DoublyLinkedList) order(dir string) {
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

func (l *DoublyLinkedList) OrderAscending() {
	l.order("asc")
}

func (l *DoublyLinkedList) OrderDescending() {
	l.order("desc")
}

func (l *DoublyLinkedList) ReverseDefault() {
	head := l.head
	tail := l.tail
	for i, j := 0, l.count-1; i < j; i, j = i+1, j-1 {
		temp := tail.val
		tail.val = head.val
		head.val = temp
		head = head.next
		tail = tail.prev
	}
}

func (l *DoublyLinkedList) ReverseIterative() {
	l.tail = l.head
	var prev, curr, next *DoublyNode = nil, l.head, nil
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *DoublyLinkedList) ReverseRecursive(head *DoublyNode) *DoublyNode {
	if head == nil || head.next == nil {
		l.head = head
		return head
	}
	rest := l.ReverseRecursive(head.next)
	if head.next.next != nil {
		l.tail = head.next.next
	}
	head.next.next = head
	head.next = nil
	return rest
}

func (l *DoublyLinkedList) Clear() {
	l.count = 0
	l.head = nil
	l.tail = nil
}

func (l *DoublyLinkedList) remove(node *DoublyNode) {
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

func (l *DoublyLinkedList) checkHead() bool {
	if l.head == nil {
		utils.PrintDiv()
		fmt.Println("[FAIL] Head is null!")
		fmt.Scanln()
		return false
	}
	return true
}

func (l *DoublyLinkedList) checkTail() bool {
	if l.tail == nil {
		utils.PrintDiv()
		fmt.Println("[FAIL] Tail is null!")
		fmt.Scanln()
		return false
	}
	return true
}

func (l *DoublyLinkedList) checkPosition(pos int) bool {
	if pos < 0 || pos >= l.count {
		utils.PrintDiv()
		fmt.Println("[FAIL] Invalid position!")
		fmt.Scanln()
		return false
	}
	return true
}

func (l *DoublyLinkedList) print() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.val, " ")
		curr = curr.next
	}
}

func DoublyMenu() {
	l := DoublyLinkedList{}
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
		fmt.Println("***      Doubly Linked List      ***")
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
		fmt.Printf("%2d | %s\n", 15, "Reverse (default)")
		fmt.Printf("%2d | %s\n", 16, "Reverse (iterative)")
		fmt.Printf("%2d | %s\n", 17, "Reverse (recursive)")
		fmt.Printf("%2d | %s\n", 18, "Clear list")
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
				l.ReverseDefault()
			case 16:
				l.ReverseIterative()
			case 17:
				l.ReverseRecursive(l.head)
			case 18:
				l.Clear()
			case 0:
			default:
				utils.InvalidOption()
			}
		}
	}
}
