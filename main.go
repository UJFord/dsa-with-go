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

	fmt.Println(ViewLinkedList)
}

func (list *List) GetLength() int {
	return list.length
}

func (list *List) Prepend(value any) {
	newHead := &Node{value: value}
	secondNode := list.head
	list.head = newHead
	newHead.next = secondNode
	list.length++
}
