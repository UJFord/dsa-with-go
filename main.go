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

var ViewLinkedList string

func (list List) Display() {

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

	fmt.Println(ViewLinkedList + "\n")
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

func (list *List) RemoveHead() *Node {

	currentHead := list.head
	if currentHead != nil {

		if list.head.next != nil {
			list.head = list.head.next
		} else {
			list.head = nil
		}

		return list.head
	}

	return nil
}
