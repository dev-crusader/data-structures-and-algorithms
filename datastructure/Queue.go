/* This is an implementation of doubly linkedlist which holds integer in its node.
Queue has a pointer to both head and tail, and its size.
Each node has a 'value' that stores the integer, and pointer to its next and previous node.
 This queue implmentation adds elements to the front of the list
 and removes element from the back of the queue with the regular methods
 i.e. Push() and Pop()
 AddToEnd and RemoveFromFront method is also added as a part of doubly linkedlist
*/

package datastructure

import "fmt"

type Queue struct {
	head, tail *Node
	size       int
}

type Node struct {
	next, prev *Node
	value      int
}

func (a *Queue) Push2(v int) {
	if a.tail == nil {
		a.tail = &Node{
			value: v,
		}
		a.head = a.tail
	} else if a.head == a.tail {
		a.head = &Node{
			value: v,
		}
		a.head.next = a.tail
		a.tail.prev = a.head
	} else {
		temp := a.head
		a.head = &Node{
			value: v,
		}
		a.head.next = temp
		temp.prev = a.head
	}
	a.size++
}

//Push method add element to the front of Queue
func (a *Queue) Push(v int) {
	node := &Node{value: v}
	if a.head == nil {
		a.tail = node
	} else {
		node.next = a.head
		a.head.prev = node
	}
	a.head = node
	a.size++
}

func (a *Queue) isEmpty() bool {
	return a.size == 0
}

// Pop method removes elements from back of the queue.
func (a *Queue) Pop() (int, bool) {
	if a.isEmpty() {
		return -1, false
	}

	if a.tail.prev == nil {
		a.head = nil
	} else {
		a.tail.prev.next = nil
	}

	popped := a.tail
	a.tail = a.tail.prev
	a.size--
	popped.prev = nil
	return popped.value, true
}

func (a *Queue) AddToEnd(v int) {
	node := &Node{value: v}

	if a.head == nil {
		a.head = node
	} else {
		node.prev = a.tail
		a.tail.next = node
	}

	a.tail = node
	a.size++
}

func (a *Queue) RemoveFromFront() (int, bool) {
	if a.isEmpty() {
		return -1, false
	}

	if a.head.next == nil {
		a.tail = nil
	} else {
		a.head.next.prev = nil
	}

	popped := a.head
	a.head = a.head.next
	a.size--
	popped.next = nil
	return popped.value, true
}

func (a *Queue) GetAll() []int {
	result := make([]int, 0)

	for i := a.head; i != nil; i = i.next {
		result = append(result, i.value)
	}
	return result
}

func RunQueue() {
	o := Queue{}
	o.Push(5)
	o.Push(-45)
	o.Push(432)
	o.Push(57)
	o.Push(86)
	o.Push(-12)
	o.AddToEnd(75)
	o.AddToEnd(93)
	fmt.Println(o.GetAll())
	fmt.Println("Size: ", o.size)
	fmt.Println(o.Pop())
	fmt.Println(o.GetAll())
	fmt.Println(o.Pop())
	fmt.Println(o.Pop())
	fmt.Println(o.Pop())
	fmt.Println(o.isEmpty())
	fmt.Println("Size: ", o.size)
	fmt.Println(o.Pop())
	fmt.Println(o.Pop())
	fmt.Println(o.RemoveFromFront())
	fmt.Println(o.RemoveFromFront())
	fmt.Println(o.Pop())
}
