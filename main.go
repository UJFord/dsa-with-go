package main

import (
	"fmt"
	"math"
)

func LinearSearch(list []int, target int) bool {

	for i := 0; i < len(list); i++ {

		if list[i] == target {
			return true
		}

	}

	return false
}

func BinarySearch(list []int, target int) bool {

	hi := len(list)
	lo := 0

	for lo < hi {
		m := lo + ((hi - lo) / 2)

		if list[m] == target {
			return true
		}
		if list[m] > target {
			hi = m
		}
		if list[m] < target {
			lo = m + 1
		}
	}

	return false
}

func CrystalBall(broke []bool) int {

	n := len(broke)
	jmpAmount := int(math.Floor(math.Sqrt(float64(n))))

	var i = jmpAmount
	for ; i < n; i += jmpAmount {

		if broke[i] {

			break
		}
	}

	i -= jmpAmount

	for j := 0; j <= jmpAmount && j < n; j++ {

		if broke[i] {

			return i
		}
		i++
	}

	return -1
}

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {

			if arr[j] > arr[j+1] {
				n := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = n
			}
		}
	}

	return arr
}

type Node struct {
	value any
	next  *Node
	prev  *Node
}

type SinglyLinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (list *SinglyLinkedList) Display() string {

	var ViewLinkedSinglyLinkedList string
	if list.head == nil {
		return "empty"
	}

	currentNode := list.head

	countdown := list.length
	for countdown > 0 {
		if countdown > 1 {
			ViewLinkedSinglyLinkedList += fmt.Sprintf("%s -> ", currentNode.value)
		} else {
			ViewLinkedSinglyLinkedList += fmt.Sprintf("%s", currentNode.value)
		}
		currentNode = currentNode.next
		countdown--
	}

	fmt.Println("===================\nList Items: " + ViewLinkedSinglyLinkedList)
	fmt.Printf("List Count: %d\n===================\n\n", list.length)

	return ViewLinkedSinglyLinkedList
}

func (list *SinglyLinkedList) Prepend(node *Node) {

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head = node
	}

	list.length++
}

func (list *SinglyLinkedList) Append(node *Node) {

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}

	list.length++
}

func (list *SinglyLinkedList) InsertAt(countdown int, node *Node) {

	if countdown == 0 {

		node.next = list.head
		list.head = node

	} else {

		currentNode := list.head
		nextNode := currentNode.next

		for ; countdown > 0; countdown-- {

			if countdown > 1 {

				currentNode = nextNode

			} else if countdown == 1 {

				currentNode.next = node
				node.next = nextNode
			}
		}
	}

	list.length++
}

func (list *SinglyLinkedList) RemoveAt(countdown int) *Node {

	var removed *Node

	if list.head == nil {
		return nil
	}

	removeFirstNode := countdown == 0
	if removeFirstNode {
		removed = list.head
		list.head = list.head.next
		list.length--
		return removed
	}

	currentNode := list.head
	for ; countdown > 0; countdown-- {

		nextNode := currentNode.next

		if countdown == 1 {
			removed = nextNode
			currentNode.next = currentNode.next.next
		}

		currentNode = nextNode
	}

	list.length--

	return removed
}

func (list *SinglyLinkedList) RemoveHead() *Node {

	if list.head != nil {
		head := list.head
		list.head = list.head.next
		list.length--

		return head
	}

	return nil
}

func (list *SinglyLinkedList) RemoveTail() *Node {

	var removed *Node

	if list.head == nil {
		return nil
	}

	if list.head.next == nil {
		removed = list.head
		list.head = nil
		list.length--
		return removed
	}

	currentNode := list.head
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}

	removed = currentNode.next
	currentNode.next = nil

	list.length--
	return removed
}

func (list *SinglyLinkedList) Get(countdown int) *Node {

	currentNode := list.head

	for ; countdown > -1; countdown-- {
		if countdown == 0 {
			return currentNode
		}

		currentNode = currentNode.next
	}

	return currentNode
}

func (queue *SinglyLinkedList) Enqueue(value *Node) {

	if queue.head == nil {
		queue.head = value
		queue.tail = value
	} else {
		queue.tail.next = value
		queue.tail = value
	}

	queue.length++
}

func (queue *SinglyLinkedList) Deque() *Node {

	if queue.head != nil {
		head := queue.head
		queue.head = queue.head.next
		queue.length--

		return head
	}

	return nil
}

func (queue *SinglyLinkedList) Peek() *Node {
	if queue.head != nil {
		return queue.head
	}

	return nil
}
