package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func reverseIterative(head *Node) *Node {
	var prev *Node
	current := head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

func reverseRecurse(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	newHead := reverseRecurse(head.next)

	head.next.next = head
	head.next = nil

	return newHead
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d -> ", head.value)
		head = head.next
	}
	fmt.Println("nil")
}

func main() {
	head := &Node{value: 1}
	head.next = &Node{value: 2}
	head.next.next = &Node{value: 3}
	head.next.next.next = &Node{value: 4}
	head.next.next.next.next = &Node{value: 5}

	fmt.Println("Lista original:")
	printList(head)

	reversedIterative := reverseIterative(head)
	fmt.Println("Lista reversa (iterativa):")
	printList(reversedIterative)

	reversedRecursive := reverseRecurse(reversedIterative)
	fmt.Println("Lista reversa novamente (recursiva):")
	printList(reversedRecursive)
}
