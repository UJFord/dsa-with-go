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
}

type List struct {
	length int
	head   *Node
}

func (list List) Display() {

	var ViewLinkedList string
	currentNode := list.head

	for list.length != 0 {
		if list.length != 1 {
			ViewLinkedList += fmt.Sprintf("%s -> ", currentNode.value)
		} else {
			ViewLinkedList += fmt.Sprintf("%s", currentNode.value)
		}
		list.length--
		currentNode = currentNode.next
	}

	fmt.Println("\n===================\nList Items: " + ViewLinkedList)
}

func (list *List) GetLength() int {
	return list.length
}

func (list *List) Prepend(node *Node) {
	newHead := node
	secondNode := list.head
	list.head = newHead
	newHead.next = secondNode
	list.length++
}

func (list *List) Append(node *Node) {

	emptyList := list.head == nil
	if emptyList {

		list.head = node

	} else {

		currentNode := list.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = node
	}

	list.length++
}

func (list *List) InsertAt(countdown int, node *Node) {

	currentNode := list.head

	for ; countdown > 0; countdown-- {

		nextNode := currentNode.next

		if countdown == 1 {
			currentNode.next = node
			node.next = nextNode
		}

		currentNode = nextNode
	}

	list.length++
}

func (list *List) RemoveAt(countdown int) *Node {

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

func (list *List) RemoveHead() *Node {

	head := list.head
	list.head = list.head.next

	list.length--

	return head
}

func (list *List) RemoveTail() *Node {

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

func (list *List) Get(countdown int) *Node {

	currentNode := list.head

	for ; countdown > -1; countdown-- {
		if countdown == 0 {
			return currentNode
		}

		currentNode = currentNode.next
	}

	return currentNode
}
