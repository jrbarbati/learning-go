package main

import "fmt"

type node[T comparable] struct {
	val  T
	next *node[T]
}

type linkedList[T comparable] struct {
	head *node[T]
}

func newLinkedList[T comparable]() linkedList[T] {
	return linkedList[T]{}
}

// Add adds a new element to the end of the linked list
func (ll *linkedList[T]) Add(val T) {
	if ll.head == nil {
		ll.head = &node[T]{val: val}
		return
	}

	curr := ll.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &node[T]{val: val, next: nil}
}

// Insert adds an element at the specified position in the linked list
func (ll *linkedList[T]) Insert(val T, index int) {
	if index <= 0 {
		newHead := &node[T]{val: val, next: ll.head}
		ll.head = newHead
		return
	}

	var idx = 0
	var prev *node[T]
	curr := ll.head

	for curr != nil {
		if idx == index {
			break
		}
		prev = curr
		curr = curr.next
		idx++
	}

	if prev != nil {
		prev.next = &node[T]{val: val, next: curr}
	}
}

// Index return the position of the supplied value, -1 if it's not present
func (ll *linkedList[T]) Index(val T) int {
	var idx = 0
	curr := ll.head

	for curr != nil {
		if curr.val == val {
			return idx
		}

		curr = curr.next
		idx++
	}

	return -1
}

func main() {
	ll := newLinkedList[uint16]()
	ll.Add(5)
	ll.Add(10)
	fmt.Println(ll.Index(5))
	fmt.Println(ll.Index(10))
	fmt.Println(ll.Index(20))

	ll.Insert(100, 0)
	fmt.Println(ll.Index(5))
	fmt.Println(ll.Index(10))
	fmt.Println(ll.Index(20))
	fmt.Println(ll.Index(100))

	ll.Insert(200, 1)
	fmt.Println(ll.Index(5))
	fmt.Println(ll.Index(10))
	fmt.Println(ll.Index(200))
	fmt.Println(ll.Index(20))
	fmt.Println(ll.Index(100))

	for curNode := ll.head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.val)
	}

	ll.Insert(300, 10)
	for curNode := ll.head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.val)
	}

	ll.Add(400)
	for curNode := ll.head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.val)
	}

	ll.Insert(500, 6)
	for curNode := ll.head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.val)
	}
}
