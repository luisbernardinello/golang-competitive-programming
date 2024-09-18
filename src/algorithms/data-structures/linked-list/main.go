package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedList struct {
	head   *Node
	length int
}

func (l *linkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length --

	}
	fmt.Println()
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0{
		return 
	}

	if l.head.data == value{
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value{
		if previousToDelete.next.next == nil{
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}


func main() {
	myList := linkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 27}
	node3 := &Node{data: 34}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	myList.printListData()

	myList.deleteWithValue(27)

	myList.printListData()

}
